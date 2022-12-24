package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Application struct {
	width   int
	height  int
	panels  []panel
	program *tea.Program
	help    help.Model
	keymap  keymap
}

type keymap struct {
	start   key.Binding
	stop    key.Binding
	setTime key.Binding
	quit    key.Binding
	cancel  key.Binding
	save    key.Binding
}

type ConfigUpdatedMsg struct{}

func NewApplication() *Application {
	var a = &Application{
		keymap: keymap{
			start: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "start"),
			),
			stop: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "stop"),
			),
			setTime: key.NewBinding(
				key.WithKeys("t"),
				key.WithHelp("t", "time set"),
			),
			quit: key.NewBinding(
				key.WithKeys("q", "ctrl+c"),
				key.WithHelp("q", "quit"),
			),
			cancel: key.NewBinding(
				key.WithKeys("esc"),
				key.WithHelp("esc", "cancel"),
			),
			save: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "save"),
			),
		},
	}

	a.keymap.stop.SetEnabled(false)
	a.keymap.save.SetEnabled(false)
	a.keymap.cancel.SetEnabled(false)

	a.program = tea.NewProgram(a, tea.WithAltScreen(), tea.WithMouseAllMotion())
	return a
}

func (a *Application) Start() {
	zone.NewGlobal()

	if _, err := a.program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (a *Application) Init() tea.Cmd {
	a.panels = append(a.panels, NewDatePanel())

	a.panels = append(a.panels, NewQuotePanel("quote"))

	a.panels = append(a.panels, NewPomodoroPanel())

	a.panels = append(a.panels, NewStatsPanel())

	var appCmds = []tea.Cmd{}

	for _, panel := range a.panels {
		appCmds = append(appCmds, panel.Init())
	}

	return tea.Batch(appCmds...)
}

func (a *Application) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// log.Println("Updated", msg)

	var appCmds = []tea.Cmd{}

	// application level messages
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, a.keymap.quit):
			return a, tea.Quit
		case key.Matches(msg, a.keymap.start):
			appCmds = append(appCmds, func() tea.Msg {
				return startPomodoroTimerMsg{}
			})
		case key.Matches(msg, a.keymap.stop):
			appCmds = append(appCmds, func() tea.Msg {
				return stopPomodoroTimerMsg{}
			})
		case key.Matches(msg, a.keymap.setTime):
			appCmds = append(appCmds, func() tea.Msg {
				return setTimePomodoroMsg{}
			})

		case key.Matches(msg, a.keymap.cancel):
			appCmds = append(appCmds, func() tea.Msg {
				return cancelSetTimePomodoroMsg{}
			})

		case key.Matches(msg, a.keymap.save):
			appCmds = append(appCmds, func() tea.Msg {
				return saveSetTimePomodoroMsg{}
			})
		}

	case startPomodoroTimerMsg:
		a.keymap.start.SetEnabled(false)

		a.keymap.stop.SetEnabled(true)

	case stopPomodoroTimerMsg:
		a.keymap.stop.SetEnabled(false)

		a.keymap.start.SetEnabled(true)

	case setTimePomodoroMsg:
		a.keymap.stop.SetEnabled(false)
		a.keymap.start.SetEnabled(false)
		a.keymap.setTime.SetEnabled(false)
		a.keymap.quit.SetEnabled(false)

		a.keymap.cancel.SetEnabled(true)
		a.keymap.save.SetEnabled(true)

	case cancelSetTimePomodoroMsg, saveSetTimePomodoroMsg:
		a.keymap.stop.SetEnabled(false)
		a.keymap.start.SetEnabled(true)
		a.keymap.setTime.SetEnabled(true)
		a.keymap.quit.SetEnabled(true)

		a.keymap.cancel.SetEnabled(false)
		a.keymap.save.SetEnabled(false)

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
	}

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

	var help = a.helpView()
	var helpHeight = lipgloss.Height(help)

	var panelHeight = a.height/2 - bordersWidth - titleHeight/2 - helpHeight/2
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
		panelStyle.Render(a.panels[3].Render()),
	)

	var screen = lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		topPanels,
		bottomPanels,
		help,
	)
	return zone.Scan(screen)
}

func (a *Application) helpView() string {
	return a.help.ShortHelpView([]key.Binding{
		a.keymap.start,
		a.keymap.stop,
		a.keymap.setTime,
		a.keymap.save,
		a.keymap.quit,
		a.keymap.cancel,
	})
}

func (a *Application) HandleConfigChange() {
	a.program.Send(ConfigUpdatedMsg{})
}
