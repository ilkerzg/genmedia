package tui

import (
	"time"
)

// StreamChunkMsg carries a chunk of streamed assistant content.
type StreamChunkMsg struct {
	Content string
}

// StreamDoneMsg signals the stream is complete.
type StreamDoneMsg struct {
	Message map[string]interface{}
}

// StreamErrorMsg signals a stream error.
type StreamErrorMsg struct {
	Err error
}

// ToolCall represents a single tool call from the assistant.
type ToolCall struct {
	ID        string
	Name      string
	Arguments string
}

// ToolCallsMsg carries tool calls extracted from the stream.
type ToolCallsMsg struct {
	ToolCalls []ToolCall
}

// ToolProgressMsg reports progress on a running tool.
type ToolProgressMsg struct {
	ToolCallID string
	State      string
	Details    map[string]interface{}
}

// ToolDoneMsg signals a tool has finished.
type ToolDoneMsg struct {
	ToolCallID string
	Name       string
	Result     string
	Elapsed    time.Duration
	IsError    bool
}

// AskUserRequestMsg requests input from the user (e.g. choice selection).
type AskUserRequestMsg struct {
	Question   string
	Options    []string
	ResponseCh chan string
}

// AskUserResponseMsg carries the user's response.
type AskUserResponseMsg struct {
	Answer string
}

// MediaLoadedMsg signals that media has been downloaded/loaded.
type MediaLoadedMsg struct {
	URL       string
	MediaType string
	LocalPath string
	Err       error
}

// GenerationLoopMsg triggers the next generation iteration.
type GenerationLoopMsg struct{}

// SubmitMsg is sent when the user presses Enter in the input.
type SubmitMsg struct {
	Content string
}

// LoginResultMsg carries the result of a login attempt.
type LoginResultMsg struct {
	Key string
	Err error
}
