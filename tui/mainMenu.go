package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"

	"github.com/shellserve/Satisfactory_Calc/internal/domain/file"
	"github.com/shellserve/Satisfactory_Calc/internal/domain/scraper"
)

type mainMenuModel struct {
	Choice int
	Chosen bool
	logger *log.Logger
}

func mainMenu() mainMenuModel {
	return mainMenuModel{
		logger: LoggerFor("mainmenu"),
	}
}

// Init functiom
func (m mainMenuModel) Init() tea.Cmd {
	if !file.FileExists("satisfactory_recipies.json") {
		m.logger.Info("Fetching recipe JSON")
		data, err := scraper.FetchRecipes()
		if err != nil {
			m.logger.Error("Error fetching recipe JSON", "error", err)
			panic(
				fmt.Sprintf("Failed to FetchRecipes: %s", err))
		}
		m.logger.Info("Fetching completed - writting to file")
		_, err = file.WriteStringToFile("satisfactory_recipies.json", data)
		if err != nil {
			m.logger.Error("Error writing to file", "error", err)
			panic(
				fmt.Sprintf("Failed to WriteStringToFile: %s", err))
		}
		m.logger.Info("Recipe succesfully fetched and written to file!")
	}
	return nil
}

// Update functions
func (m mainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl-c", "esc", "q":
			m.logger.Info("Exiting")
			return m, tea.Quit
		default:
			if !m.Chosen {
				return updateChoices(msg, m)
			}
			return updateChosen(msg, m)
		}
	}
	return m, nil
}

// Sub Update functions
func updateChoices(msg tea.Msg, m mainMenuModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice++
			if m.Choice > 3 {
				m.Choice = 3
			}
		case "k", "up":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = 0
			}
		case "enter":
			m.logger.Info(
				fmt.Sprintf("User selected: %d", m.Choice))
			m.Chosen = true
			return m, nil
		}
	}
	return m, nil
}

func updateChosen(msg tea.Msg, m mainMenuModel) (tea.Model, tea.Cmd) {
	// to be implemented
	switch m.Choice {
	case 0:
		searchMenu := searchMenu()
		return RootScreen().SwitchScreen(&searchMenu)

	default:
		return m, nil
	}
}

// View Function
func (m mainMenuModel) View() string {
	var s string
	if !m.Chosen {
		s = choicesView(m)
	} else {
		s = choicesView(m)
	}
	return MainStyle.Render("\n" + s + "\n\n")
}

// Sub View functions
func choicesView(m mainMenuModel) string {
	c := m.Choice
	tpl := "Welcome to the Satisfactory Calculation Suite\n\n"
	tpl += "%s\n\n"
	tpl += SubTitleStyle.Render("j/k, up/down: select") + DotStyle +
		SubTitleStyle.Render("enter: choose") + DotStyle +
		SubTitleStyle.Render("ctrl-C, esc: quit")

	choices := fmt.Sprintf("%s\n%s\n",
		checkbox("Search Recipe", c == 0),
		checkbox("To be implemented", c == 1),
	)

	return fmt.Sprintf(tpl, choices)
}

// helper functions
func checkbox(label string, checked bool) string {
	if checked {
		return CheckboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}
