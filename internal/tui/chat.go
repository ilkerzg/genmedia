package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

// ChatModel manages the chat viewport and rendered message blocks.
type ChatModel struct {
	viewport viewport.Model
	blocks   []string // rendered content blocks
	width    int
	height   int
	theme    *Theme

	// streaming state
	currentAssistantText string
	streaming            bool

	// tool block index tracking
	toolBlockIndex map[string]int // toolCallID → block index
}

// NewChatModel creates a new ChatModel.
func NewChatModel(width, height int, theme *Theme) ChatModel {
	vp := viewport.New(width, height)
	vp.SetContent("")
	return ChatModel{
		viewport: vp,
		blocks:   nil,
		width:    width,
		height:   height,
		theme:    theme,
	}
}

// SetSize resizes the chat viewport.
func (c *ChatModel) SetSize(w, h int) {
	c.width = w
	c.height = h
	c.viewport.Width = w
	c.viewport.Height = h
	c.refreshContent()
}

// AddUserMessage adds a user message with a secondary-colored left border.
func (c *ChatModel) AddUserMessage(content string) {
	rendered := c.renderWithBorder(content, c.theme.UserBorderStyle())
	c.blocks = append(c.blocks, rendered)
	c.refreshContent()
	c.viewport.GotoBottom()
}

// AddAssistantChunk appends content to the current streaming assistant block.
func (c *ChatModel) AddAssistantChunk(content string) {
	c.currentAssistantText += content
	// Render the streaming text as plain (no glamour yet)
	rendered := c.renderWithBorder(c.currentAssistantText, c.theme.AssistantBorderStyle())
	if c.streaming {
		// Replace the last block
		c.blocks[len(c.blocks)-1] = rendered
	} else {
		c.streaming = true
		c.blocks = append(c.blocks, rendered)
	}
	c.refreshContent()
	c.viewport.GotoBottom()
}

// FinalizeAssistant renders the complete assistant message with glamour markdown.
func (c *ChatModel) FinalizeAssistant() {
	if !c.streaming {
		return
	}
	contentWidth := c.contentWidth()
	md := c.renderMarkdown(c.currentAssistantText, contentWidth)
	rendered := c.renderWithBorder(md, c.theme.AssistantBorderStyle())
	c.blocks[len(c.blocks)-1] = rendered
	c.currentAssistantText = ""
	c.streaming = false
	c.refreshContent()
	c.viewport.GotoBottom()
}

// AddToolIndicator adds a tool call indicator with a muted border.
func (c *ChatModel) AddToolIndicator(toolCallID, name, argsPreview string) {
	icon := "○"
	line := fmt.Sprintf("%s %s %s", icon, name, lipgloss.NewStyle().Foreground(c.theme.TextMuted).Render(argsPreview))
	rendered := c.renderWithBorder(line, c.theme.ToolBorderStyle())
	c.blocks = append(c.blocks, rendered)
	if c.toolBlockIndex == nil {
		c.toolBlockIndex = make(map[string]int)
	}
	c.toolBlockIndex[toolCallID] = len(c.blocks) - 1
	c.refreshContent()
	c.viewport.GotoBottom()
}

// UpdateToolProgress updates the correct tool block with progress info.
func (c *ChatModel) UpdateToolProgress(toolCallID, line string) {
	idx, ok := c.toolBlockIndex[toolCallID]
	if !ok || idx >= len(c.blocks) {
		return
	}
	icon := "◎"
	content := fmt.Sprintf("%s %s", icon, line)
	rendered := c.renderWithBorder(content, c.theme.ToolBorderStyle())
	c.blocks[idx] = rendered
	c.refreshContent()
	c.viewport.GotoBottom()
}

// UpdateToolDone marks a tool as complete.
func (c *ChatModel) UpdateToolDone(toolCallID string, elapsed time.Duration, isError bool) {
	idx, ok := c.toolBlockIndex[toolCallID]
	if !ok || idx >= len(c.blocks) {
		return
	}
	icon := "●"
	style := c.theme.ToolBorderStyle()
	elapsedStr := fmt.Sprintf("%.1fs", elapsed.Seconds())
	if isError {
		icon = "✗"
		style = lipgloss.NewStyle().Foreground(c.theme.Error)
	}
	content := fmt.Sprintf("%s done %s", icon, lipgloss.NewStyle().Foreground(c.theme.TextMuted).Render(elapsedStr))
	rendered := c.renderWithBorder(content, style)
	c.blocks[idx] = rendered
	c.refreshContent()
	c.viewport.GotoBottom()
}

// AddSystemMessage adds a dim system message.
func (c *ChatModel) AddSystemMessage(content string) {
	rendered := c.renderWithBorder(
		lipgloss.NewStyle().Foreground(c.theme.TextMuted).Render(content),
		c.theme.SystemBorderStyle(),
	)
	c.blocks = append(c.blocks, rendered)
	c.refreshContent()
	c.viewport.GotoBottom()
}

// ClearBlocks removes all chat blocks.
func (c *ChatModel) ClearBlocks() {
	c.blocks = nil
	c.currentAssistantText = ""
	c.streaming = false
	c.refreshContent()
}

// Update handles viewport messages.
func (c ChatModel) Update(msg tea.Msg) (ChatModel, tea.Cmd) {
	var cmd tea.Cmd
	c.viewport, cmd = c.viewport.Update(msg)
	return c, cmd
}

// View returns the viewport view.
func (c ChatModel) View() string {
	return c.viewport.View()
}

// contentWidth returns the usable width for content (minus border + padding).
func (c *ChatModel) contentWidth() int {
	// " ┃ content" = 1 + 1 + 1 + content = 3 chars prefix, 1 char right pad
	w := c.width - 4
	if w < 20 {
		w = 20
	}
	return w
}

// renderWithBorder renders content with a colored left border bar.
func (c *ChatModel) renderWithBorder(content string, borderStyle lipgloss.Style) string {
	bar := borderStyle.Render(ThickBar)
	lines := strings.Split(content, "\n")
	var sb strings.Builder
	for i, line := range lines {
		if i > 0 {
			sb.WriteByte('\n')
		}
		sb.WriteString(" ")
		sb.WriteString(bar)
		sb.WriteString(" ")
		sb.WriteString(line)
	}
	return sb.String()
}

// renderMarkdown renders markdown text using glamour.
func (c *ChatModel) renderMarkdown(text string, width int) string {
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return text
	}
	out, err := r.Render(text)
	if err != nil {
		return text
	}
	return strings.TrimRight(out, "\n")
}

// refreshContent rebuilds the viewport content from blocks.
func (c *ChatModel) refreshContent() {
	content := strings.Join(c.blocks, "\n\n")
	c.viewport.SetContent(content)
}
