// Package satisfactorycalc is for caling
package main

import (
	"fmt"
	"os"

	"github.com/shellserve/Satisfactory_Calc/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.RootScreen(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
}
