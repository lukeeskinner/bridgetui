package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type appState int

const (
	stateSetup appState = iota
	stateMain
)

type configSavedMsg struct{}

func newNameInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "e.g. Luke"
	ti.CharLimit = 32
	ti.Width = 28
	ti.Focus()
	return ti
}

func saveConfigCmd(cfg Config) tea.Cmd {
	return func() tea.Msg {
		_ = saveConfig(cfg)
		return configSavedMsg{}
	}
}

var (
	setupBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("205")).
			Padding(2, 4)

	setupTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true).
			MarginBottom(1)

	setupHintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			MarginTop(1)
)

func renderSetup(m model) string {
	title := setupTitleStyle.Render("Welcome to Bridge")
	prompt := "What's your name?\n\n" + m.nameInput.View()
	hint := setupHintStyle.Render("enter to continue · ctrl+c to quit")

	box := setupBoxStyle.Render(fmt.Sprintf("%s\n%s\n%s", title, prompt, hint))

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		box,
	)
}
