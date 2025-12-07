package tui

import (
	"fmt"
	"math"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/sahilm/fuzzy"
	"github.com/shellserve/Satisfactory_Calc/internal/domain/file"
	"github.com/shellserve/Satisfactory_Calc/internal/domain/recipes"
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
	rec           recipes.Recipes
	paginator     paginator.Model
	showResults   bool
	logger        *log.Logger
	err           error
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
	contentByes, err := file.ReadFromFile("satisfactory_recipies.json")
	if err != nil {
		panic(err)
	}

	rec, err := recipes.LoadRecipeFromJSON(contentByes)
	if err != nil {
		panic(err)
	}

	return searchMenuModel{
		queryInput: ti,
		paginator:  p,
		logger:     LoggerFor("searchmenu"),
		rec:        rec,
	}
}

func (m searchMenuModel) Init() tea.Cmd {
	// Load recipe map here to reduce overhead during start of program
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
			return m, searchRecipesCmd(m.rec, m.queryString)

		case "left":
			if m.showResults {
				m.paginator.PrevPage()
			}
		case "right":
			if m.showResults {
				m.paginator.NextPage()
			}
		}

	case []recipes.RecipeEntry:
		// m.queryResults = msg
		m.queryRunning = false
		m.showResults = true
		m.logger.Info("Results", msg)
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
	m.logger.Info(
		fmt.Sprintf("Page Results: %s", pageResults),
	)

	s += "Results:\n"
	for _, r := range pageResults {
		if r != "" {
			s += fmt.Sprintf(":: %s\n", r)
		}
	}

	s += "\n" + m.paginator.View() + "\n"

	s += "\n(< > to change page, esc to quit)\n"

	return s
}

// Commands
func searchRecipesCmd(rec recipes.Recipes, query string) tea.Cmd {
	// return MSG
	// performs fuzz match
	// returns;[]Recipe Name, RecipeComponents{componentName: component Cost}, Building

	return func() tea.Msg {
		var allNames []string
		var allEntries []recipes.RecipeEntry

		for _, entries := range rec {
			for _, entry := range entries {
				allNames = append(allNames, entry.Name)
				allEntries = append(allEntries, entry)
			}
		}

		matches := fuzzy.Find(query, allNames)
		var results []recipes.RecipeEntry
		for _, match := range matches {
			results = append(results, allEntries[match.Index])
		}

		return results
	}
}
