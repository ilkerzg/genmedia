package tools

import (
	"fmt"
	"net/url"
)

type GetPricingTool struct {
	AuthKey string
}

func (t *GetPricingTool) Name() string { return "get_pricing" }
func (t *GetPricingTool) Description() string {
	return "Get pricing information for a model endpoint."
}
func (t *GetPricingTool) Parameters() map[string]interface{} {
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

func (t *GetPricingTool) Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	endpointID, _ := args["endpoint_id"].(string)

	params := url.Values{}
	params.Set("endpoint_id", endpointID)

	u := fmt.Sprintf("%s/models/pricing?%s", falAPIBase, params.Encode())
	result, err := apiGetRaw(u, t.AuthKey)
	if err != nil {
		return map[string]interface{}{"ok": false, "error": err.Error()}, nil
	}
	result["ok"] = true
	return result, nil
}
