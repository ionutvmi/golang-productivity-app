package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Application struct {
	width   int
	height  int
	panels  []panel
	program *tea.Program
}

type ConfigUpdatedMsg struct{}

func NewApplication() *Application {
	var a = &Application{}
	a.program = tea.NewProgram(a, tea.WithAltScreen())
	return a
}

func (a *Application) Start() {
	if _, err := a.program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (a *Application) Init() tea.Cmd {
	a.panels = append(a.panels, NewDatePanel())

	a.panels = append(a.panels, NewQuotePanel("quote"))

	a.panels = append(a.panels, NewPomodoroPanel())

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

	var appCmds = []tea.Cmd{}

	// panel specific updates
	for _, panel := range a.panels {
		appCmds = append(appCmds, panel.Update(msg))
	}

	return a, tea.Batch(appCmds...)
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
		Align(lipgloss.Center, lipgloss.Center).
		Padding(1)

	var topPanels = lipgloss.JoinHorizontal(lipgloss.Top,
		panelStyle.Render(a.panels[0].Render()),
		panelStyle.Render(a.panels[1].Render()),
	)

	var bottomPanels = lipgloss.JoinHorizontal(lipgloss.Top,
		panelStyle.Render(a.panels[2].Render()),
		panelStyle.Render("Hello bottom right"),
	)

	return lipgloss.JoinVertical(lipgloss.Left, title, topPanels, bottomPanels)
}

func (a *Application) HandleConfigChange() {
	a.program.Send(ConfigUpdatedMsg{})
}
