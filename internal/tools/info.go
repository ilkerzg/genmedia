package tools

import (
	"fmt"
	"net/url"
)

// ModelInfoTool gets detailed model information including input/output schemas.
type ModelInfoTool struct {
	AuthKey string
}

func (t *ModelInfoTool) Name() string { return "model_info" }
func (t *ModelInfoTool) Description() string {
	return "Get detailed information about a specific model, including its input parameters (with types, defaults, constraints) and output schema."
}
func (t *ModelInfoTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"endpoint_id": map[string]interface{}{
				"type":        "string",
				"description": "The model endpoint ID (e.g. 'fal-ai/flux/dev', 'fal-ai/kling-video/v2.5/standard/text-to-video')",
			},
		},
		"required": []string{"endpoint_id"},
	}
}

func (t *ModelInfoTool) Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	endpointID, _ := args["endpoint_id"].(string)

	params := url.Values{}
	params.Set("endpoint_id", endpointID)
	params.Set("expand", "openapi-3.0")

	u := fmt.Sprintf("%s/models?%s", falAPIBase, params.Encode())
	result, err := apiGetRaw(u, t.AuthKey)
	if err != nil {
		return map[string]interface{}{"ok": false, "error": err.Error()}, nil
	}

	// Resolve $ref in schemas
	for _, key := range []string{"input_schema", "output_schema"} {
		if schema, ok := result[key].(map[string]interface{}); ok {
			result[key] = resolveRefs(schema, nil)
		}
	}
	// Remove top-level definition containers
	delete(result, "components")
	delete(result, "definitions")
	delete(result, "$defs")

	result["ok"] = true
	return result, nil
}

// resolveRefs recursively resolves $ref in JSON Schema.
func resolveRefs(obj interface{}, definitions map[string]interface{}) interface{} {
	if definitions == nil {
		definitions = make(map[string]interface{})
	}

	switch v := obj.(type) {
	case map[string]interface{}:
		// Collect definitions
		for _, key := range []string{"definitions", "$defs", "components"} {
			if defs, ok := v[key].(map[string]interface{}); ok {
				d := defs
				if key == "components" {
					if schemas, ok := d["schemas"].(map[string]interface{}); ok {
						d = schemas
					}
				}
				for k, val := range d {
					definitions[k] = val
				}
			}
		}

		if ref, ok := v["$ref"].(string); ok {
			// Extract name from "#/definitions/Foo" or "#/components/schemas/Foo"
			parts := splitLast(ref, "/")
			refName := parts
			if resolved, ok := definitions[refName].(map[string]interface{}); ok {
				merged := make(map[string]interface{})
				for k, val := range resolved {
					merged[k] = val
				}
				for k, val := range v {
					if k != "$ref" {
						merged[k] = val
					}
				}
				return resolveRefs(merged, definitions)
			}
			// Can't resolve — generic object
			result := map[string]interface{}{"type": "object", "description": "(schema: " + refName + ")"}
			for k, val := range v {
				if k != "$ref" {
					result[k] = val
				}
			}
			return result
		}

		result := make(map[string]interface{})
		for k, val := range v {
			if k != "definitions" && k != "$defs" {
				result[k] = resolveRefs(val, definitions)
			}
		}
		return result

	case []interface{}:
		result := make([]interface{}, len(v))
		for i, item := range v {
			result[i] = resolveRefs(item, definitions)
		}
		return result
	}

	return obj
}

func splitLast(s, sep string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == sep {
			return s[i+1:]
		}
	}
	return s
}
