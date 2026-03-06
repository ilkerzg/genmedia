package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// tokenDir returns the fal auth directory (~/.fal).
func tokenDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".fal"), nil
}

// tokenPath returns the path to the auth0 token file.
func tokenPath() (string, error) {
	dir, err := tokenDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "auth0_token"), nil
}

// SaveTokens saves refresh and access tokens to ~/.fal/auth0_token.
// Format: first line = refresh_token, second line = access_token (optional).
func SaveTokens(refreshToken, accessToken string) error {
	dir, err := tokenDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	path, _ := tokenPath()
	content := refreshToken
	if accessToken != "" {
		content += "\n" + accessToken
	}
	return os.WriteFile(path, []byte(content+"\n"), 0600)
}

// LoadTokens loads refresh and access tokens from ~/.fal/auth0_token.
func LoadTokens() (refreshToken, accessToken string, err error) {
	path, err := tokenPath()
	if err != nil {
		return "", "", err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) >= 1 {
		refreshToken = strings.TrimSpace(lines[0])
	}
	if len(lines) >= 2 {
		accessToken = strings.TrimSpace(lines[1])
	}
	return
}

// RefreshAccessToken uses a refresh token to get a new access token from Auth0.
func RefreshAccessToken(refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	resp, err := http.PostForm(
		fmt.Sprintf("https://%s/oauth/token", auth0Domain),
		url.Values{
			"grant_type":    {"refresh_token"},
			"client_id":     {auth0Client},
			"refresh_token": {refreshToken},
		},
	)
	if err != nil {
		return "", "", fmt.Errorf("refresh request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		// Token might be revoked — clear local tokens
		_ = ClearTokens()
		return "", "", fmt.Errorf("token refresh failed (%d): %s", resp.StatusCode, body)
	}

	var token struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.Unmarshal(body, &token); err != nil {
		return "", "", fmt.Errorf("failed to parse refresh response: %w", err)
	}

	// Auth0 Refresh Token Rotation: new refresh token replaces old one
	newRefresh := refreshToken
	if token.RefreshToken != "" {
		newRefresh = token.RefreshToken
	}

	// Save updated tokens
	_ = SaveTokens(newRefresh, token.AccessToken)

	return token.AccessToken, newRefresh, nil
}

// ClearTokens removes the stored auth tokens.
func ClearTokens() error {
	path, err := tokenPath()
	if err != nil {
		return err
	}
	return os.Remove(path)
}
