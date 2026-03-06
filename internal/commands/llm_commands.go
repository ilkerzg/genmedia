package commands

// LLMCommandPrompt holds the prompt template for an LLM-delegating command.
type LLMCommandPrompt struct {
	// WithArg is the prompt format when argument is provided (use %s for arg)
	WithArg string
	// WithoutArg is the prompt when no argument is given, or "" if arg is required
	WithoutArg string
}

// SearchCommand searches models.
type SearchCommand struct{}

func (c *SearchCommand) Name() string        { return "/search" }
func (c *SearchCommand) Aliases() []string   { return nil }
func (c *SearchCommand) ArgsHint() string    { return "[query]" }
func (c *SearchCommand) Description() string { return "Search models" }

// InfoCommand gets model details.
type InfoCommand struct{}

func (c *InfoCommand) Name() string        { return "/info" }
func (c *InfoCommand) Aliases() []string   { return nil }
func (c *InfoCommand) ArgsHint() string    { return "<endpoint>" }
func (c *InfoCommand) Description() string { return "Model details + params" }

// PriceCommand checks pricing.
type PriceCommand struct{}

func (c *PriceCommand) Name() string        { return "/price" }
func (c *PriceCommand) Aliases() []string   { return nil }
func (c *PriceCommand) ArgsHint() string    { return "<endpoint>" }
func (c *PriceCommand) Description() string { return "Check pricing" }

// EstimateCommand estimates cost.
type EstimateCommand struct{}

func (c *EstimateCommand) Name() string        { return "/estimate" }
func (c *EstimateCommand) Aliases() []string   { return nil }
func (c *EstimateCommand) ArgsHint() string    { return "<endpoint=qty>" }
func (c *EstimateCommand) Description() string { return "Cost estimate" }

// UsageCommand shows usage.
type UsageCommand struct{}

func (c *UsageCommand) Name() string        { return "/usage" }
func (c *UsageCommand) Aliases() []string   { return nil }
func (c *UsageCommand) ArgsHint() string    { return "[query]" }
func (c *UsageCommand) Description() string { return "Usage & cost info" }

// AnalyticsCommand shows analytics.
type AnalyticsCommand struct{}

func (c *AnalyticsCommand) Name() string        { return "/analytics" }
func (c *AnalyticsCommand) Aliases() []string   { return nil }
func (c *AnalyticsCommand) ArgsHint() string    { return "<endpoint>" }
func (c *AnalyticsCommand) Description() string { return "Performance analytics" }

// WorkflowsCommand lists workflows.
type WorkflowsCommand struct{}

func (c *WorkflowsCommand) Name() string        { return "/workflows" }
func (c *WorkflowsCommand) Aliases() []string   { return nil }
func (c *WorkflowsCommand) ArgsHint() string    { return "[query]" }
func (c *WorkflowsCommand) Description() string { return "List workflows" }

// FilesCommand lists files.
type FilesCommand struct{}

func (c *FilesCommand) Name() string        { return "/files" }
func (c *FilesCommand) Aliases() []string   { return nil }
func (c *FilesCommand) ArgsHint() string    { return "[path]" }
func (c *FilesCommand) Description() string { return "List files" }

// WhoamiCommand shows identity.
type WhoamiCommand struct{}

func (c *WhoamiCommand) Name() string        { return "/whoami" }
func (c *WhoamiCommand) Aliases() []string   { return nil }
func (c *WhoamiCommand) ArgsHint() string    { return "" }
func (c *WhoamiCommand) Description() string { return "Auth identity" }

// GetLLMPrompt returns the prompt for an LLM-delegating command, or "" if not an LLM command.
func GetLLMPrompt(cmdName, arg string) string {
	switch cmdName {
	case "/search":
		if arg != "" {
			return "Search for models matching \"" + arg + "\". Use the search_models tool."
		}
		return "Show me popular and recently updated models. Use the search_models tool."
	case "/info":
		if arg == "" {
			return ""
		}
		return "Get detailed info about model " + arg + ". Use the model_info tool. Show the key parameters, required inputs, and valid values."
	case "/price":
		if arg == "" {
			return ""
		}
		return "Check pricing for " + arg + ". Use the get_pricing tool."
	case "/estimate":
		if arg == "" {
			return ""
		}
		return "Estimate cost for " + arg + ". Use the estimate_cost tool."
	case "/usage":
		if arg != "" {
			return "Show me usage and cost information for " + arg + ". Use the check_usage and get_analytics tools."
		}
		return "Show me my recent usage and cost summary. Use the check_usage tool."
	case "/analytics":
		if arg == "" {
			return ""
		}
		return "Show performance analytics for " + arg + ". Use the get_analytics tool. Include latency, throughput, and error rates."
	case "/workflows":
		if arg != "" {
			return "List my workflows matching \"" + arg + "\". Use the list_workflows tool."
		}
		return "List all my saved workflows. Use the list_workflows tool."
	case "/files":
		if arg != "" {
			return "List files at path \"" + arg + "\". Use the list_files tool."
		}
		return "List files in my fal storage root. Use the list_files tool."
	case "/whoami":
		return "Show my current identity and key info. Use the whoami tool."
	}
	return ""
}
