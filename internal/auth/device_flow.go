package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	auth0Domain  = "auth.fal.ai"
	auth0Client  = "TwXR51Vz8JbY8GUUMy6EyuVR0fTO7N4N"
	auth0Audience = "fal-cloud"
	auth0Scope   = "openid profile email offline_access"
	sessionSeedURL = "https://fal.ai/api/auth/cli/session-seed"
)

// DeviceCodeResponse from Auth0.
type DeviceCodeResponse struct {
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	Interval                int    `json:"interval"`
	ExpiresIn               int    `json:"expires_in"`
}

// TokenResponse from Auth0.
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// DeviceCodeLogin performs the Auth0 Device Code Flow.
// It opens the browser for the user to authenticate, then polls for the token.
// onStatus is called with status messages for UI updates.
func DeviceCodeLogin(onStatus func(string)) (*TokenResponse, error) {
	// Step 1: Request device code
	resp, err := http.PostForm(
		fmt.Sprintf("https://%s/oauth/device/code", auth0Domain),
		url.Values{
			"audience":  {auth0Audience},
			"client_id": {auth0Client},
			"scope":     {auth0Scope},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("device code request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("device code request failed (%d): %s", resp.StatusCode, body)
	}

	var dcResp DeviceCodeResponse
	if err := json.Unmarshal(body, &dcResp); err != nil {
		return nil, fmt.Errorf("failed to parse device code response: %w", err)
	}

	// Step 2: Build login URL and open browser
	loginURL := buildLoginURL(dcResp)
	if onStatus != nil {
		onStatus(fmt.Sprintf("Opening browser... Code: %s", dcResp.UserCode))
	}
	openBrowser(loginURL)

	// Step 3: Poll for token
	interval := dcResp.Interval
	if interval < 5 {
		interval = 5
	}
	deadline := time.Now().Add(time.Duration(dcResp.ExpiresIn) * time.Second)

	for time.Now().Before(deadline) {
		time.Sleep(time.Duration(interval) * time.Second)

		if onStatus != nil {
			onStatus("Waiting for browser login...")
		}

		token, err := pollToken(dcResp.DeviceCode)
		if err != nil {
			// Check for "authorization_pending" — keep polling
			if strings.Contains(err.Error(), "authorization_pending") {
				continue
			}
			if strings.Contains(err.Error(), "slow_down") {
				interval += 2
				continue
			}
			return nil, err
		}

		return token, nil
	}

	return nil, fmt.Errorf("login timed out — please try again")
}

func buildLoginURL(dcResp DeviceCodeResponse) string {
	// Use fal.ai session-seed endpoint for a cleaner login experience
	params := url.Values{
		"user_code":                   {dcResp.UserCode},
		"verification_uri_complete":   {dcResp.VerificationURIComplete},
	}
	return sessionSeedURL + "?" + params.Encode()
}

func pollToken(deviceCode string) (*TokenResponse, error) {
	resp, err := http.PostForm(
		fmt.Sprintf("https://%s/oauth/token", auth0Domain),
		url.Values{
			"grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
			"device_code": {deviceCode},
			"client_id":   {auth0Client},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("token poll failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		// Parse error response
		var errResp struct {
			Error       string `json:"error"`
			Description string `json:"error_description"`
		}
		if json.Unmarshal(body, &errResp) == nil && errResp.Error != "" {
			return nil, fmt.Errorf("%s: %s", errResp.Error, errResp.Description)
		}
		return nil, fmt.Errorf("token request failed (%d): %s", resp.StatusCode, body)
	}

	var token TokenResponse
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w", err)
	}
	return &token, nil
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		cmd = exec.Command("open", url)
	}
	_ = cmd.Start()
}
