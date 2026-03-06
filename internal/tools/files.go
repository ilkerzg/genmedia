package tools

import "fmt"

type ListFilesTool struct {
	AuthKey string
}

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
	path := ""
	if p, ok := args["path"].(string); ok && p != "" {
		path = p
	}

	u := fmt.Sprintf("%s/files/list/%s", falRESTBase, path)
	result, err := apiGetRaw(u, t.AuthKey)
	if err != nil {
		return map[string]interface{}{"ok": false, "error": err.Error()}, nil
	}
	result["ok"] = true
	return result, nil
}
