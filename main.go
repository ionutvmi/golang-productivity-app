package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type model struct {
	width  int
	height int
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m model) View() string {
	var bordersWidth = 2

	var titleStyle = lipgloss.
		NewStyle().
		Width(m.width).
		Padding(1).
		Align(lipgloss.Center)

	var title = titleStyle.Render("Productivity app")

	var titleHeight = lipgloss.Height(title)

	var panelStyle = lipgloss.
		NewStyle().
		Width(m.width/2-bordersWidth).
		Height(m.height/2-bordersWidth-titleHeight/2).
		BorderStyle(lipgloss.RoundedBorder()).
		Align(lipgloss.Center, lipgloss.Center)

	var topPanels = lipgloss.JoinHorizontal(lipgloss.Top, panelStyle.Render("Hello top left"), panelStyle.Render("Hello top right"))
	var bottomPanels = lipgloss.JoinHorizontal(lipgloss.Top, panelStyle.Render("Hello bottom left"), panelStyle.Render("Hello bottom right"))

	var s string
	s += lipgloss.JoinVertical(lipgloss.Left, title, topPanels, bottomPanels)
	// s += strconv.Itoa(m.width) + " / " + strconv.Itoa(m.height)

	// Send the UI for rendering
	return s
}

func initialModel() model {
	return model{}
}
