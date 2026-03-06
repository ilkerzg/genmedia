package tui

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fal-ai/genmedia/internal/auth"
	"github.com/fal-ai/genmedia/internal/commands"
	"github.com/fal-ai/genmedia/internal/config"
	"github.com/fal-ai/genmedia/internal/provider"
	"github.com/fal-ai/genmedia/internal/session"
	"github.com/fal-ai/genmedia/internal/tools"
)

// AppModel is the root Bubble Tea model.
type AppModel struct {
	chat      ChatModel
	input     InputModel
	status    StatusModel
	slashMenu *SlashMenuModel
	choice    *ChoiceModel

	// State
	model        string
	messages     []map[string]interface{}
	isGenerating bool
	themeKey     string
	theme        *Theme

	// Registries
	toolRegistry    *tools.Registry
	commandRegistry *commands.Registry

	// Provider
	provider provider.Provider

	// Session
	sessionID     string
	sessionStore  *session.SessionStore
	sessionTitled bool

	// Preferences
	preferences *config.Preferences

	// Auth
	authKey string

	// Program reference for goroutine sends.
	// Uses a holder so the pointer survives Bubble Tea's value copies.
	programRef *programRef

	// Window
	width  int
	height int

	// Flags
	showWelcome bool

	// Pending tool calls
	pendingTools     int
	toolResults      []map[string]interface{}
	currentToolCalls []ToolCall
}

// AppConfig holds initialization parameters.
type AppConfig struct {
	Model       string
	AuthKey     string
	ThemeKey    string
	SessionID   string
	Store       *session.SessionStore
	Preferences *config.Preferences
}

// NewAppModel creates the root model.
func NewAppModel(cfg AppConfig) AppModel {
	theme := GetTheme(cfg.ThemeKey)
	cmdRegistry := commands.NewRegistry()

	// Build system prompt
	systemPrompt := config.BuildSystemPrompt(cfg.Preferences)

	app := AppModel{
		model:           cfg.Model,
		authKey:         cfg.AuthKey,
		themeKey:        cfg.ThemeKey,
		theme:           theme,
		chat:            NewChatModel(80, 20, theme),
		input:           NewInputModel(theme),
		status:          NewStatusModel(cfg.Model, theme),
		slashMenu:       NewSlashMenuModel(cmdRegistry, theme),
		choice:          NewChoiceModel(theme),
		toolRegistry:    tools.NewRegistry(cfg.AuthKey),
		commandRegistry: cmdRegistry,
		provider:        &provider.OpenRouterProvider{AuthKey: cfg.AuthKey, BaseURL: config.OpenRouterBase},
		sessionStore:    cfg.Store,
		sessionID:       cfg.SessionID,
		preferences:     cfg.Preferences,
		programRef:      &programRef{},
		showWelcome:     true,
		messages: []map[string]interface{}{
			{"role": "system", "content": systemPrompt},
		},
	}

	// Load existing session if resuming
	if cfg.SessionID != "" && cfg.Store != nil {
		if msgs := cfg.Store.LoadMessages(cfg.SessionID); len(msgs) > 0 {
			app.messages = msgs
			app.showWelcome = false
			app.sessionTitled = true
			// Replay messages to chat view
			for _, msg := range msgs {
				role, _ := msg["role"].(string)
				content, _ := msg["content"].(string)
				switch role {
				case "user":
					app.chat.AddUserMessage(content)
				case "assistant":
					if content != "" {
						app.chat.AddAssistantChunk(content)
						app.chat.FinalizeAssistant()
					}
				}
			}
		}
		// Load session model
		if sess := cfg.Store.GetSession(cfg.SessionID); sess != nil {
			if m, ok := sess["model"].(string); ok && m != "" {
				app.model = m
				app.status.SetModel(m)
			}
		}
	}

	return app
}

// programRef holds a tea.Program pointer.
// Shared across all value copies of AppModel.
type programRef struct {
	p *tea.Program
}

// SetProgram sets the tea.Program reference for goroutine sends.
func (m *AppModel) SetProgram(p *tea.Program) {
	m.programRef.p = p
}

// Init returns the initial command.
func (m AppModel) Init() tea.Cmd {
	return m.input.textInput.Cursor.BlinkCmd()
}

