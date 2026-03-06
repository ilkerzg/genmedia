package tools

import "time"

type ListFilesTool struct{}

func (t *ListFilesTool) Name() string { return "list_files" }
func (t *ListFilesTool) Description() string {
	return "List files in fal storage."
}
func (t *ListFilesTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"path": map[string]interface{}{
				"type":        "string",
				"description": "Path to list files from",
			},
		},
	}
}

func (t *ListFilesTool) Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	cmd := []string{"files"}
	if p, ok := args["path"].(string); ok && p != "" {
		cmd = append(cmd, p)
	}
	return RunCLI(cmd, 120*time.Second), nil
}
