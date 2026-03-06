package tools

import (
	"fmt"
	"net/url"
)

type ListWorkflowsTool struct {
	AuthKey string
}

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
	params := url.Values{}
	params.Set("limit", "20")

	if q, ok := args["query"].(string); ok && q != "" {
		params.Set("search", q)
	}

	u := fmt.Sprintf("%s/workflows?%s", falAPIBase, params.Encode())
	result, err := apiGetRaw(u, t.AuthKey)
	if err != nil {
		return map[string]interface{}{"ok": false, "error": err.Error()}, nil
	}
	result["ok"] = true
	return result, nil
}
