package ui

import (
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// type pomodoroPanelTickMsg struct{}

type pomodoroPanel struct {
	setTimeButton *Button
	startButton   *Button
	stopButton    *Button

	timer timer.Model
}

func NewPomodoroPanel() *pomodoroPanel {
	return &pomodoroPanel{
		setTimeButton: NewButton("pomodoroSetTime", "Set time", ButtonSecondary),
		startButton:   NewButton("pomodoroStart", "Start", ButtonPrimary),
		stopButton:    NewButton("pomodoroStop", "Stop", ButtonDanger),
		timer:         timer.Model{Timeout: 25 * time.Minute},
	}
}

func (d *pomodoroPanel) Init() tea.Cmd {
	return nil
}

func (d *pomodoroPanel) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft {
			if d.setTimeButton.InBounds(msg) {
				log.Println("clicked on set time")
			}
			if d.startButton.InBounds(msg) {
				d.timer = timer.New(10 * time.Second)
				return d.timer.Init()
			}
			if d.stopButton.InBounds(msg) {
				return d.timer.Stop()
			}
		}
	case timer.TickMsg:
		var cmd tea.Cmd
		d.timer, cmd = d.timer.Update(msg)

		return cmd

	case timer.StartStopMsg:
		var cmd tea.Cmd
		d.timer, cmd = d.timer.Update(msg)
		return cmd

	case timer.TimeoutMsg:
		log.Println("timer timeout")

		// case pomodoroPanelTickMsg:
		// 	return tea.Tick(1*time.Second, func(now time.Time) tea.Msg {
		// 		d.now = now
		// 		return pomodoroPanelTickMsg{}
		// 	})
	}
	return nil
}

func (d *pomodoroPanel) Render() string {
	var seconds = int(d.timer.Timeout.Seconds())

	var timerDisplay = "" +
		fmt.Sprintf("%02d", seconds/60) +
		":" +
		fmt.Sprintf("%02d", seconds%60)

	var buttons = []string{d.setTimeButton.Render()}

	if d.timer.Running() {
		buttons = append(buttons, d.stopButton.Render())
	} else {
		buttons = append(buttons, d.startButton.Render())
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		timerDisplay,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			buttons...,
		),
	)
}
