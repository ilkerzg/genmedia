package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var KnownCategories = []string{
	"text-to-image",
	"image-to-image",
	"text-to-video",
	"image-to-video",
	"text-to-audio",
	"text-to-speech",
}

type Preferences struct {
	Defaults map[string]string `json:"defaults"`
	path     string
}

func preferencesPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "fal-dev", "preferences.json")
}

func LoadPreferences() (*Preferences, error) {
	p := &Preferences{
		Defaults: make(map[string]string),
		path:     preferencesPath(),
	}

	data, err := os.ReadFile(p.path)
	if err != nil {
		if os.IsNotExist(err) {
			return p, nil
		}
		return p, err
	}

	if err := json.Unmarshal(data, p); err != nil {
		return p, err
	}
	if p.Defaults == nil {
		p.Defaults = make(map[string]string)
	}
	return p, nil
}

func (p *Preferences) Save() error {
	dir := filepath.Dir(p.path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p.path, data, 0600)
}

func (p *Preferences) GetDefaults() map[string]string {
	return p.Defaults
}

func (p *Preferences) GetDefault(category string) string {
	return p.Defaults[category]
}

func (p *Preferences) SetDefault(category, model string) {
	p.Defaults[category] = model
}

func (p *Preferences) FormatForSystemPrompt() string {
	if len(p.Defaults) == 0 {
		return ""
	}
	var lines []string
	for category, model := range p.Defaults {
		lines = append(lines, fmt.Sprintf("- %s: %s", category, model))
	}
	return "When the user asks to generate without specifying a model, use these defaults:\n" + strings.Join(lines, "\n")
}
