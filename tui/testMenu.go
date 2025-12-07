package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type testMenuModel struct {
	spinner spinner.Model
	err     error
}

func testMenu() testMenuModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return testMenuModel{
		spinner: s,
	}
}

func (m testMenuModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m testMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		default:
			mainMenu := mainMenu()
			return RootScreen().SwitchScreen(&mainMenu)
			// return mainMenu, mainMenu.Init()
		}

	case error:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	}
	// return m, nil
}

func (m testMenuModel) View() string {
	str := fmt.Sprintf("\n %s This is the test screen.\n\n\n", m.spinner.View())
	return str
}
