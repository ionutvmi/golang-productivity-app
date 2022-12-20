package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Application struct {
	width  int
	height int
	panels []panel
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Start() {
	var program = tea.NewProgram(a, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (a *Application) Init() tea.Cmd {
	a.panels = append(a.panels, NewDatePanel())

	var appCmds = []tea.Cmd{}

	for _, panel := range a.panels {
		appCmds = append(appCmds, panel.Init())
	}

	return tea.Batch(appCmds...)
}

func (a *Application) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// log.Println("Updated", msg)

	// application level messages
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return a, tea.Quit
		}

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
	}

	// panel specific messages
	for _, panel := range a.panels {
		panel.Update(msg)
	}

	return a, nil
}

func (a *Application) View() string {
	var bordersWidth = 2

	var titleStyle = lipgloss.
		NewStyle().
		Width(a.width).
		Padding(1).
		Bold(true).
		Align(lipgloss.Center)

	var title = titleStyle.Render("Productivity App")
	var titleHeight = lipgloss.Height(title)

	var panelHeight = a.height/2 - bordersWidth - titleHeight/2
	var panelWidth = a.width/2 - bordersWidth
	var panelStyle = lipgloss.
		NewStyle().
		Width(panelWidth).
		Height(panelHeight).
		BorderStyle(lipgloss.RoundedBorder()).
		Align(lipgloss.Center, lipgloss.Center)

	var topPanels = lipgloss.JoinHorizontal(lipgloss.Top,
		panelStyle.Render(a.panels[0].Render()),
		panelStyle.Render("Hello top right"),
	)

	var bottomPanels = lipgloss.JoinHorizontal(lipgloss.Top,
		panelStyle.Render("Hello bottom left"),
		panelStyle.Render("Hello bottom right"),
	)

	return lipgloss.JoinVertical(lipgloss.Left, title, topPanels, bottomPanels)
}
