package tools

import "fmt"

type EstimateCostTool struct {
	AuthKey string
}

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

	payload := map[string]interface{}{
		"endpoint_id": endpointID,
	}
	if qty, ok := args["quantity"].(float64); ok {
		payload["quantity"] = int(qty)
	}

	u := fmt.Sprintf("%s/models/pricing/estimate", falAPIBase)
	result, err := apiPost(u, t.AuthKey, payload)
	if err != nil {
		return map[string]interface{}{"ok": false, "error": err.Error()}, nil
	}
	result["ok"] = true
	return result, nil
}
