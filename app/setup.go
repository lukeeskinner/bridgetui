package app

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"bridge-tui/config"
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

func saveConfigCmd(cfg config.Config) tea.Cmd {
	return func() tea.Msg {
		_ = config.Save(cfg)
		return configSavedMsg{}
	}
}

func (m Model) updateSetup(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if msg.String() == KeyEnter && len(strings.TrimSpace(m.nameInput.Value())) > 0 {
		m.cfg.Name = strings.TrimSpace(m.nameInput.Value())
		return m, saveConfigCmd(m.cfg)
	}
	var cmd tea.Cmd
	m.nameInput, cmd = m.nameInput.Update(msg)
	return m, cmd
}

var (
	setupBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(FocusedBorder).
			Padding(2, 4)

	setupTitleStyle = lipgloss.NewStyle().
			Foreground(FocusedBorder).
			Bold(true).
			MarginBottom(1)

	setupHintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			MarginTop(1)
)

func renderSetup(m Model) string {
	title := setupTitleStyle.Render("Welcome to Bridge")
	prompt := "What's your name?\n\n" + m.nameInput.View()
	hint := setupHintStyle.Render("enter to continue · ctrl+c to quit")

	box := setupBoxStyle.Render(fmt.Sprintf("%s\n%s\n%s", title, prompt, hint))

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}
