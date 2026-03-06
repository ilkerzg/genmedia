package tools

import (
	"encoding/json"
	"fmt"
)

// Tool is the interface all tools must implement.
type Tool interface {
	Name() string
	Description() string
	Parameters() map[string]interface{}
	Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error)
}

// Registry holds registered tools.
type Registry struct {
	tools map[string]Tool
}

// NewRegistry creates a registry with all tools registered.
func NewRegistry(authKey string) *Registry {
	r := &Registry{tools: make(map[string]Tool)}
	r.Register(&GenerateTool{AuthKey: authKey})
	r.Register(&AskUserTool{})
	r.Register(&GetSkillTool{})
	r.Register(&ModelInfoTool{AuthKey: authKey})
	r.Register(&SearchModelsTool{AuthKey: authKey})
	r.Register(&GetPricingTool{AuthKey: authKey})
	r.Register(&EstimateCostTool{AuthKey: authKey})
	r.Register(&CheckUsageTool{AuthKey: authKey})
	r.Register(&GetAnalyticsTool{AuthKey: authKey})
	r.Register(&RequestHistoryTool{AuthKey: authKey})
	r.Register(&ListWorkflowsTool{AuthKey: authKey})
	r.Register(&ListFilesTool{AuthKey: authKey})
	r.Register(&WhoamiTool{AuthKey: authKey})
	return r
}

// Register adds a tool.
func (r *Registry) Register(t Tool) {
	r.tools[t.Name()] = t
}

// Get returns a tool by name.
func (r *Registry) Get(name string) Tool {
	return r.tools[name]
}

// Execute runs a tool by name and returns JSON string result.
func (r *Registry) Execute(name string, arguments interface{}, onProgress func(map[string]interface{})) string {
	tool := r.tools[name]
	if tool == nil {
		b, _ := json.Marshal(map[string]interface{}{"ok": false, "error": fmt.Sprintf("Unknown tool: %s", name)})
		return string(b)
	}

	var args map[string]interface{}
	switch v := arguments.(type) {
	case map[string]interface{}:
		args = v
	case string:
		if err := json.Unmarshal([]byte(v), &args); err != nil {
			b, _ := json.Marshal(map[string]interface{}{"ok": false, "error": fmt.Sprintf("Invalid JSON args: %v", err)})
			return string(b)
		}
	default:
		b, _ := json.Marshal(map[string]interface{}{"ok": false, "error": "Invalid arguments type"})
		return string(b)
	}

	data, err := tool.Execute(args, onProgress)
	if err != nil {
		b, _ := json.Marshal(map[string]interface{}{"ok": false, "error": err.Error()})
		return string(b)
	}

	result, err := json.Marshal(data)
	if err != nil {
		b, _ := json.Marshal(map[string]interface{}{"ok": false, "error": fmt.Sprintf("JSON marshal error: %v", err)})
		return string(b)
	}

	// Truncate very large results
	const maxLen = 32000
	if len(result) > maxLen {
		truncated := string(result[:maxLen])
		ok := true
		if v, exists := data["ok"]; exists {
			if b, isBool := v.(bool); isBool {
				ok = b
			}
		}
		wrapped, _ := json.Marshal(map[string]interface{}{
			"ok":             ok,
			"data_truncated": true,
			"partial":        truncated,
			"note":           fmt.Sprintf("Result truncated from %d chars to %d", len(result), maxLen),
		})
		return string(wrapped)
	}

	return string(result)
}

// OpenAISchemas returns tool schemas in OpenAI function calling format.
func (r *Registry) OpenAISchemas() []map[string]interface{} {
	schemas := make([]map[string]interface{}, 0, len(r.tools))
	for _, t := range r.tools {
		schemas = append(schemas, map[string]interface{}{
			"type": "function",
			"function": map[string]interface{}{
				"name":        t.Name(),
				"description": t.Description(),
				"parameters": mergeMap(map[string]interface{}{
					"type": "object",
				}, t.Parameters()),
			},
		})
	}
	return schemas
}

func mergeMap(base, overlay map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range base {
		result[k] = v
	}
	for k, v := range overlay {
		result[k] = v
	}
	return result
}
