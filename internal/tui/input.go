package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// InputModel wraps a textinput with history and paste support.
type InputModel struct {
	textInput textinput.Model
	history   []string
	histIdx   int    // -1 = not browsing history
	pasted    string // multi-line paste buffer
	width     int
	theme     *Theme
}

// NewInputModel creates a new InputModel.
func NewInputModel(theme *Theme) InputModel {
	ti := textinput.New()
	ti.Placeholder = "Describe what to generate, or ask anything..."
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(theme.TextMuted)
	ti.CharLimit = 0 // unlimited
	ti.Width = 80
	ti.PromptStyle = lipgloss.NewStyle().Foreground(theme.Accent)
	ti.TextStyle = lipgloss.NewStyle().Foreground(theme.Text)
	ti.Prompt = "❯ "
	ti.Focus()

	return InputModel{
		textInput: ti,
		histIdx:   -1,
		theme:     theme,
	}
}

// SetWidth sets the input width.
func (m *InputModel) SetWidth(w int) {
	m.width = w
	// textinput width = total - border(2) - padding(2) - prompt("❯ " = 2)
	m.textInput.Width = w - 6
	if m.textInput.Width < 10 {
		m.textInput.Width = 10
	}
}

// GetEffectiveValue returns pasted text if available, else input value.
func (m *InputModel) GetEffectiveValue() string {
	if m.pasted != "" {
		return m.pasted
	}
	return m.textInput.Value()
}

// Focus focuses the input.
func (m *InputModel) Focus() tea.Cmd {
	return m.textInput.Focus()
}

// Blur blurs the input.
func (m *InputModel) Blur() {
	m.textInput.Blur()
}

// Update handles input messages.
func (m InputModel) Update(msg tea.Msg) (InputModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			value := m.GetEffectiveValue()
			value = strings.TrimSpace(value)
			if value == "" {
				return m, nil
			}
			// Add to history
			m.history = append(m.history, value)
			m.histIdx = -1
			m.pasted = ""
			m.textInput.Reset()
			return m, func() tea.Msg {
				return SubmitMsg{Content: value}
			}

		case tea.KeyUp:
			if len(m.history) == 0 {
				break
			}
			if m.histIdx == -1 {
				m.histIdx = len(m.history) - 1
			} else if m.histIdx > 0 {
				m.histIdx--
			}
			m.textInput.SetValue(m.history[m.histIdx])
			m.textInput.CursorEnd()
			return m, nil

		case tea.KeyDown:
			if m.histIdx == -1 {
				break
			}
			if m.histIdx < len(m.history)-1 {
				m.histIdx++
				m.textInput.SetValue(m.history[m.histIdx])
			} else {
				m.histIdx = -1
				m.textInput.Reset()
			}
			m.textInput.CursorEnd()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View renders the input with a rounded border.
func (m InputModel) View() string {
	borderColor := m.theme.BorderDim
	if m.textInput.Focused() {
		borderColor = m.theme.Border
	}

	contentWidth := m.width - 4 // border(2) + padding(2)
	if contentWidth < 10 {
		contentWidth = 10
	}

	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Padding(0, 1).
		Width(contentWidth)

	return style.Render(m.textInput.View())
}