// Update handles all messages.
func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.layout()
		return m, nil

	case tea.KeyMsg:
		// Priority: choice picker → slash menu → global keys → input
		if m.choice.IsVisible() {
			cmd := m.choice.Update(msg)
			return m, cmd
		}

		if m.slashMenu.IsVisible() {
			switch msg.Type {
			case tea.KeyUp, tea.KeyDown, tea.KeyEscape:
				cmd := m.slashMenu.Update(msg)
				return m, cmd
			case tea.KeyTab, tea.KeyEnter:
				// Both Tab and Enter autocomplete the selected command without submitting
				if selected := m.slashMenu.SelectedCommand(); selected != "" {
					m.input.textInput.SetValue(selected + " ")
					m.input.textInput.CursorEnd()
					m.slashMenu.Hide()
					return m, nil
				}
			}
		}

		switch {
		case msg.Type == tea.KeyCtrlC:
			return m, tea.Quit

		case msg.Type == tea.KeyCtrlL:
			m.clearChat()
			return m, nil

		case msg.Type == tea.KeyPgUp:
			m.chat.viewport.HalfViewUp()
			return m, nil

		case msg.Type == tea.KeyPgDown:
			m.chat.viewport.HalfViewDown()
			return m, nil

		case msg.Type == tea.KeyShiftUp:
			m.chat.viewport.LineUp(1)
			return m, nil

		case msg.Type == tea.KeyShiftDown:
			m.chat.viewport.LineDown(1)
			return m, nil

		case msg.Type == tea.KeyEscape:
			if m.slashMenu.IsVisible() {
				m.slashMenu.Hide()
			}
			cmd := m.input.Focus()
			return m, cmd

		default:
			var cmd tea.Cmd
			m.input, cmd = m.input.Update(msg)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}

			// Update slash menu based on input content
			val := m.input.textInput.Value()
			if strings.HasPrefix(val, "/") && !strings.Contains(val, " ") {
				m.slashMenu.Show(val)
			} else {
				m.slashMenu.Hide()
			}
		}

	case SlashCompleteMsg:
		m.input.textInput.SetValue(msg.Command + " ")
		m.input.textInput.CursorEnd()
		m.slashMenu.Hide()
		return m, nil

	case SubmitMsg:
		cmd := m.handleSubmit(msg.Content)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)

	case StreamChunkMsg:
		m.chat.AddAssistantChunk(msg.Content)
		return m, nil

	case StreamDoneMsg:
		m.chat.FinalizeAssistant()
		// Check for tool calls
		if tc, ok := msg.Message["tool_calls"]; ok && tc != nil {
			// Parse tool calls and execute
			return m, m.handleToolCalls(msg.Message)
		}
		// No tool calls — generation complete
		m.isGenerating = false
		m.status.SetGenerating(false, "")
		// Save assistant message
		m.messages = append(m.messages, msg.Message)
		m.saveMessage(msg.Message)
		return m, nil

	case StreamErrorMsg:
		m.chat.FinalizeAssistant()
		m.chat.AddSystemMessage("Error: " + msg.Err.Error())
		m.isGenerating = false
		m.status.SetGenerating(false, "")
		return m, nil

	case ToolCallsMsg:
		for _, tc := range msg.ToolCalls {
			preview := tc.Arguments
			if len(preview) > 60 {
				preview = preview[:60] + "..."
			}
			m.chat.AddToolIndicator(tc.ID, tc.Name, preview)
		}
		return m, nil

	case ToolProgressMsg:
		m.chat.UpdateToolProgress(msg.ToolCallID, msg.State)
		return m, nil

	case ToolDoneMsg:
		m.chat.UpdateToolDone(msg.ToolCallID, msg.Elapsed, msg.IsError)
		m.pendingTools--
		// Store tool result
		toolResult := map[string]interface{}{
			"role":         "tool",
			"content":      msg.Result,
			"tool_call_id": msg.ToolCallID,
		}
		m.toolResults = append(m.toolResults, toolResult)
		m.messages = append(m.messages, toolResult)
		m.saveMessage(toolResult)

		// Handle media URLs from generate results
		if msg.Name == "generate" && !msg.IsError {
			var resultData map[string]interface{}
			if err := json.Unmarshal([]byte(msg.Result), &resultData); err == nil {
				if result, ok := resultData["result"]; ok {
					mediaURLs := ExtractMediaURLs(result)
					for _, mu := range mediaURLs {
						media := NewMediaModel(mu.URL, mu.MediaType, m.theme, m.width)
						rendered := media.View()
						m.chat.AddSystemMessage(rendered)
						// Launch async load
						cmds = append(cmds, media.LoadCmd())
					}
				}
			}
		}

		// All tools done → continue generation loop
		if m.pendingTools <= 0 {
			return m, m.startGeneration()
		}
		return m, tea.Batch(cmds...)

	case AskUserRequestMsg:
		m.choice.Show(msg.Question, msg.Options, msg.ResponseCh)
		m.input.Blur()
		return m, nil

	case AskUserResponseMsg:
		m.input.Focus()
		if msg.Answer != "" {
			m.chat.AddSystemMessage("→ " + msg.Answer)
		}
		return m, nil

	case GenerationLoopMsg:
		return m, m.startGeneration()

	case MediaLoadedMsg:
		if msg.Err != nil {
			m.chat.AddSystemMessage("Media load error: " + msg.Err.Error())
		}
		return m, nil

	case MediaTickMsg:
		// Animation tick — no-op for now
		return m, nil

	case modelSelectedMsg:
		m.model = msg.model
		m.status.SetModel(msg.model)
		if m.sessionStore != nil && m.sessionID != "" {
			m.sessionStore.UpdateModel(m.sessionID, msg.model)
		}
		m.chat.AddSystemMessage("Model changed to: " + msg.model)
		cmd := m.input.Focus()
		return m, cmd

	case LoginResultMsg:
		cmd := m.input.Focus()
		if msg.Err != nil {
			if msg.Err.Error() != "cancelled" {
				m.chat.AddSystemMessage("Login failed: " + msg.Err.Error())
			}
			return m, cmd
		}
		if msg.Key != "" {
			m.authKey = msg.Key
			m.provider = &provider.OpenRouterProvider{AuthKey: msg.Key, BaseURL: config.OpenRouterBase}
			m.toolRegistry = tools.NewRegistry(msg.Key)
			m.chat.AddSystemMessage("Logged in successfully.")
		}
		return m, cmd

	case loginBrowserMsg:
		m.chat.AddSystemMessage("Opening browser for fal.ai login...")
		prog := m.programRef.p
		return m, func() tea.Msg {
			token, err := auth.DeviceCodeLogin(func(status string) {
				prog.Send(loginStatusMsg{status: status})
			})
			if err != nil {
				return LoginResultMsg{Err: err}
			}
			// Save tokens
			if err := auth.SaveTokens(token.RefreshToken, token.AccessToken); err != nil {
				return LoginResultMsg{Err: fmt.Errorf("failed to save tokens: %w", err)}
			}
			return LoginResultMsg{Key: "Bearer " + token.AccessToken}
		}

	case loginStatusMsg:
		m.chat.AddSystemMessage(msg.status)
		return m, nil

	case loginManualKeyMsg:
		// Show choice picker with "Other..." to get the key
		ch := make(chan string, 1)
		m.choice.Show("Paste your fal.ai API key (from fal.ai/dashboard/keys):", []string{}, ch)
		m.input.Blur()
		return m, func() tea.Msg {
			answer := <-ch
			if answer == "" {
				return LoginResultMsg{Err: fmt.Errorf("cancelled")}
			}
			key := strings.TrimSpace(answer)
			if err := auth.SaveAuthKey(key); err != nil {
				return LoginResultMsg{Err: err}
			}
			return LoginResultMsg{Key: auth.GetCredentials()}
		}

	default:
		// Pass through to sub-models
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		m.chat, cmd = m.chat.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

