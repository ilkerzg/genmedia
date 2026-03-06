package auth

import (
	"os"
	"path/filepath"
	"strings"
)

// GetCredentials returns the Authorization header value, checking all sources.
// Returns "" if not authenticated.
// Priority: FAL_KEY env → ~/.fal/auth0_token (OAuth) → ~/.cache/fal-dev/api_key (cached key)
func GetCredentials() string {
	// 1. FAL_KEY environment variable
	if key := os.Getenv("FAL_KEY"); key != "" {
		return formatKeyHeader(key)
	}

	// 2. OAuth tokens (~/.fal/auth0_token)
	refreshToken, accessToken, err := LoadTokens()
	if err == nil && refreshToken != "" {
		// Try cached access token first
		if accessToken != "" {
			return "Bearer " + accessToken
		}
		// Refresh to get a new access token
		newAccess, _, err := RefreshAccessToken(refreshToken)
		if err == nil && newAccess != "" {
			return "Bearer " + newAccess
		}
	}

	// 3. Cached API key (~/.cache/fal-dev/api_key)
	if key := readCachedKey(); key != "" {
		return formatKeyHeader(key)
	}

	return ""
}

// SaveAuthKey saves an API key to the cache file.
func SaveAuthKey(key string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dir := filepath.Join(home, ".cache", "fal-dev")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	// Strip any existing prefix before saving raw key
	raw := key
	raw = strings.TrimPrefix(raw, "Key ")
	raw = strings.TrimPrefix(raw, "Bearer ")
	return os.WriteFile(filepath.Join(dir, "api_key"), []byte(strings.TrimSpace(raw)+"\n"), 0600)
}

// ClearCachedKey removes the cached API key file.
func ClearCachedKey() {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}
	_ = os.Remove(filepath.Join(home, ".cache", "fal-dev", "api_key"))
}

// readCachedKey reads the cached API key from ~/.cache/fal-dev/api_key.
func readCachedKey() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	data, err := os.ReadFile(filepath.Join(home, ".cache", "fal-dev", "api_key"))
	if err != nil {
		return ""
	}
	key := strings.TrimSpace(string(data))
	if key == "" {
		return ""
	}
	return key
}

// formatKeyHeader formats a raw API key as a proper Authorization header value.
func formatKeyHeader(key string) string {
	key = strings.TrimSpace(key)
	// Already has a prefix
	if strings.HasPrefix(key, "Key ") || strings.HasPrefix(key, "Bearer ") {
		return key
	}
	// API key format: key_id:key_secret
	if strings.Contains(key, ":") {
		return "Key " + key
	}
	// JWT token (starts with ey)
	if strings.HasPrefix(key, "ey") {
		return "Bearer " + key
	}
	// Default to Key prefix
	return "Key " + key
}

