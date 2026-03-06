package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/fal-ai/genmedia/internal/config"
)

// ThickBar is the thick vertical bar used for left borders.
const ThickBar = "\u2503" // ┃

// Theme holds all colors for the TUI.
type Theme struct {
	Name           string
	Primary        lipgloss.Color
	Secondary      lipgloss.Color
	Accent         lipgloss.Color
	Text           lipgloss.Color
	TextMuted      lipgloss.Color
	TextEmphasized lipgloss.Color
	Bg             lipgloss.Color
	BgSecondary    lipgloss.Color
	BgDarker       lipgloss.Color
	Border         lipgloss.Color
	BorderFocused  lipgloss.Color
	BorderDim      lipgloss.Color
	Success        lipgloss.Color
	Warning        lipgloss.Color
	Error          lipgloss.Color
	Info           lipgloss.Color
}

// themeCache stores converted themes.
var themeCache = map[string]*Theme{}

// convertTheme converts a config.ThemeConfig to a Theme.
func convertTheme(name string, tc config.ThemeConfig) *Theme {
	return &Theme{
		Name:           name,
		Primary:        lipgloss.Color(tc.Primary),
		Secondary:      lipgloss.Color(tc.Secondary),
		Accent:         lipgloss.Color(tc.Accent),
		Text:           lipgloss.Color(tc.Text),
		TextMuted:      lipgloss.Color(tc.TextMuted),
		TextEmphasized: lipgloss.Color(tc.TextEmphasized),
		Bg:             lipgloss.Color(tc.Bg),
		BgSecondary:    lipgloss.Color(tc.BgSecondary),
		BgDarker:       lipgloss.Color(tc.BgDarker),
		Border:         lipgloss.Color(tc.Border),
		BorderFocused:  lipgloss.Color(tc.BorderFocused),
		BorderDim:      lipgloss.Color(tc.BorderDim),
		Success:        lipgloss.Color(tc.Success),
		Warning:        lipgloss.Color(tc.Warning),
		Error:          lipgloss.Color(tc.Error),
		Info:           lipgloss.Color(tc.Info),
	}
}

// GetTheme returns the theme for the given key. Falls back to "tokyonight".
func GetTheme(key string) *Theme {
	if t, ok := themeCache[key]; ok {
		return t
	}
	tc, ok := config.Themes[key]
	if !ok {
		tc = config.Themes["tokyonight"]
		key = "tokyonight"
	}
	t := convertTheme(key, tc)
	themeCache[key] = t
	return t
}

// Style helpers

func (t *Theme) UserBorderStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(t.Secondary)
}

func (t *Theme) AssistantBorderStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(t.Accent)
}

func (t *Theme) ToolBorderStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(t.TextMuted)
}

func (t *Theme) SystemBorderStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(t.BorderDim)
}

func (t *Theme) StatusBarStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(t.BgSecondary).
		Foreground(t.TextMuted)
}

func (t *Theme) InputStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Text)
}
