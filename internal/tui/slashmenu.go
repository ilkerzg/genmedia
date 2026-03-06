package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fal-ai/genmedia/internal/commands"
)

// SlashMenuModel is the autocomplete overlay for slash commands.
type SlashMenuModel struct {
	visible  bool
	filter   string
	items    []commands.Command
	filtered []commands.Command
	selected int
	theme    *Theme
	width    int
}

// NewSlashMenuModel creates a new slash menu.
func NewSlashMenuModel(cmdRegistry *commands.Registry, theme *Theme) *SlashMenuModel {
	return &SlashMenuModel{
		items: cmdRegistry.AllUnique(),
		theme: theme,
	}
}

// Show shows the menu with a filter string.
func (m *SlashMenuModel) Show(filter string) {
	m.filter = filter
	m.visible = true
	m.selected = 0
	m.updateFiltered()
}

// Hide hides the menu.
func (m *SlashMenuModel) Hide() {
	m.visible = false
	m.filter = ""
	m.filtered = nil
}

// IsVisible returns whether the menu is shown.
func (m *SlashMenuModel) IsVisible() bool {
	return m.visible
}

// SelectedCommand returns the currently selected command name, if any.
func (m *SlashMenuModel) SelectedCommand() string {
	if len(m.filtered) == 0 || m.selected >= len(m.filtered) {
		return ""
	}
	return m.filtered[m.selected].Name()
}

// SetWidth sets the menu width.
func (m *SlashMenuModel) SetWidth(w int) {
	m.width = w
}

// Update handles keys when the menu is visible.
func (m *SlashMenuModel) Update(msg tea.KeyMsg) tea.Cmd {
	switch msg.Type {
	case tea.KeyUp:
		if m.selected > 0 {
			m.selected--
		}
	case tea.KeyDown:
		if m.selected < len(m.filtered)-1 {
			m.selected++
		}
	case tea.KeyTab:
		// Complete with selected command
		if cmd := m.SelectedCommand(); cmd != "" {
			m.Hide()
			return func() tea.Msg {
				return SlashCompleteMsg{Command: cmd}
			}
		}
	case tea.KeyEscape:
		m.Hide()
	}
	return nil
}

// UpdateFilter updates the filter and re-filters items.
func (m *SlashMenuModel) UpdateFilter(filter string) {
	m.filter = filter
	m.selected = 0
	m.updateFiltered()
	if len(m.filtered) == 0 {
		m.visible = false
	}
}

func (m *SlashMenuModel) updateFiltered() {
	m.filtered = nil
	prefix := strings.ToLower(m.filter)
	for _, cmd := range m.items {
		if strings.HasPrefix(strings.ToLower(cmd.Name()), prefix) {
			m.filtered = append(m.filtered, cmd)
		}
	}
}

// View renders the slash menu.
func (m *SlashMenuModel) View() string {
	if !m.visible || len(m.filtered) == 0 {
		return ""
	}

	var sb strings.Builder
	for i, cmd := range m.filtered {
		nameStyle := lipgloss.NewStyle().Foreground(m.theme.Primary)
		descStyle := lipgloss.NewStyle().Foreground(m.theme.TextMuted)

		line := "  " + nameStyle.Render(cmd.Name())
		if cmd.ArgsHint() != "" {
			line += " " + lipgloss.NewStyle().Foreground(m.theme.TextMuted).Render(cmd.ArgsHint())
		}
		line += "  " + descStyle.Render(cmd.Description())

		if i == m.selected {
			line = lipgloss.NewStyle().
				Background(m.theme.BgSecondary).
				Render(line)
		}

		sb.WriteString(line)
		if i < len(m.filtered)-1 {
			sb.WriteByte('\n')
		}
	}

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.theme.Border).
		Padding(0, 1).
		Width(m.width - 4).
		Render(sb.String())
}

// SlashCompleteMsg is sent when Tab completes a command.
type SlashCompleteMsg struct {
	Command string
}
