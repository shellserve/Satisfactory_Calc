package application

import (
	tea "github.com/charmbracelet/bubbletea"
)

type CommandFactory func(args []string) tea.Model

type CommandRegistry struct {
	commands map[string]CommandFactory
}

func NewRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: make(map[string]CommandFactory),
	}
}

func (r *CommandRegistry) Register(name string, create CommandFactory) {
	r.commands[name] = create
}

func (r *CommandRegistry) Get(name string) (CommandFactory, bool) {
	f, ok := r.commands[name]
	return f, ok
}
