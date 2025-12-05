package application

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	input     string
	activeCmd tea.Model
	registry  *CommandRegistry
}

func New(registry *CommandRegistry) Model {
	return Model{registry: registry}
}
