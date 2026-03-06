package auth

import (
	"fmt"
	"os/exec"
)

// RunFalLogin runs `fal auth login` as a subprocess.
// Returns nil on success, error on failure.
func RunFalLogin() error {
	path, err := FindFalBinary()
	if err != nil {
		return fmt.Errorf("fal CLI not found. Install it: pip install fal")
	}

	cmd := exec.Command(path, "auth", "login")
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

// FindFalBinary locates the fal or fal-dev binary in PATH.
func FindFalBinary() (string, error) {
	// Try fal first, then fal-dev
	if path, err := exec.LookPath("fal"); err == nil {
		return path, nil
	}
	if path, err := exec.LookPath("fal-dev"); err == nil {
		return path, nil
	}
	return "", fmt.Errorf("not found")
}
