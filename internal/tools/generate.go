package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
	"time"
)

const queueBase = "https://queue.fal.run"

// GenerateTool runs a model via the fal queue API with live progress.
type GenerateTool struct {
	AuthKey string
}

func (t *GenerateTool) Name() string { return "generate" }
func (t *GenerateTool) Description() string {
	return "Run a model with the given input parameters. Submits the job, waits for completion, and returns the result including output URLs."
}
func (t *GenerateTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"endpoint_id": map[string]interface{}{
				"type":        "string",
				"description": "Model endpoint ID (e.g. 'fal-ai/flux/dev')",
			},
			"input": map[string]interface{}{
				"type":        "object",
				"description": `Input parameters for the model (e.g. {"prompt": "a cat", "image_size": "landscape_16_9"})`,
			},
		},
		"required": []string{"endpoint_id", "input"},
	}
}

// ExecuteWithContext runs a generation with cancellation support.
func (t *GenerateTool) ExecuteWithContext(ctx context.Context, args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	return t.executeInternal(ctx, args, onProgress)
}

func (t *GenerateTool) Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	return t.executeInternal(context.Background(), args, onProgress)
}

func (t *GenerateTool) executeInternal(ctx context.Context, args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	endpointID, _ := args["endpoint_id"].(string)
	inputData, _ := args["input"].(map[string]interface{})

	// Auto-unwrap: if input is empty but extra keys exist, use them as input
	if len(inputData) == 0 {
		knownMeta := map[string]bool{"endpoint_id": true, "input": true}
		extra := make(map[string]interface{})
		for k, v := range args {
			if !knownMeta[k] {
				extra[k] = v
			}
		}
		if len(extra) > 0 {
			inputData = extra
		}
	}

	if t.AuthKey == "" {
		return map[string]interface{}{"ok": false, "error": "Not authenticated"}, nil
	}

	if onProgress != nil {
		onProgress(map[string]interface{}{"state": "SUBMITTING", "endpoint_id": endpointID})
	}

	startTime := time.Now()
	client := &http.Client{Timeout: 30 * time.Second}

	// Submit
	submitBody, _ := json.Marshal(inputData)
	submitURL := fmt.Sprintf("%s/%s", queueBase, endpointID)
	req, err := http.NewRequest("POST", submitURL, strings.NewReader(string(submitBody)))
	if err != nil {
		return map[string]interface{}{"ok": false, "error": fmt.Sprintf("Submit failed: %v", err)}, nil
	}
	req.Header.Set("Authorization", t.AuthKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return map[string]interface{}{"ok": false, "error": fmt.Sprintf("Submit failed: %v", err)}, nil
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == 422 {
		var detail interface{}
		if err := json.Unmarshal(body, &detail); err != nil {
			detail = truncStr(string(body), 500)
		}
		return map[string]interface{}{
			"ok":    false,
			"error": fmt.Sprintf("422 Validation Error: %v", detail),
			"hint":  "Call model_info to check correct parameter names, types, and valid values, then retry.",
		}, nil
	}

	if resp.StatusCode >= 400 {
		var detail interface{}
		if err := json.Unmarshal(body, &detail); err != nil {
			detail = truncStr(string(body), 500)
		}
		return map[string]interface{}{"ok": false, "error": fmt.Sprintf("HTTP %d: %v", resp.StatusCode, detail)}, nil
	}

	var submitData map[string]interface{}
	if err := json.Unmarshal(body, &submitData); err != nil {
		return map[string]interface{}{"ok": false, "error": fmt.Sprintf("Invalid submit response: %v", err)}, nil
	}

	requestID, _ := submitData["request_id"].(string)
	urls := buildURLs(endpointID, requestID)

	statusURL := strOr(submitData["status_url"], urls["status_url"])
	responseURL := strOr(submitData["response_url"], urls["response_url"])

	// Poll with progress
	pollInterval := 0.5
	pollClient := &http.Client{Timeout: 15 * time.Second}

	for {
		select {
		case <-ctx.Done():
			return map[string]interface{}{"ok": false, "error": "generation cancelled"}, nil
		default:
		}

		pollReq, _ := http.NewRequestWithContext(ctx, "GET", statusURL+"?logs=1", nil)
		pollReq.Header.Set("Authorization", t.AuthKey)

		pollResp, err := pollClient.Do(pollReq)
		if err != nil {
			time.Sleep(time.Duration(pollInterval * float64(time.Second)))
			continue
		}

		pollBody, _ := io.ReadAll(pollResp.Body)
		pollResp.Body.Close()

		var status map[string]interface{}
		if err := json.Unmarshal(pollBody, &status); err != nil {
			time.Sleep(time.Duration(pollInterval * float64(time.Second)))
			continue
		}

		state, _ := status["status"].(string)
		if state == "" {
			state = "UNKNOWN"
		}
		elapsed := time.Since(startTime).Seconds()

		if onProgress != nil {
			progress := map[string]interface{}{
				"state":       state,
				"elapsed":     elapsed,
				"endpoint_id": endpointID,
				"request_id":  requestID,
			}
			if state == "IN_QUEUE" {
				if pos, ok := status["queue_position"]; ok {
					progress["position"] = pos
				} else {
					progress["position"] = "?"
				}
			} else if state == "IN_PROGRESS" {
				if logs, ok := status["logs"].([]interface{}); ok {
					logMsgs := make([]string, 0, len(logs))
					for _, l := range logs {
						if lm, ok := l.(map[string]interface{}); ok {
							if msg, ok := lm["message"].(string); ok {
								logMsgs = append(logMsgs, msg)
							}
						}
					}
					progress["logs"] = logMsgs
				}
			}
			onProgress(progress)
		}

		if state == "COMPLETED" {
			break
		}

		if errVal, ok := status["error"]; ok && errVal != nil {
			var errorMsg string
			switch e := errVal.(type) {
			case map[string]interface{}:
				b, _ := json.Marshal(e)
				errorMsg = truncStr(string(b), 500)
			default:
				errorMsg = truncStr(fmt.Sprintf("%v", e), 500)
			}
			return map[string]interface{}{
				"ok":    false,
				"error": fmt.Sprintf("Execution failed: %s", errorMsg),
				"hint":  "Check model_info for valid parameters and retry with corrected input.",
			}, nil
		}

		time.Sleep(time.Duration(pollInterval * float64(time.Second)))
		pollInterval = math.Min(pollInterval*1.2, 2.0)
	}

	// Fetch result
	resultClient := &http.Client{Timeout: 60 * time.Second}
	resultReq, _ := http.NewRequest("GET", responseURL, nil)
	resultReq.Header.Set("Authorization", t.AuthKey)

	resultResp, err := resultClient.Do(resultReq)
	if err != nil {
		return map[string]interface{}{"ok": false, "error": fmt.Sprintf("Failed to fetch result: %v", err)}, nil
	}
	defer resultResp.Body.Close()

	resultBody, _ := io.ReadAll(resultResp.Body)

	if resultResp.StatusCode == 422 {
		var detail interface{}
		json.Unmarshal(resultBody, &detail)
		return map[string]interface{}{
			"ok":    false,
			"error": fmt.Sprintf("Result fetch 422: %v", detail),
			"hint":  "The input parameters were likely invalid. Call model_info to check correct params and retry.",
		}, nil
	}

	if resultResp.StatusCode >= 400 {
		var detail interface{}
		json.Unmarshal(resultBody, &detail)
		return map[string]interface{}{"ok": false, "error": fmt.Sprintf("Result fetch HTTP %d: %v", resultResp.StatusCode, detail)}, nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resultBody, &result); err != nil {
		return map[string]interface{}{"ok": false, "error": fmt.Sprintf("Invalid result JSON: %v", err)}, nil
	}

	elapsed := time.Since(startTime).Seconds()
	if onProgress != nil {
		onProgress(map[string]interface{}{"state": "COMPLETED", "elapsed": elapsed, "endpoint_id": endpointID})
	}

	return map[string]interface{}{
		"ok":              true,
		"request_id":      requestID,
		"endpoint_id":     endpointID,
		"elapsed_seconds": math.Round(elapsed*10) / 10,
		"result":          result,
	}, nil
}

func buildURLs(endpointID, requestID string) map[string]string {
	parts := strings.Split(endpointID, "/")
	baseApp := endpointID
	if len(parts) >= 2 {
		baseApp = strings.Join(parts[:2], "/")
	}
	base := fmt.Sprintf("%s/%s/requests/%s", queueBase, baseApp, requestID)
	return map[string]string{
		"status_url":   base + "/status",
		"response_url": base,
		"cancel_url":   base + "/cancel",
	}
}

func strOr(val interface{}, fallback string) string {
	if s, ok := val.(string); ok && s != "" {
		return s
	}
	return fallback
}

func truncStr(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen]
	}
	return s
}
