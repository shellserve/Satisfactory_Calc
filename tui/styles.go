package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// colors #
var (
	LightWhite   = "#FFFFFF"
	FicsitOrange = "#FA9549"
)

// generic styles
var (
	MainStyle     = lipgloss.NewStyle().MarginLeft(2)
	SubTitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#87af00"))
	DotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("##af00ff ")).Render(" • ")
	KeywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#87d7ff"))
	CheckboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#34f"))
)
