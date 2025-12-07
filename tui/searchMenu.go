package tui

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

type (
	errMsg error
)

type searchMenuModel struct {
	queryInput    textinput.Model
	queryString   string
	querySelected bool
	queryRunning  bool
	queryResults  []string

	paginator   paginator.Model
	showResults bool
	logger      *log.Logger
	err         error
}

func searchMenu() searchMenuModel {
	ti := textinput.New()
	ti.Placeholder = "nuclear speghetti"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 30

	p := paginator.New()
	p.PerPage = 5
	p.Type = paginator.Dots

	return searchMenuModel{
		queryInput: ti,
		paginator:  p,
		logger:     LoggerFor("searchmenu"),
	}
}

func (m searchMenuModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m searchMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl-c", "esc", "q":
			m.logger.Info("Quitting")
			return m, tea.Quit

		case "enter":
			m.queryString = m.queryInput.Value()
			m.showResults = false
			m.queryRunning = true

			m.logger.Info("Running search", "query", m.queryString)
			return m, searchRecipesCmd(m.queryString)

		case "left":
			if m.showResults {
				m.paginator.PrevPage()
			}
		case "right":
			if m.showResults {
				m.paginator.NextPage()
			}
		}

	case []string:
		m.queryResults = msg
		m.queryRunning = false
		m.showResults = true

		m.paginator.SetTotalPages(
			int(math.Ceil(float64(len(m.queryResults)) / float64(m.paginator.PerPage))),
		)

		return m, nil

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.queryInput, cmd = m.queryInput.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m searchMenuModel) View() string {
	s := fmt.Sprintf("Search for Recipe: %s\n\n\n", m.queryInput.View())

	if m.queryRunning {
		s += "Searching...\n"
		return s
	}

	if !m.showResults {
		return s + "(enter to search, esc to quit)\n"
	}

	start, end := m.paginator.GetSliceBounds(len(m.queryResults))
	pageResults := m.queryResults[start:end]

	s += "Results:\n"
	for _, r := range pageResults {
		s += fmt.Sprintf("• %s\n", r)
	}

	s += "\n" + m.paginator.View() + "\n"

	s += "\n(< > to change page, esc to quit)\n"

	return s
}

// Commands
func searchRecipesCmd(query string) tea.Cmd {
	return func() tea.Msg {
		results := []string{
			"Pasta Carbonara",
			"Nuclear Spaghetti",
			"Atomic Alfredo",
			"Radioactive Ragu",
			"Quantum Lasagna",
			"Particle Penne",
		}

		filtered := []string{}
		for _, r := range results {
			if strings.Contains(strings.ToLower(r), strings.ToLower(query)) {
				filtered = append(filtered, r)
			}
		}
		return filtered
	}
}
