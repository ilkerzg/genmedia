package tools

import "time"

type WhoamiTool struct{}

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
	return RunCLI([]string{"whoami"}, 120*time.Second), nil
}
