package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type view int

const (
	viewWeather view = iota
	viewPomodoro
	viewSysmon
	viewChat
	viewSettings
)

var navItems = []struct {
	label string
	id    view
}{
	{"Weather", viewWeather},
	{"Pomodoro", viewPomodoro},
	{"Sys Monitor", viewSysmon},
	{"Chat", viewChat},
	{"Settings", viewSettings},
}

var (
	sidebarStyle = lipgloss.NewStyle().
			Width(18).
			Padding(1, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240"))

	activeItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true)

	itemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("250"))

	contentStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240"))

	focusedBorder = lipgloss.Color("205")
	blurredBorder = lipgloss.Color("240")
)

type model struct {
	state        appState
	config       Config
	nameInput    textinput.Model
	activeView   view
	cursor       int
	sidebarFocus bool
	width        int
	height       int
}

func initialModel(cfg Config, firstBoot bool) model {
	m := model{
		config:       cfg,
		activeView:   viewWeather,
		sidebarFocus: true,
	}
	if firstBoot {
		m.state = stateSetup
		m.nameInput = newNameInput()
	} else {
		m.state = stateMain
	}
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case configSavedMsg:
		m.state = stateMain
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if m.state == stateSetup {
			return m.updateSetup(msg)
		}
		return m.updateMain(msg)
	}
	return m, nil
}

func (m model) updateSetup(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if msg.String() == "enter" && len(strings.TrimSpace(m.nameInput.Value())) > 0 {
		m.config.Name = strings.TrimSpace(m.nameInput.Value())
		return m, saveConfigCmd(m.config)
	}
	var cmd tea.Cmd
	m.nameInput, cmd = m.nameInput.Update(msg)
	return m, cmd
}

func (m model) updateMain(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q":
		return m, tea.Quit
	case "tab":
		m.sidebarFocus = !m.sidebarFocus
	case "up", "k":
		if m.sidebarFocus && m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.sidebarFocus && m.cursor < len(navItems)-1 {
			m.cursor++
		}
	case "enter":
		if m.sidebarFocus {
			m.activeView = navItems[m.cursor].id
			m.sidebarFocus = false
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.state == stateSetup {
		return renderSetup(m)
	}
	sidebar := m.renderSidebar()
	content := m.renderContent()
	return lipgloss.JoinHorizontal(lipgloss.Top, sidebar, content)
}

func (m model) renderSidebar() string {
	var sb strings.Builder
	for i, item := range navItems {
		if i == m.cursor {
			sb.WriteString(activeItemStyle.Render("> " + item.label))
		} else {
			sb.WriteString(itemStyle.Render("  " + item.label))
		}
		sb.WriteByte('\n')
	}
	s := sb.String()

	borderColor := blurredBorder
	if m.sidebarFocus {
		borderColor = focusedBorder
	}

	return sidebarStyle.BorderForeground(borderColor).Render(s)
}

func (m model) renderContent() string {
	var body string
	switch m.activeView {
	case viewWeather:
		body = "Weather — coming soon"
	case viewPomodoro:
		body = "Pomodoro — coming soon"
	case viewSysmon:
		body = "Sys Monitor — coming soon"
	case viewChat:
		body = "Chat — coming soon"
	}

	borderColor := blurredBorder
	if !m.sidebarFocus {
		borderColor = focusedBorder
	}

	contentW := max(m.width-24, 20)

	return contentStyle.
		BorderForeground(borderColor).
		Width(contentW).
		Height(m.height - 4).
		Render(body)
}

func main() {
	cfg, firstBoot, err := loadConfig()
	if err != nil {
		fmt.Printf("error loading config: %v\n", err)
		os.Exit(1)
	}
	p := tea.NewProgram(initialModel(cfg, firstBoot), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
