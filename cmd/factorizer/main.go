package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"Satisfactory_Calc/internal/domain/io"
	"Satisfactory_Calc/internal/domain/recipes"
	"Satisfactory_Calc/internal/domain/scraper"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fogleman/ease"
)

const (
	dotChar = " • "
)

type (
	tickMsg  struct{}
	frameMsg struct{}
)

// Styles
var (
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
	ticksStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
)

type mainModel struct {
	Choice   int
	Chosen   bool
	Ticks    int
	Loaded   bool
	Quitting bool
	Frames   int
	Progress float64
	Recipes  recipes.Recipes
}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[] %s", label)
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return frameMsg{}
	})
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func UpdateChosen(msg tea.Msg, m mainModel) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case frameMsg:
		if !m.Loaded {
			m.Frames++
			m.Progress = ease.OutBounce(float64(m.Frames) / float64(100))
			if m.Progress >= 1 {
				m.Progress = 1
				m.Loaded = true
				m.Ticks = 3
				return m, tick()
			}
			return m, frame()
		}
	}
	return m, nil
}

func UpdateChoices(msg tea.Msg, m mainModel) (tea.Model, tea.Cmd) {
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
			m.Chosen = true
		}
	}
	return m, nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl-c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func choicesView(m mainModel) string {
	c := m.Choice

	tpl := "Satisfactory Calculator Tool"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("q, esc: quit")

	choices := fmt.Sprintf("%s\n", checkbox("Search Recipe: ", c == 0))

	return fmt.Sprintf(tpl, choices, ticksStyle.Render(strconv.Itoa(m.Ticks)))
}

func (m mainModel) View() string {
	var s string
	if m.Quitting {
		return "\n Quitting Time!\n\n"
	}
	if !m.Chosen {
		s = choicesView(m)
	}
	return mainStyle.Render("\n" + s + "\n\n")
}

func main() {
	fmt.Printf("Satisfactory Factorizor\n")

	// itemRate := "Items Per Minute"
	// liquidRate := "Meters Cubed Per Minute"

	// Data collection phase
	outputFilename := "satisfactory_recipies.json"

	if !io.FileExists(outputFilename) {
		if data, err := scraper.FetchRecipes(); err != nil {
			log.Fatalf("Error during scraping process: %s", err)
		} else {
			_, err := io.WriteStringToFile(outputFilename, data)
			if err != nil {
				log.Fatalf("Error WriteStringToFile: %v", err)
			}
		}
	}

	contentBytes, err := io.ReadFromFile(outputFilename)
	if err != nil {
		fmt.Printf("Error in opening file: %v", err)
		return
	}

	// Main Logic
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter desired factory output: ")

	outputItem, err := reader.ReadString('\n')
	log.Println(outputItem)

	if err != nil {
		log.Fatalf("Error during input: %v", err)
	}

	r, err := recipes.LoadRecipeFromJSON(contentBytes)
	if err != nil {
		log.Fatalf("Error while fetching recipe: %v", err)
	}

	// print out of recipe names for debugging
	//for key := range r {
	//	log.Println(key)
	//}

	initalModel := mainModel{0, false, 10, false, false, 0, 0, r}
	p := tea.NewProgram(initalModel)

	if _, err := p.Run(); err != nil {
		log.Fatalf("Could not start program: %v", err)
	}
}
