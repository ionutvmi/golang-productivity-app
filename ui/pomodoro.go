package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// type pomodoroPanelTickMsg struct{}

type pomodoroPanel struct {
}

func NewPomodoroPanel() *pomodoroPanel {
	return &pomodoroPanel{}
}

func (d *pomodoroPanel) Init() tea.Cmd {

	return func() tea.Msg {
		return nil
	}
}

func (d *pomodoroPanel) Update(msg tea.Msg) tea.Cmd {
	// switch msg.(type) {
	// case pomodoroPanelTickMsg:
	// 	return tea.Tick(1*time.Second, func(now time.Time) tea.Msg {
	// 		d.now = now
	// 		return pomodoroPanelTickMsg{}
	// 	})
	// }
	return nil
}

func (d *pomodoroPanel) Render() string {
	var primaryButton = lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#0c450e")).
		Padding(0, 2).
		MarginRight(1).
		Margin(1)

	var secondaryButton = lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#222222")).
		Padding(0, 2).
		MarginRight(1).
		Margin(1)

	var dangerButton = lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#450a0a")).
		Padding(0, 2).
		MarginRight(1).
		Margin(1)

	return lipgloss.JoinVertical(
		lipgloss.Center,
		"25:00",
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			secondaryButton.Render("Set time"),
			primaryButton.Render("Start"),
			dangerButton.Render("Stop"),
		),
	)
}
