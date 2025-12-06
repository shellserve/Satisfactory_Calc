package main

import (
	"Satisfactory_Calc/internal/domain/recipes"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Recipes     recipes.Recipes
	SearchQuery textinput.Model
	Quitting    bool
}

func (m model) Init() tea.Cmd {
	return nil /// may load db here not fully sure yet
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl-c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}
