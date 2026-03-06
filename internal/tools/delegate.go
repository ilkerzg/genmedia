package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

// FindFalDev locates the fal or fal-dev binary.
func FindFalDev() (string, error) {
	if path, err := exec.LookPath("fal"); err == nil {
		return path, nil
	}
	if path, err := exec.LookPath("fal-dev"); err == nil {
		return path, nil
	}
	return "", fmt.Errorf("fal CLI not found in PATH (install: pip install fal)")
}

// RunCLI runs a fal-dev CLI command in agent mode and returns parsed JSON.
func RunCLI(args []string, timeout time.Duration) map[string]interface{} {
	if timeout == 0 {
		timeout = 120 * time.Second
	}

	falDev, err := FindFalDev()
	if err != nil {
		return map[string]interface{}{"ok": false, "error": err.Error()}
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	fullArgs := append([]string{"--agent"}, args...)
	cmd := exec.CommandContext(ctx, falDev, fullArgs...)

	output, err := cmd.Output()
	if ctx.Err() == context.DeadlineExceeded {
		return map[string]interface{}{"ok": false, "error": fmt.Sprintf("Command timed out after %s", timeout)}
	}
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			stderr := string(exitErr.Stderr)
			if stderr != "" {
				return map[string]interface{}{"ok": false, "error": stderr}
			}
		}
		return map[string]interface{}{"ok": false, "error": err.Error()}
	}

	stdout := string(output)
	if stdout == "" {
		return map[string]interface{}{"ok": false, "error": "No output from command"}
	}

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(stdout), &result); err != nil {
		return map[string]interface{}{"ok": false, "error": fmt.Sprintf("Invalid JSON: %s", stdout[:min(200, len(stdout))])}
	}
	return result
}
