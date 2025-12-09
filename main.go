// Package satisfactorycalc is for caling
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/shellserve/Satisfactory_Calc/internal/domain/file"
	"github.com/shellserve/Satisfactory_Calc/internal/domain/scraper"
	"github.com/shellserve/Satisfactory_Calc/tui"
)

const (
	DataFile = "satisfactory_recipies.json"
)

func main() {
	if !file.FileExists(DataFile) {
		var data string
		var err error

		// Scrape recipes from remote resource
		if data, err = scraper.FetchRecipes(); err != nil {
			panic(fmt.Sprintf("Failed to FetchRecipes: %v", err))
		}

		// Write json data to file
		if _, err = file.WriteStringToFile(DataFile, data); err != nil {
			panic(fmt.Sprintf("Failed to write to file: %v", err))
		}
	}

	p := tea.NewProgram(tui.RootScreen(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
}
