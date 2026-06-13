package app

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"bridge-tui/config"
	"bridge-tui/sidebar"
)

type panelID int

const (
	panelWeather panelID = iota
	panelMonitor
	panelGit
	panelJournal
	panelChat
)

var navItems = []sidebar.Item{
	{Label: "Weather"},
	{Label: "Monitor"},
	{Label: "Git"},
	{Label: "Journal"},
	{Label: "Chat"},
}

type Model struct {
	state     appState
	cfg       config.Config
	nameInput textinput.Model
	sidebar   sidebar.Model
	active    panelID
	width     int
	height    int
}

func New(cfg config.Config, firstBoot bool) Model {
	m := Model{
		cfg:     cfg,
		sidebar: sidebar.New(navItems, cfg.Name),
	}
	if firstBoot {
		m.state = stateSetup
		m.nameInput = newNameInput()
	} else {
		m.state = stateMain
	}
	return m
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case configSavedMsg:
		m.state = stateMain
		m.sidebar = m.sidebar.WithGreeting(m.cfg.Name)
		return m, nil

	case tea.KeyMsg:
		if msg.String() == KeyQuit {
			return m, tea.Quit
		}
		if m.state == stateSetup {
			return m.updateSetup(msg)
		}
		return m.updateMain(msg)
	}
	return m, nil
}

func (m Model) updateMain(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.sidebar.Focus() {
		switch msg.String() {
		case KeyQuit2:
			return m, tea.Quit
		case KeyTab:
			m.sidebar = m.sidebar.SetFocus(false)
		case KeyUp, KeyUpK, KeyDown, KeyDownJ:
			var cmd tea.Cmd
			m.sidebar, cmd = m.sidebar.Update(msg)
			return m, cmd
		case KeyEnter:
			m.active = panelID(m.sidebar.Cursor())
			m.sidebar = m.sidebar.SetFocus(false)
		}
	} else {
		switch msg.String() {
		case KeyQuit2:
			return m, tea.Quit
		case KeyTab:
			m.sidebar = m.sidebar.SetFocus(true)
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.state == stateSetup {
		return renderSetup(m)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, m.sidebar.View(), m.renderContent())
}

func (m Model) renderContent() string {
	var body string
	switch m.active {
	case panelWeather:
		body = "Weather — coming soon"
	case panelMonitor:
		body = "Monitor — coming soon"
	case panelGit:
		body = "Git — coming soon"
	case panelJournal:
		body = "Journal — coming soon"
	case panelChat:
		body = "Chat — coming soon"
	}

	border := BlurredBorder
	if !m.sidebar.Focus() {
		border = FocusedBorder
	}

	return contentStyle.
		BorderForeground(border).
		Width(max(m.width-24, 20)).
		Height(max(m.height-4, 5)).
		Render(body)
}