// View renders the full UI.
func (m AppModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	chatH := m.chatHeight()
	statusView := m.status.View()
	inputView := m.input.View()

	// Overlays reduce available chat height
	overlay := ""
	overlayHeight := 0
	if m.choice.IsVisible() {
		overlay = m.choice.View()
		overlayHeight = lipgloss.Height(overlay)
	} else if m.slashMenu.IsVisible() {
		overlay = m.slashMenu.View()
		overlayHeight = lipgloss.Height(overlay)
	}

	effectiveChatH := chatH - overlayHeight
	if effectiveChatH < 1 {
		effectiveChatH = 1
	}

	var chatView string
	if m.showWelcome && len(m.chat.blocks) == 0 {
		welcome := RenderWelcome(m.width, m.theme, m.model)
		chatView = padToHeight(welcome, effectiveChatH, m.width)
	} else {
		// Temporarily resize viewport for overlay
		if overlayHeight > 0 {
			m.chat.viewport.Height = effectiveChatH
		}
		chatView = m.chat.View()
		if overlayHeight > 0 {
			m.chat.viewport.Height = chatH
		}
	}

	if overlay != "" {
		return lipgloss.JoinVertical(lipgloss.Left,
			chatView,
			overlay,
			statusView,
			inputView,
		)
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		chatView,
		statusView,
		inputView,
	)
}

