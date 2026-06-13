package app

import "github.com/charmbracelet/lipgloss"

var (
	FocusedBorder = lipgloss.Color("205")
	BlurredBorder = lipgloss.Color("240")

	contentStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder())
)
