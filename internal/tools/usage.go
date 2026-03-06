package tools

import "time"

type CheckUsageTool struct{}

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
	cmd := []string{"usage"}
	if period, ok := args["period"].(string); ok && period != "" {
		cmd = append(cmd, "--period", period)
	}
	return RunCLI(cmd, 120*time.Second), nil
}