// padToHeight pads content to exactly h lines, vertically centered (upper-third bias).
func padToHeight(content string, h, _ int) string {
	lines := strings.Split(content, "\n")

	// Trim trailing empty lines from content
	for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}

	if len(lines) >= h {
		return strings.Join(lines[:h], "\n")
	}

	// Position content in the upper third for a balanced feel
	topPad := (h - len(lines)) / 4
	if topPad < 2 {
		topPad = 2
	}

	result := make([]string, 0, h)
	for i := 0; i < topPad; i++ {
		result = append(result, "")
	}
	result = append(result, lines...)
	for len(result) < h {
		result = append(result, "")
	}
	return strings.Join(result[:h], "\n")
}

// handleSubmit processes user input.
func (m *AppModel) handleSubmit(content string) tea.Cmd {
	m.showWelcome = false
	m.slashMenu.Hide()

	// Check for slash commands
	if strings.HasPrefix(content, "/") {
		return m.handleSlashCommand(content)
	}

	if m.isGenerating {
		return nil
	}

	m.chat.AddUserMessage(content)

	// Create session if needed
	if m.sessionID == "" && m.sessionStore != nil {
		m.sessionID = m.sessionStore.CreateSession(m.model, "")
	}

	// Store message
	msg := map[string]interface{}{
		"role":    "user",
		"content": content,
	}
	m.messages = append(m.messages, msg)
	m.saveMessage(msg)

	// Auto-title from first user message
	if !m.sessionTitled && m.sessionStore != nil && m.sessionID != "" {
		title := content
		if len(title) > 60 {
			title = title[:60]
		}
		m.sessionStore.UpdateTitle(m.sessionID, title)
		m.sessionTitled = true
	}

	// Start generation
	return m.startGeneration()
}

// handleSlashCommand processes a slash command.
func (m *AppModel) handleSlashCommand(input string) tea.Cmd {
	parts := strings.SplitN(input, " ", 2)
	cmdName := strings.ToLower(parts[0])
	arg := ""
	if len(parts) > 1 {
		arg = strings.TrimSpace(parts[1])
	}

	cmd := m.commandRegistry.Get(cmdName)
	if cmd == nil {
		m.chat.AddSystemMessage(fmt.Sprintf("Unknown command: %s. Type /help for available commands.", cmdName))
		return nil
	}

	// Handle built-in commands directly
	switch cmdName {
	case "/help":
		m.showHelp()
	case "/clear":
		m.clearChat()
	case "/quit", "/exit", "/q":
		return tea.Quit
	case "/model":
		return m.handleModelCommand(arg)
	case "/theme":
		m.handleThemeCommand(arg)
	case "/login":
		return m.handleLoginCommand()
	case "/logout":
		m.handleLogoutCommand()
		return nil
	default:
		// LLM-delegating commands
		prompt := commands.GetLLMPrompt(cmdName, arg)
		if prompt == "" {
			m.chat.AddSystemMessage(fmt.Sprintf("Usage: %s %s", cmd.Name(), cmd.ArgsHint()))
			return nil
		}
		m.chat.AddUserMessage(input)
		msg := map[string]interface{}{"role": "user", "content": prompt}
		m.messages = append(m.messages, msg)
		m.saveMessage(msg)
		return m.startGeneration()
	}
	return nil
}

func (m *AppModel) showHelp() {
	var sb strings.Builder
	sb.WriteString("Commands:\n\n")
	for _, cmd := range m.commandRegistry.AllUnique() {
		name := cmd.Name()
		if cmd.ArgsHint() != "" {
			name += " " + cmd.ArgsHint()
		}
		sb.WriteString(fmt.Sprintf("  %-25s %s\n", name, cmd.Description()))
	}
	sb.WriteString(fmt.Sprintf("\nThemes: %s\n", strings.Join(themeNames(), ", ")))
	m.chat.AddSystemMessage(sb.String())
}

func (m *AppModel) clearChat() {
	m.chat.ClearBlocks()
	m.showWelcome = true
	// Keep system message, clear the rest
	if len(m.messages) > 0 {
		m.messages = m.messages[:1]
	}
}

