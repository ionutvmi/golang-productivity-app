package ui

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type pomodoroPanelAction int

const (
	pomodoroView pomodoroPanelAction = iota
	pomodoroSetTime
)

type pomodoroPanel struct {
	setTimeButton *Button
	startButton   *Button
	stopButton    *Button

	timer              timer.Model
	textInput          textinput.Model
	lastTextInputValue string

	action pomodoroPanelAction
}

func NewPomodoroPanel() *pomodoroPanel {
	var panel = &pomodoroPanel{
		setTimeButton: NewButton("pomodoroSetTime", "Set time", ButtonSecondary),
		startButton:   NewButton("pomodoroStart", "Start", ButtonPrimary),
		stopButton:    NewButton("pomodoroStop", "Stop", ButtonDanger),
		timer:         timer.Model{Timeout: 25 * time.Minute},
		textInput:     textinput.New(),
		action:        pomodoroView,
	}

	panel.textInput.Placeholder = "Time in minutes"
	panel.textInput.CharLimit = 3
	panel.textInput.Width = 0
	panel.textInput.SetCursorMode(textinput.CursorStatic)
	panel.textInput.SetValue("25")

	panel.textInput.Validate = func(s string) error {
		_, err := strconv.Atoi(s)
		return err
	}

	return panel
}

func (d *pomodoroPanel) Init() tea.Cmd {
	return nil
}

func (d *pomodoroPanel) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if d.action == pomodoroSetTime {
				d.timer.Timeout = d.timeout()
				d.action = pomodoroView
			}
		case tea.KeyEsc:
			if d.action == pomodoroSetTime {
				d.action = pomodoroView
				d.textInput.SetValue(d.lastTextInputValue)
			}
		}
	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft {
			if d.setTimeButton.InBounds(msg) {
				d.action = pomodoroSetTime
				d.textInput.Focus()
				d.lastTextInputValue = d.textInput.Value()
			}
			if d.startButton.InBounds(msg) {
				d.timer = timer.New(d.timeout())
				return d.timer.Init()
			}
			if d.stopButton.InBounds(msg) {
				return d.timer.Stop()
			}
		}
	case timer.TickMsg:
		d.timer, cmd = d.timer.Update(msg)

		return cmd

	case timer.StartStopMsg:
		d.timer, cmd = d.timer.Update(msg)

		if !d.timer.Running() {
			d.timer.Timeout = d.timeout()
		}

		return cmd

	case timer.TimeoutMsg:
		log.Println("timer timeout")

		// case pomodoroPanelTickMsg:
		// 	return tea.Tick(1*time.Second, func(now time.Time) tea.Msg {
		// 		d.now = now
		// 		return pomodoroPanelTickMsg{}
		// 	})
	}

	d.textInput, cmd = d.textInput.Update(msg)

	return cmd
}

func (d *pomodoroPanel) Render() string {
	if d.action == pomodoroSetTime {
		return lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.NewStyle().Width(30).Align(lipgloss.Center).Render("Set Time"),
			"",
			"",
			"",
			d.textInput.View(),
		)
	}

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
		lipgloss.NewStyle().Width(30).Align(lipgloss.Center).Render("Pomodoro Timer"),
		"",
		timerDisplay,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			buttons...,
		),
	)
}

func (d *pomodoroPanel) timeout() time.Duration {
	var val, err = strconv.Atoi(d.textInput.Value())

	if err != nil {
		log.Println("Failed to process the time value", err.Error(), d.textInput.Value())
		val = 25
	}

	return time.Duration(val) * time.Minute
}
