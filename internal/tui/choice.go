package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ChoiceModel is the choice picker overlay for ask_user and model selection.
type ChoiceModel struct {
	visible    bool
	question   string
	options    []string
	selected   int
	customMode bool // "Other..." custom text mode
	customText string
	responseCh chan string // channel to send response back to tool
	theme      *Theme
	width      int
}

// NewChoiceModel creates a new choice picker.
func NewChoiceModel(theme *Theme) *ChoiceModel {
	return &ChoiceModel{
		theme: theme,
	}
}

// Show shows the choice picker with a question and options.
func (m *ChoiceModel) Show(question string, options []string, responseCh chan string) {
	m.visible = true
	m.question = question
	m.options = options
	m.selected = 0
	m.customMode = false
	m.customText = ""
	m.responseCh = responseCh
}

// Hide hides the picker.
func (m *ChoiceModel) Hide() {
	m.visible = false
	m.question = ""
	m.options = nil
	m.responseCh = nil
	m.customMode = false
	m.customText = ""
}

// IsVisible returns whether the picker is shown.
func (m *ChoiceModel) IsVisible() bool {
	return m.visible
}

// SetWidth sets the picker width.
func (m *ChoiceModel) SetWidth(w int) {
	m.width = w
}

// Update handles keys when the picker is visible.
func (m *ChoiceModel) Update(msg tea.KeyMsg) tea.Cmd {
	if m.customMode {
		return m.updateCustomMode(msg)
	}

	totalItems := len(m.options) + 1 // +1 for "Other..."

	switch msg.Type {
	case tea.KeyUp:
		if m.selected > 0 {
			m.selected--
		}
	case tea.KeyDown:
		if m.selected < totalItems-1 {
			m.selected++
		}
	case tea.KeyEnter:
		if m.selected == len(m.options) {
			// "Other..." selected
			m.customMode = true
			m.customText = ""
			return nil
		}
		answer := m.options[m.selected]
		ch := m.responseCh
		m.Hide()
		if ch != nil {
			ch <- answer
		}
		return func() tea.Msg {
			return AskUserResponseMsg{Answer: answer}
		}
	case tea.KeyEscape:
		ch := m.responseCh
		m.Hide()
		if ch != nil {
			ch <- ""
		}
		return func() tea.Msg {
			return AskUserResponseMsg{Answer: ""}
		}
	}
	return nil
}

func (m *ChoiceModel) updateCustomMode(msg tea.KeyMsg) tea.Cmd {
	switch msg.Type {
	case tea.KeyEnter:
		answer := strings.TrimSpace(m.customText)
		if answer == "" {
			return nil
		}
		ch := m.responseCh
		m.Hide()
		if ch != nil {
			ch <- answer
		}
		return func() tea.Msg {
			return AskUserResponseMsg{Answer: answer}
		}
	case tea.KeyEscape:
		m.customMode = false
		return nil
	case tea.KeyBackspace:
		if len(m.customText) > 0 {
			runes := []rune(m.customText)
			m.customText = string(runes[:len(runes)-1])
		}
	default:
		if msg.Type == tea.KeyRunes {
			m.customText += string(msg.Runes)
		} else if msg.Type == tea.KeySpace {
			m.customText += " "
		}
	}
	return nil
}

// View renders the choice picker.
func (m *ChoiceModel) View() string {
	if !m.visible {
		return ""
	}

	var sb strings.Builder

	// Question
	qStyle := lipgloss.NewStyle().Foreground(m.theme.Accent).Bold(true)
	sb.WriteString(qStyle.Render("? " + m.question))
	sb.WriteByte('\n')

	// Options
	for i, opt := range m.options {
		prefix := "  "
		style := lipgloss.NewStyle().Foreground(m.theme.Text)
		if i == m.selected && !m.customMode {
			prefix = "> "
			style = lipgloss.NewStyle().Foreground(m.theme.Primary).Bold(true)
		}
		sb.WriteString(prefix + style.Render(opt))
		sb.WriteByte('\n')
	}

	// "Other..." option
	otherPrefix := "  "
	otherStyle := lipgloss.NewStyle().Foreground(m.theme.TextMuted)
	if m.selected == len(m.options) && !m.customMode {
		otherPrefix = "> "
		otherStyle = lipgloss.NewStyle().Foreground(m.theme.Primary).Bold(true)
	}
	sb.WriteString(otherPrefix + otherStyle.Render("Other..."))

	if m.customMode {
		sb.WriteByte('\n')
		prompt := lipgloss.NewStyle().Foreground(m.theme.Primary).Render("  > ")
		cursor := lipgloss.NewStyle().Foreground(m.theme.Accent).Render("█")
		sb.WriteString(prompt + m.customText + cursor)
	}

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.theme.Border).
		Padding(0, 1).
		Width(m.width - 4).
		Render(sb.String())
}
