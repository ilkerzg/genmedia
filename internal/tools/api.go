package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	falAPIBase  = "https://api.fal.ai/v1"
	falRESTBase = "https://rest.alpha.fal.ai"
)

var apiClient = &http.Client{Timeout: 30 * time.Second}

// apiGet performs an authenticated GET request and returns parsed JSON.
func apiGet(url, authKey string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}
	req.Header.Set("Authorization", authKey)

	resp, err := apiClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, truncStr(string(body), 500))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		// Try as array — some endpoints return JSON arrays
		var arr []interface{}
		if err2 := json.Unmarshal(body, &arr); err2 == nil {
			return map[string]interface{}{"ok": true, "data": arr}, nil
		}
		return nil, fmt.Errorf("invalid JSON response: %s", truncStr(string(body), 200))
	}
	return result, nil
}

// apiPost performs an authenticated POST request with JSON body and returns parsed JSON.
func apiPost(url, authKey string, payload interface{}) (map[string]interface{}, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonBody)))
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}
	req.Header.Set("Authorization", authKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := apiClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, truncStr(string(body), 500))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		var arr []interface{}
		if err2 := json.Unmarshal(body, &arr); err2 == nil {
			return map[string]interface{}{"ok": true, "data": arr}, nil
		}
		return nil, fmt.Errorf("invalid JSON response: %s", truncStr(string(body), 200))
	}
	return result, nil
}

// apiGetRaw performs an authenticated GET and returns the raw JSON (could be object or array).
func apiGetRaw(url, authKey string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}
	req.Header.Set("Authorization", authKey)

	resp, err := apiClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, truncStr(string(body), 500))
	}

	// Try object first
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err == nil {
		return result, nil
	}

	// Try array
	var arr []interface{}
	if err := json.Unmarshal(body, &arr); err == nil {
		return map[string]interface{}{"ok": true, "results": arr}, nil
	}

	return nil, fmt.Errorf("invalid JSON response: %s", truncStr(string(body), 200))
}
