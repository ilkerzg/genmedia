package tools

import (
	"fmt"
	"net/url"
	"time"
)

type CheckUsageTool struct {
	AuthKey string
}

func (t *CheckUsageTool) Name() string { return "check_usage" }
func (t *CheckUsageTool) Description() string {
	return "Check current usage and cost information."
}
func (t *CheckUsageTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"period": map[string]interface{}{
				"type":        "string",
				"description": "Time period (e.g. 'today', 'week', 'month')",
			},
		},
	}
}

func (t *CheckUsageTool) Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	params := url.Values{}
	params.Add("expand[]", "summary")

	// Calculate start/end based on period
	now := time.Now().UTC()
	end := now.Format(time.RFC3339)
	var start string

	period, _ := args["period"].(string)
	switch period {
	case "today":
		start = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
	case "week":
		start = now.AddDate(0, 0, -7).Format(time.RFC3339)
	default: // "month" or empty
		start = now.AddDate(0, -1, 0).Format(time.RFC3339)
	}

	params.Set("start", start)
	params.Set("end", end)

	u := fmt.Sprintf("%s/models/usage?%s", falAPIBase, params.Encode())
	result, err := apiGetRaw(u, t.AuthKey)
	if err != nil {
		return map[string]interface{}{"ok": false, "error": err.Error()}, nil
	}
	result["ok"] = true
	return result, nil
}
