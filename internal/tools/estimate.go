package tools

import (
	"fmt"
	"time"
)

type EstimateCostTool struct{}

func (t *EstimateCostTool) Name() string { return "estimate_cost" }
func (t *EstimateCostTool) Description() string {
	return "Estimate the cost of running a model with given parameters."
}
func (t *EstimateCostTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"endpoint_id": map[string]interface{}{
				"type":        "string",
				"description": "The model endpoint ID",
			},
			"quantity": map[string]interface{}{
				"type":        "integer",
				"description": "Number of runs to estimate",
			},
		},
		"required": []string{"endpoint_id"},
	}
}

func (t *EstimateCostTool) Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	endpointID, _ := args["endpoint_id"].(string)
	cmd := []string{"estimate", endpointID}
	if qty, ok := args["quantity"].(float64); ok {
		cmd = append(cmd, "--qty", fmt.Sprintf("%d", int(qty)))
	}
	return RunCLI(cmd, 120*time.Second), nil
}
