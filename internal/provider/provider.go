package provider

import "context"

// StreamEvent represents an event from the SSE stream.
type StreamEvent struct {
	Type string      // "content", "tool_calls", "done", "error"
	Data interface{} // string for content/error, []ToolCall for tool_calls, map[string]interface{} for done
}

// ToolCall from the streaming response.
type ToolCall struct {
	ID       string
	Type     string // always "function"
	Function ToolCallFunction
}

// ToolCallFunction contains the function name and arguments.
type ToolCallFunction struct {
	Name      string
	Arguments string
}

// Provider streams chat completions.
type Provider interface {
	StreamChat(ctx context.Context, model string, messages []map[string]interface{}, tools []map[string]interface{}) <-chan StreamEvent
}
