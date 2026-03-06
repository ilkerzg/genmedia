package tools

import "time"

type SearchModelsTool struct{}

func (t *SearchModelsTool) Name() string { return "search_models" }
func (t *SearchModelsTool) Description() string {
	return "Search for AI models on the fal platform. Returns a list of matching models with endpoint IDs, names, categories, and descriptions."
}
func (t *SearchModelsTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"query": map[string]interface{}{
				"type":        "string",
				"description": "Free-text search query (e.g. 'text to image', 'video generation', 'flux')",
			},
			"category": map[string]interface{}{
				"type":        "string",
				"description": "Filter by category (e.g. text-to-image, image-to-video, text-to-video, image-to-image)",
			},
			"sort": map[string]interface{}{
				"type":        "string",
				"enum":        []string{"newest", "oldest", "updated", "name"},
				"description": "Sort order for results",
			},
		},
	}
}

func (t *SearchModelsTool) Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	cmd := []string{"search"}
	if q, ok := args["query"].(string); ok && q != "" {
		cmd = append(cmd, q)
	}
	if c, ok := args["category"].(string); ok && c != "" {
		cmd = append(cmd, "-c", c)
	}
	if s, ok := args["sort"].(string); ok && s != "" {
		cmd = append(cmd, "-s", s)
	}
	cmd = append(cmd, "-n", "15")
	return RunCLI(cmd, 120*time.Second), nil
}
