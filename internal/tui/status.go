package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// StatusModel displays a single-line status bar.
type StatusModel struct {
	tagline    string
	model      string
	generating bool
	width      int
	theme      *Theme
}

// NewStatusModel creates a new StatusModel.
func NewStatusModel(model string, theme *Theme) StatusModel {
	return StatusModel{
		model: model,
		theme: theme,
	}
}

// SetWidth sets the status bar width.
func (s *StatusModel) SetWidth(w int) {
	s.width = w
}

// SetModel updates the displayed model name.
func (s *StatusModel) SetModel(m string) {
	s.model = m
}

// SetGenerating sets the generating state and tagline.
func (s *StatusModel) SetGenerating(g bool, tagline string) {
	s.generating = g
	s.tagline = tagline
}

// View renders the status bar.
func (s StatusModel) View() string {
	// Use Border color as status bar bg — distinct from terminal background
	barBgColor := s.theme.BgDarker

	brandStyle := lipgloss.NewStyle().
		Background(s.theme.Primary).
		Foreground(s.theme.BgDarker).
		Bold(true).
		Padding(0, 1)

	tagStyle := lipgloss.NewStyle().
		Background(barBgColor).
		Foreground(s.theme.TextMuted).
		Italic(true).
		Padding(0, 1)

	modelStyle := lipgloss.NewStyle().
		Background(barBgColor).
		Foreground(s.theme.Text).
		Padding(0, 1)

	helpStyle := lipgloss.NewStyle().
		Background(barBgColor).
		Foreground(s.theme.TextMuted).
		Padding(0, 1)

	barBg := lipgloss.NewStyle().
		Background(barBgColor).
		Width(s.width)

	brand := brandStyle.Render(" genmedia ")

	tag := ""
	if s.tagline != "" {
		tag = tagStyle.Render(s.tagline)
	} else if s.generating {
		tag = tagStyle.Render("generating...")
	}

	model := modelStyle.Render(s.model)
	help := helpStyle.Render("/help  Ctrl+C")

	left := brand + tag + model
	right := help

	// Calculate gap
	leftWidth := lipgloss.Width(left)
	rightWidth := lipgloss.Width(right)
	gap := s.width - leftWidth - rightWidth
	if gap < 0 {
		gap = 0
	}
	gapStr := lipgloss.NewStyle().
		Background(barBgColor).
		Render(spaces(gap))

	return barBg.Render(left + gapStr + right)
}

func spaces(n int) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = ' '
	}
	return string(b)
}