func (m *AppModel) handleModelCommand(arg string) tea.Cmd {
	if arg != "" {
		m.model = arg
		m.status.SetModel(arg)
		if m.sessionStore != nil && m.sessionID != "" {
			m.sessionStore.UpdateModel(m.sessionID, arg)
		}
		m.chat.AddSystemMessage("Model changed to: " + arg)
		return nil
	}

	// Show choice picker
	ch := make(chan string, 1)
	m.choice.Show("Select LLM model:", config.LLMModels, ch)
	m.input.Blur()

	// The response will come via AskUserResponseMsg
	// We need a goroutine to wait for the channel
	return func() tea.Msg {
		answer := <-ch
		if answer != "" {
			return modelSelectedMsg{model: answer}
		}
		return AskUserResponseMsg{Answer: ""}
	}
}

type modelSelectedMsg struct {
	model string
}

func (m *AppModel) handleThemeCommand(arg string) {
	if arg != "" {
		if _, ok := config.Themes[arg]; ok {
			m.themeKey = arg
			m.theme = GetTheme(arg)
			m.chat.theme = m.theme
			m.input.theme = m.theme
			m.status.theme = m.theme
			m.chat.AddSystemMessage("Theme changed to: " + config.Themes[arg].Name)
		} else {
			m.chat.AddSystemMessage("Unknown theme: " + arg + ". Available: " + strings.Join(themeNames(), ", "))
		}
		return
	}
	m.chat.AddSystemMessage("Available themes: " + strings.Join(themeNames(), ", ") + " (current: " + m.themeKey + ")")
}

// handleLoginCommand handles the /login slash command.
func (m *AppModel) handleLoginCommand() tea.Cmd {
	// Check if already authenticated
	if m.authKey != "" {
		m.chat.AddSystemMessage("Already logged in. Key is active.")
		return nil
	}

	// Show login method options
	options := []string{
		"Log in with browser",
		"Enter API key manually",
	}

	ch := make(chan string, 1)
	m.choice.Show("How would you like to log in?", options, ch)
	m.input.Blur()

	return func() tea.Msg {
		answer := <-ch
		if answer == "" {
			return LoginResultMsg{Err: fmt.Errorf("cancelled")}
		}
		switch {
		case strings.HasPrefix(answer, "Log in with browser"):
			return loginBrowserMsg{}
		case answer == "Enter API key manually":
			return loginManualKeyMsg{}
		default:
			// "Other..." — treat as a pasted API key
			key := strings.TrimSpace(answer)
			if err := auth.SaveAuthKey(key); err != nil {
				return LoginResultMsg{Err: err}
			}
			return LoginResultMsg{Key: auth.GetCredentials()}
		}
	}
}

type loginBrowserMsg struct{}
type loginManualKeyMsg struct{}
type loginStatusMsg struct{ status string }

// startGeneration launches the SSE stream in a goroutine.
func (m *AppModel) handleLogoutCommand() {
	_ = auth.ClearTokens()
	auth.ClearCachedKey()
	m.authKey = ""
	m.chat.AddSystemMessage("Logged out. Use /login to authenticate again.")
}

func (m *AppModel) startGeneration() tea.Cmd {
	m.isGenerating = true
	m.status.SetGenerating(true, config.RandomTagline())
	m.toolResults = nil

	messages := make([]map[string]interface{}, len(m.messages))
	copy(messages, m.messages)

	toolSchemas := m.toolRegistry.OpenAISchemas()
	model := m.model
	p := m.provider
	prog := m.programRef.p

	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		events := p.StreamChat(ctx, model, messages, toolSchemas)

		for event := range events {
			switch event.Type {
			case "content":
				if content, ok := event.Data.(string); ok {
					prog.Send(StreamChunkMsg{Content: content})
				}
			case "tool_calls":
				if tcList, ok := event.Data.([]provider.ToolCall); ok {
					tuiCalls := make([]ToolCall, len(tcList))
					for i, tc := range tcList {
						tuiCalls[i] = ToolCall{ID: tc.ID, Name: tc.Function.Name, Arguments: tc.Function.Arguments}
					}
					prog.Send(ToolCallsMsg{ToolCalls: tuiCalls})
				}
			case "done":
				if msg, ok := event.Data.(map[string]interface{}); ok {
					return StreamDoneMsg{Message: msg}
				}
			case "error":
				if errMsg, ok := event.Data.(string); ok {
					return StreamErrorMsg{Err: fmt.Errorf("%s", errMsg)}
				}
			}
		}

		return StreamErrorMsg{Err: fmt.Errorf("stream ended unexpectedly")}
	}
}

