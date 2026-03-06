package commands

// Command is the interface for slash commands.
type Command interface {
	Name() string
	Aliases() []string
	ArgsHint() string
	Description() string
}

// Registry holds registered commands.
type Registry struct {
	commands map[string]Command
	ordered  []Command // unique, in registration order
}

// NewRegistry creates a new command registry with all builtins registered.
func NewRegistry() *Registry {
	r := &Registry{commands: make(map[string]Command)}
	for _, cmd := range allCommands() {
		r.Register(cmd)
	}
	return r
}

// Register adds a command with its name and aliases.
func (r *Registry) Register(cmd Command) {
	r.commands[cmd.Name()] = cmd
	r.ordered = append(r.ordered, cmd)
	for _, alias := range cmd.Aliases() {
		r.commands[alias] = cmd
	}
}

// Get returns a command by name or alias.
func (r *Registry) Get(name string) Command {
	return r.commands[name]
}

// AllUnique returns all unique commands in registration order.
func (r *Registry) AllUnique() []Command {
	return r.ordered
}

// allCommands returns all built-in commands.
func allCommands() []Command {
	return []Command{
		&HelpCommand{},
		&ClearCommand{},
		&QuitCommand{},
		&ModelCommand{},
		&ThemeCommand{},
		&HistoryCommand{},
		&DefaultCommand{},
		&LoginCommand{},
		&LogoutCommand{},
		&SearchCommand{},
		&InfoCommand{},
		&PriceCommand{},
		&EstimateCommand{},
		&UsageCommand{},
		&AnalyticsCommand{},
		&WorkflowsCommand{},
		&FilesCommand{},
		&WhoamiCommand{},
	}
}
