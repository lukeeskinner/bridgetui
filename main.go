package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"bridge-tui/app"
	"bridge-tui/config"
)

func main() {
	cfg, firstBoot, err := config.Load()
	if err != nil {
		fmt.Printf("error loading config: %v\n", err)
		os.Exit(1)
	}
	p := tea.NewProgram(app.New(cfg, firstBoot), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
