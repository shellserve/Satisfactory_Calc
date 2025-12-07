// Package tui holds the different bubbletea tui menus and menu logic for the project
package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

type rootScreenModel struct {
	logger *log.Logger
	model  tea.Model
}

func RootScreen() rootScreenModel {
	var rootModel tea.Model

	mainMenuScreen := mainMenu()
	rootModel = &mainMenuScreen

	return rootScreenModel{
		logger: LoggerFor("root"),
		model:  rootModel,
	}
}

func (m rootScreenModel) Init() tea.Cmd {
	return m.model.Init()
}

func (m rootScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.model.Update(msg)
}

func (m rootScreenModel) View() string {
	return m.model.View()
}

func (m rootScreenModel) SwitchScreen(model tea.Model) (tea.Model, tea.Cmd) {
	m.model = model
	return m.model, m.model.Init()
}
