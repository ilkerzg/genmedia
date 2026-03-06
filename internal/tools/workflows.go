package tools

import "time"

type ListWorkflowsTool struct{}

func (t *ListWorkflowsTool) Name() string { return "list_workflows" }
func (t *ListWorkflowsTool) Description() string {
	return "List saved workflows."
}
func (t *ListWorkflowsTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"query": map[string]interface{}{
				"type":        "string",
				"description": "Optional search query to filter workflows",
			},
		},
	}
}

func (t *ListWorkflowsTool) Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	cmd := []string{"workflows"}
	if q, ok := args["query"].(string); ok && q != "" {
		cmd = append(cmd, q)
	}
	return RunCLI(cmd, 120*time.Second), nil
}
