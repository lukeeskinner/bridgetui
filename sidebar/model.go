package sidebar

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Item struct {
	Label string
}

type Model struct {
	items    []Item
	cursor   int
	focus    bool
	greeting string
}

func New(items []Item, greeting string) Model {
	return Model{items: items, focus: true, greeting: greeting}
}

func (m Model) Cursor() int { return m.cursor }
func (m Model) Focus() bool { return m.focus }

func (m Model) SetFocus(f bool) Model {
	m.focus = f
	return m
}

func (m Model) WithGreeting(name string) Model {
	m.greeting = name
	return m
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	var sb strings.Builder
	sb.WriteString(greetingStyle.Render("Hey, " + m.greeting))
	sb.WriteString("\n\n")
	for i, item := range m.items {
		if i == m.cursor {
			sb.WriteString(activeItemStyle.Render("> " + item.Label))
		} else {
			sb.WriteString(itemStyle.Render("  " + item.Label))
		}
		sb.WriteByte('\n')
	}

	border := blurredBorder
	if m.focus {
		border = focusedBorder
	}
	return boxStyle.BorderForeground(border).Render(sb.String())
}
