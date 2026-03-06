package commands

// HelpCommand shows help.
type HelpCommand struct{}

func (c *HelpCommand) Name() string        { return "/help" }
func (c *HelpCommand) Aliases() []string   { return nil }
func (c *HelpCommand) ArgsHint() string    { return "" }
func (c *HelpCommand) Description() string { return "Show help" }

// ClearCommand clears conversation.
type ClearCommand struct{}

func (c *ClearCommand) Name() string        { return "/clear" }
func (c *ClearCommand) Aliases() []string   { return nil }
func (c *ClearCommand) ArgsHint() string    { return "" }
func (c *ClearCommand) Description() string { return "Clear conversation" }

// QuitCommand exits.
type QuitCommand struct{}

func (c *QuitCommand) Name() string        { return "/quit" }
func (c *QuitCommand) Aliases() []string   { return []string{"/exit", "/q"} }
func (c *QuitCommand) ArgsHint() string    { return "" }
func (c *QuitCommand) Description() string { return "Exit" }

// ModelCommand switches LLM model.
type ModelCommand struct{}

func (c *ModelCommand) Name() string        { return "/model" }
func (c *ModelCommand) Aliases() []string   { return nil }
func (c *ModelCommand) ArgsHint() string    { return "[name]" }
func (c *ModelCommand) Description() string { return "Switch LLM model" }

// ThemeCommand switches theme.
type ThemeCommand struct{}

func (c *ThemeCommand) Name() string        { return "/theme" }
func (c *ThemeCommand) Aliases() []string   { return nil }
func (c *ThemeCommand) ArgsHint() string    { return "[name]" }
func (c *ThemeCommand) Description() string { return "Switch theme" }

// HistoryCommand browses request history.
type HistoryCommand struct{}

func (c *HistoryCommand) Name() string        { return "/history" }
func (c *HistoryCommand) Aliases() []string   { return nil }
func (c *HistoryCommand) ArgsHint() string    { return "<endpoint>" }
func (c *HistoryCommand) Description() string { return "Browse request history" }

// DefaultCommand manages default models.
type DefaultCommand struct{}

func (c *DefaultCommand) Name() string        { return "/default" }
func (c *DefaultCommand) Aliases() []string   { return nil }
func (c *DefaultCommand) ArgsHint() string    { return "[category [endpoint]]" }
func (c *DefaultCommand) Description() string { return "Set default models" }

// LoginCommand authenticates with fal.ai.
type LoginCommand struct{}

func (c *LoginCommand) Name() string        { return "/login" }
func (c *LoginCommand) Aliases() []string   { return nil }
func (c *LoginCommand) ArgsHint() string    { return "" }
func (c *LoginCommand) Description() string { return "Log in to fal.ai" }

// LogoutCommand clears authentication.
type LogoutCommand struct{}

func (c *LogoutCommand) Name() string        { return "/logout" }
func (c *LogoutCommand) Aliases() []string   { return nil }
func (c *LogoutCommand) ArgsHint() string    { return "" }
func (c *LogoutCommand) Description() string { return "Log out of fal.ai" }
