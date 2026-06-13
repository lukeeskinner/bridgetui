package sidebar

import "github.com/charmbracelet/lipgloss"

var (
	focusedBorder = lipgloss.Color("205")
	blurredBorder = lipgloss.Color("240")

	boxStyle = lipgloss.NewStyle().
			Width(18).
			Padding(1, 1).
			Border(lipgloss.RoundedBorder())

	activeItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true)

	itemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("250"))

	greetingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true)
)
