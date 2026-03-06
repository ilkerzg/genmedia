package tools

import (
	"fmt"
	"net/url"
)

type SearchModelsTool struct {
	AuthKey string
}

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
	params := url.Values{}
	params.Set("limit", "15")

	if q, ok := args["query"].(string); ok && q != "" {
		params.Set("q", q)
	}
	if c, ok := args["category"].(string); ok && c != "" {
		params.Set("category", c)
	}
	if s, ok := args["sort"].(string); ok && s != "" {
		params.Set("sort", s)
	}

	u := fmt.Sprintf("%s/models?%s", falAPIBase, params.Encode())
	result, err := apiGetRaw(u, t.AuthKey)
	if err != nil {
		return map[string]interface{}{"ok": false, "error": err.Error()}, nil
	}
	result["ok"] = true
	return result, nil
}
