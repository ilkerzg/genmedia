package tools

import (
	"fmt"
	"net/url"
)

type RequestHistoryTool struct {
	AuthKey string
}

func (t *RequestHistoryTool) Name() string { return "request_history" }
func (t *RequestHistoryTool) Description() string {
	return "Get request history for an endpoint."
}
func (t *RequestHistoryTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"endpoint_id": map[string]interface{}{
				"type":        "string",
				"description": "The model endpoint ID",
			},
			"limit": map[string]interface{}{
				"type":        "integer",
				"description": "Maximum number of results",
			},
		},
		"required": []string{"endpoint_id"},
	}
}

func (t *RequestHistoryTool) Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	endpointID, _ := args["endpoint_id"].(string)

	params := url.Values{}
	params.Set("endpoint_id", endpointID)

	limit := 20
	if l, ok := args["limit"].(float64); ok {
		limit = int(l)
	}
	params.Set("limit", fmt.Sprintf("%d", limit))

	u := fmt.Sprintf("%s/models/requests/by-endpoint?%s", falAPIBase, params.Encode())
	result, err := apiGetRaw(u, t.AuthKey)
	if err != nil {
		return map[string]interface{}{"ok": false, "error": err.Error()}, nil
	}
	result["ok"] = true
	return result, nil
}
