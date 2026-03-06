package tools

import "time"

// AskUserTool lets the AI ask the user interactive questions via a choice picker.
type AskUserTool struct{}

func (t *AskUserTool) Name() string { return "ask_user" }
func (t *AskUserTool) Description() string {
	return "MANDATORY for asking questions. Shows an interactive picker UI with arrow-key navigation. " +
		"NEVER write questions as text — ALWAYS use this tool instead. " +
		"Use for: choosing between models, styles, resolutions, creative directions, or any clarification."
}
func (t *AskUserTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"question": map[string]interface{}{
				"type":        "string",
				"description": "The question to ask the user",
			},
			"options": map[string]interface{}{
				"type":        "array",
				"items":       map[string]interface{}{"type": "string"},
				"description": "List of options for the user to choose from (2-6 options)",
			},
		},
		"required": []string{"question", "options"},
	}
}

func (t *AskUserTool) Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	question, _ := args["question"].(string)
	optionsRaw, _ := args["options"].([]interface{})

	if question == "" || len(optionsRaw) == 0 {
		return map[string]interface{}{"ok": false, "error": "question and options are required"}, nil
	}

	if onProgress == nil {
		return map[string]interface{}{"ok": false, "error": "ask_user requires UI context (on_progress callback)"}, nil
	}

	options := make([]string, len(optionsRaw))
	for i, o := range optionsRaw {
		options[i], _ = o.(string)
	}

	// Create response channel — the TUI will send the answer on this channel
	responseCh := make(chan string, 1)

	onProgress(map[string]interface{}{
		"type":        "ask_user",
		"question":    question,
		"options":     options,
		"response_ch": responseCh,
	})

	// Block until the user picks an option (5 min timeout)
	select {
	case answer := <-responseCh:
		if answer == "" {
			return map[string]interface{}{"ok": false, "error": "User cancelled the question"}, nil
		}
		return map[string]interface{}{"ok": true, "answer": answer}, nil
	case <-time.After(5 * time.Minute):
		return map[string]interface{}{"ok": false, "error": "User did not respond within 5 minutes."}, nil
	}
}
