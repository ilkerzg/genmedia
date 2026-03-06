package provider

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/fal-ai/genmedia/internal/config"
)

// OpenRouterProvider implements Provider using OpenRouter's API.
type OpenRouterProvider struct {
	AuthKey string
	BaseURL string
}

func (p *OpenRouterProvider) baseURL() string {
	if p.BaseURL != "" {
		return p.BaseURL
	}
	return config.OpenRouterBase
}

// StreamChat sends a streaming chat completion request and returns events on a channel.
func (p *OpenRouterProvider) StreamChat(ctx context.Context, model string, messages []map[string]interface{}, tools []map[string]interface{}) <-chan StreamEvent {
	ch := make(chan StreamEvent, 64)

	go func() {
		defer close(ch)

		sanitized := make([]map[string]interface{}, len(messages))
		for i, msg := range messages {
			m := make(map[string]interface{}, len(msg))
			for k, v := range msg {
				m[k] = v
			}
			if m["content"] == nil {
				m["content"] = ""
			}
			sanitized[i] = m
		}

		body := map[string]interface{}{
			"model":    model,
			"messages": sanitized,
			"tools":    tools,
			"stream":   true,
		}
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			ch <- StreamEvent{Type: "error", Data: fmt.Sprintf("marshal error: %v", err)}
			return
		}

		url := p.baseURL() + "/chat/completions"

		// Use a long-lived context for the streaming connection (not cancelled after Do).
		streamCtx, streamCancel := context.WithTimeout(ctx, 5*time.Minute)
		defer streamCancel()

		var resp *http.Response
		maxRetries := 3
		for attempt := 0; attempt < maxRetries; attempt++ {
			req, err := http.NewRequestWithContext(streamCtx, http.MethodPost, url, strings.NewReader(string(bodyBytes)))
			if err != nil {
				ch <- StreamEvent{Type: "error", Data: fmt.Sprintf("request error: %v", err)}
				return
			}
			req.Header.Set("Authorization", p.AuthKey)
			req.Header.Set("Content-Type", "application/json")

			resp, err = http.DefaultClient.Do(req)

			if err != nil {
				if attempt < maxRetries-1 {
					time.Sleep(time.Duration(attempt+1) * time.Second)
					continue
				}
				p.dumpError(model, messages, tools, fmt.Sprintf("connection error: %v", err))
				ch <- StreamEvent{Type: "error", Data: fmt.Sprintf("connection error after retries: %v", err)}
				return
			}

			if resp.StatusCode >= 500 {
				errBody := make([]byte, 4096)
				n, _ := resp.Body.Read(errBody)
				resp.Body.Close()
				if attempt < maxRetries-1 {
					time.Sleep(time.Duration(attempt+1) * time.Second)
					continue
				}
				errStr := string(errBody[:n])
				p.dumpError(model, messages, tools, errStr)
				ch <- StreamEvent{Type: "error", Data: fmt.Sprintf("server error %d: %s", resp.StatusCode, errStr)}
				return
			}

			if resp.StatusCode != http.StatusOK {
				errBody := make([]byte, 4096)
				n, _ := resp.Body.Read(errBody)
				resp.Body.Close()
				errStr := string(errBody[:n])
				p.dumpError(model, messages, tools, errStr)
				ch <- StreamEvent{Type: "error", Data: fmt.Sprintf("HTTP %d: %s", resp.StatusCode, errStr)}
				return
			}

			break
		}

		defer resp.Body.Close()

		var collected strings.Builder
		toolCalls := make(map[int]*ToolCall)

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.HasPrefix(line, "data: ") {
				continue
			}
			data := strings.TrimPrefix(line, "data: ")
			if data == "[DONE]" {
				break
			}

			var chunk map[string]interface{}
			if err := json.Unmarshal([]byte(data), &chunk); err != nil {
				continue
			}

			choices, ok := chunk["choices"].([]interface{})
			if !ok || len(choices) == 0 {
				continue
			}
			choice, ok := choices[0].(map[string]interface{})
			if !ok {
				continue
			}
			delta, ok := choice["delta"].(map[string]interface{})
			if !ok {
				continue
			}

			if content, ok := delta["content"].(string); ok && content != "" {
				collected.WriteString(content)
				ch <- StreamEvent{Type: "content", Data: content}
			}

			if tcs, ok := delta["tool_calls"].([]interface{}); ok {
				for _, tc := range tcs {
					tcMap, ok := tc.(map[string]interface{})
					if !ok {
						continue
					}
					idxF, ok := tcMap["index"].(float64)
					if !ok {
						continue
					}
					idx := int(idxF)

					if _, exists := toolCalls[idx]; !exists {
						toolCalls[idx] = &ToolCall{Type: "function"}
					}
					entry := toolCalls[idx]

					if id, ok := tcMap["id"].(string); ok && id != "" {
						entry.ID = id
					}
					if fn, ok := tcMap["function"].(map[string]interface{}); ok {
						if name, ok := fn["name"].(string); ok && name != "" {
							entry.Function.Name = name
						}
						if args, ok := fn["arguments"].(string); ok {
							entry.Function.Arguments += args
						}
					}
				}
			}
		}

		if len(toolCalls) > 0 {
			indices := make([]int, 0, len(toolCalls))
			for idx := range toolCalls {
				indices = append(indices, idx)
			}
			sort.Ints(indices)

			tcList := make([]ToolCall, 0, len(indices))
			for _, idx := range indices {
				tcList = append(tcList, *toolCalls[idx])
			}
			ch <- StreamEvent{Type: "tool_calls", Data: tcList}
		}

		finalMessage := map[string]interface{}{
			"role":    "assistant",
			"content": collected.String(),
		}
		if len(toolCalls) > 0 {
			indices := make([]int, 0, len(toolCalls))
			for idx := range toolCalls {
				indices = append(indices, idx)
			}
			sort.Ints(indices)
			tcList := make([]map[string]interface{}, 0, len(indices))
			for _, idx := range indices {
				tc := toolCalls[idx]
				tcList = append(tcList, map[string]interface{}{
					"id":   tc.ID,
					"type": tc.Type,
					"function": map[string]interface{}{
						"name":      tc.Function.Name,
						"arguments": tc.Function.Arguments,
					},
				})
			}
			finalMessage["tool_calls"] = tcList
		}
		ch <- StreamEvent{Type: "done", Data: finalMessage}
	}()

	return ch
}

func (p *OpenRouterProvider) dumpError(model string, messages []map[string]interface{}, tools []map[string]interface{}, errBody string) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}
	cacheDir := filepath.Join(home, ".cache", "fal-dev")
	os.MkdirAll(cacheDir, 0o755)

	truncated := make([]map[string]interface{}, len(messages))
	for i, msg := range messages {
		m := map[string]interface{}{"role": msg["role"]}
		if content, ok := msg["content"].(string); ok && len(content) > 200 {
			m["content"] = content[:200] + "..."
		} else {
			m["content"] = msg["content"]
		}
		truncated[i] = m
	}

	dump := map[string]interface{}{
		"model":       model,
		"messages":    truncated,
		"tools_count": len(tools),
		"error_body":  errBody,
		"timestamp":   time.Now().Format(time.RFC3339),
	}

	data, err := json.MarshalIndent(dump, "", "  ")
	if err != nil {
		return
	}
	os.WriteFile(filepath.Join(cacheDir, "last_error.json"), data, 0o600)
}
