package tools

import "fmt"

type WhoamiTool struct {
	AuthKey string
}

func (t *WhoamiTool) Name() string { return "whoami" }
func (t *WhoamiTool) Description() string {
	return "Show current authentication identity and key info."
}
func (t *WhoamiTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{},
	}
}

func (t *WhoamiTool) Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	u := fmt.Sprintf("%s/users/current", falRESTBase)
	result, err := apiGetRaw(u, t.AuthKey)
	if err != nil {
		return map[string]interface{}{"ok": false, "error": err.Error()}, nil
	}
	result["ok"] = true
	return result, nil
}
