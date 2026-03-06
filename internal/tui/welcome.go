package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/fal-ai/genmedia/internal/config"
)

// RenderWelcome renders the welcome screen with ASCII logo and help commands.
func RenderWelcome(width int, theme *Theme, model string) string {
	var sb strings.Builder

	// Render two-tone ASCII logo
	logoGen := config.LogoGen
	logoMedia := config.LogoMedia

	genStyle := lipgloss.NewStyle().Foreground(theme.Primary).Bold(true)
	mediaStyle := lipgloss.NewStyle().Foreground(theme.TextMuted)

	// Pad all logo lines to uniform width so centering doesn't shift
	maxGenLen := 0
	for _, s := range logoGen {
		if n := len([]rune(s)); n > maxGenLen {
			maxGenLen = n
		}
	}
	maxMediaLen := 0
	for _, s := range logoMedia {
		if n := len([]rune(s)); n > maxMediaLen {
			maxMediaLen = n
		}
	}

	maxLines := len(logoGen)
	if len(logoMedia) > maxLines {
		maxLines = len(logoMedia)
	}

	var logoLines []string
	for i := 0; i < maxLines; i++ {
		gen := ""
		media := ""
		if i < len(logoGen) {
			gen = logoGen[i]
		}
		if i < len(logoMedia) {
			media = logoMedia[i]
		}
		// Pad to uniform width
		gen = gen + strings.Repeat(" ", maxGenLen-len([]rune(gen)))
		media = media + strings.Repeat(" ", maxMediaLen-len([]rune(media)))
		line := genStyle.Render(gen) + "  " + mediaStyle.Render(media)
		logoLines = append(logoLines, line)
	}
	logoBlock := strings.Join(logoLines, "\n")

	// Center the logo
	sb.WriteString(lipgloss.PlaceHorizontal(width, lipgloss.Center, logoBlock))
	sb.WriteString("\n\n")

	// Command table — build as a single block, then center once
	cmdStyle := lipgloss.NewStyle().Foreground(theme.Primary).Bold(true)
	descStyle := lipgloss.NewStyle().Foreground(theme.TextMuted)

	// Find max description length to build fixed-width lines
	maxDescLen := 0
	for _, cmd := range config.HelpCommands {
		if len(cmd.Description) > maxDescLen {
			maxDescLen = len(cmd.Description)
		}
	}

	var tableLines []string
	for _, cmd := range config.HelpCommands {
		cmdText := fmt.Sprintf("%14s", cmd.Command)
		descText := fmt.Sprintf("%-*s", maxDescLen, cmd.Description)
		line := cmdStyle.Render(cmdText) + "    " + descStyle.Render(descText)
		tableLines = append(tableLines, line)
	}

	// Add model line to the same block
	tableLines = append(tableLines, "")
	modelLabel := fmt.Sprintf("%14s", "model")
	modelLine := lipgloss.NewStyle().Foreground(theme.TextMuted).Render(modelLabel) +
		"    " +
		lipgloss.NewStyle().Foreground(theme.Accent).Bold(true).Render(fmt.Sprintf("%-*s", maxDescLen, model))
	tableLines = append(tableLines, modelLine)

	tableBlock := strings.Join(tableLines, "\n")
	sb.WriteString(lipgloss.PlaceHorizontal(width, lipgloss.Center, tableBlock))
	sb.WriteString("\n")

	return sb.String()
}