// handleToolCalls processes tool calls from the assistant message.
func (m *AppModel) handleToolCalls(assistantMsg map[string]interface{}) tea.Cmd {
	// Save assistant message
	m.messages = append(m.messages, assistantMsg)
	m.saveMessage(assistantMsg)

	// Parse tool calls
	tcRaw, _ := assistantMsg["tool_calls"].([]interface{})
	if len(tcRaw) == 0 {
		return nil
	}

	var toolCalls []ToolCall
	for _, raw := range tcRaw {
		tc, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		id, _ := tc["id"].(string)
		fn, _ := tc["function"].(map[string]interface{})
		name, _ := fn["name"].(string)
		args, _ := fn["arguments"].(string)
		toolCalls = append(toolCalls, ToolCall{ID: id, Name: name, Arguments: args})
	}

	m.currentToolCalls = toolCalls
	m.pendingTools = len(toolCalls)
	m.toolResults = nil

	// Add indicators only for tool calls that don't already have a block
	for _, tc := range toolCalls {
		if _, exists := m.chat.toolBlockIndex[tc.ID]; !exists {
			preview := tc.Arguments
			if len(preview) > 60 {
				preview = preview[:60] + "..."
			}
			m.chat.AddToolIndicator(tc.ID, tc.Name, preview)
		}
	}

	// Launch tool executions
	var cmds []tea.Cmd
	for _, tc := range toolCalls {
		tc := tc // capture
		prog := m.programRef.p
		registry := m.toolRegistry

		cmds = append(cmds, func() tea.Msg {
			start := time.Now()

			onProgress := func(info map[string]interface{}) {
				// Handle ask_user bridge
				if t, ok := info["type"].(string); ok && t == "ask_user" {
					q, _ := info["question"].(string)
					opts, _ := info["options"].([]string)
					ch, _ := info["response_ch"].(chan string)
					prog.Send(AskUserRequestMsg{
						Question:   q,
						Options:    opts,
						ResponseCh: ch,
					})
					return
				}

				state, _ := info["state"].(string)
				elapsed, _ := info["elapsed"].(float64)
				line := state
				if elapsed > 0 {
					line += fmt.Sprintf(" %.1fs", elapsed)
				}
				if pos, ok := info["position"]; ok {
					line += fmt.Sprintf(" (queue: %v)", pos)
				}
				prog.Send(ToolProgressMsg{
					ToolCallID: tc.ID,
					State:      line,
				})
			}

			result := registry.Execute(tc.Name, tc.Arguments, onProgress)
			elapsed := time.Since(start)

			isError := false
			var resultData map[string]interface{}
			if err := json.Unmarshal([]byte(result), &resultData); err == nil {
				if ok, exists := resultData["ok"]; exists {
					if b, isBool := ok.(bool); isBool && !b {
						isError = true
					}
				}
			}

			return ToolDoneMsg{
				ToolCallID: tc.ID,
				Name:       tc.Name,
				Result:     result,
				Elapsed:    elapsed,
				IsError:    isError,
			}
		})
	}

	return tea.Batch(cmds...)
}

// saveMessage persists a message to the session store.
func (m *AppModel) saveMessage(msg map[string]interface{}) {
	if m.sessionStore != nil && m.sessionID != "" {
		m.sessionStore.SaveMessage(m.sessionID, msg)
	}
}

// layout recalculates sizes for all sub-models.
func (m *AppModel) layout() {
	statusHeight := 1
	inputHeight := 3 // 1 content + 2 border lines (rounded border)

	chatH := m.height - statusHeight - inputHeight
	if chatH < 1 {
		chatH = 1
	}

	m.chat.SetSize(m.width, chatH)
	m.status.SetWidth(m.width)
	m.input.SetWidth(m.width)
	m.slashMenu.SetWidth(m.width)
	m.choice.SetWidth(m.width)
}

// chatHeight returns the chat viewport height.
func (m *AppModel) chatHeight() int {
	return m.height - 4 // 1 status + 3 input (1 content + 2 border lines)
}

func themeNames() []string {
	names := make([]string, 0, len(config.Themes))
	for k := range config.Themes {
		names = append(names, k)
	}
	return names
}
