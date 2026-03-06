package tools

import "time"

type GetAnalyticsTool struct{}

func (t *GetAnalyticsTool) Name() string { return "get_analytics" }
func (t *GetAnalyticsTool) Description() string {
	return "Get performance analytics for an endpoint including latency and throughput."
}
func (t *GetAnalyticsTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"endpoint_id": map[string]interface{}{
				"type":        "string",
				"description": "The model endpoint ID",
			},
		},
		"required": []string{"endpoint_id"},
	}
}

func (t *GetAnalyticsTool) Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	endpointID, _ := args["endpoint_id"].(string)
	return RunCLI([]string{"analytics", endpointID}, 120*time.Second), nil
}
