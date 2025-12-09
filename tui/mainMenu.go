package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/erikgeiser/promptkit/selection"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5AF"))
	cursorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5F5F"))
	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00D787")).
			Bold(true)
	itemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA"))
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5F5F87"))
)

type mainMenuModel struct {
	selectionPrompt *selection.Model[string]
	logger          *log.Logger
	choice          string
	isPrompt        bool
	err             error
}

func mainMenu() *mainMenuModel {
	items := []string{
		"Search Menu",
		"Calculate Factory",
	}

	m := selection.New("Select Task", items)
	m.FilterInputTextStyle = titleStyle
	mainMenuSelectionModel := selection.NewModel(m)

	return &mainMenuModel{
		selectionPrompt: mainMenuSelectionModel,
		logger:          LoggerFor("mainmenu"),
		isPrompt:        true,
	}
}

// Init functiom
func (m *mainMenuModel) Init() tea.Cmd {
	return m.selectionPrompt.Init()
}

func (m *mainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.isPrompt {
		promptModel, promptCmd := m.selectionPrompt.Update(msg)
		m.selectionPrompt = promptModel.(*selection.Model[string])

		if val, err := m.selectionPrompt.Value(); err == nil {
			m.choice = val
			m.isPrompt = false
		}
		return m, promptCmd
	}

	switch m.choice {
	case "Search Menu":
		searchMenuModel := searchMenu()
		return RootScreen().SwitchScreen(&searchMenuModel)
	}

	return m, nil
}

func (m mainMenuModel) View() string {
	return ""
}
