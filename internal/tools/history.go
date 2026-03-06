package tools

import (
	"fmt"
	"time"
)

type RequestHistoryTool struct{}

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
	cmd := []string{"history", endpointID}
	if limit, ok := args["limit"].(float64); ok {
		cmd = append(cmd, "-n", fmt.Sprintf("%d", int(limit)))
	}
	return RunCLI(cmd, 120*time.Second), nil
}
