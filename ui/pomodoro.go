package ui

import (
	"app/database"
	"app/models"
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

	activePomodoro *models.Pomodoro
}

func NewPomodoroPanel() *pomodoroPanel {
	var p = &pomodoroPanel{
		setTimeButton: NewButton("pomodoroSetTime", "Set time", ButtonSecondary),
		startButton:   NewButton("pomodoroStart", "Start", ButtonPrimary),
		stopButton:    NewButton("pomodoroStop", "Stop", ButtonDanger),
		timer:         timer.Model{Timeout: 25 * time.Minute},
		textInput:     textinput.New(),
		action:        pomodoroView,
	}

	p.textInput.Placeholder = "Time in minutes"
	p.textInput.CharLimit = 3
	p.textInput.Width = 0
	p.textInput.SetCursorMode(textinput.CursorStatic)
	p.textInput.SetValue("25")

	p.textInput.Validate = func(s string) error {
		_, err := strconv.Atoi(s)
		return err
	}

	p.setTimeButton.AddOnClick(func() tea.Cmd {
		p.action = pomodoroSetTime
		p.textInput.Focus()
		p.lastTextInputValue = p.textInput.Value()

		return nil
	})

	p.startButton.AddOnClick(func() tea.Cmd {
		p.timer = timer.New(p.timeout())
		p.startPomodoro()
		return p.timer.Init()
	})

	p.stopButton.AddOnClick(func() tea.Cmd {
		return p.timer.Stop()
	})

	return p
}

func (p *pomodoroPanel) Init() tea.Cmd {
	return nil
}

func (p *pomodoroPanel) Update(msg tea.Msg) tea.Cmd {
	var allCmds = []tea.Cmd{}
	var isStartStopMsg = false

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if p.action == pomodoroSetTime {
				p.timer.Timeout = p.timeout()
				p.action = pomodoroView
			}
		case tea.KeyEsc:
			if p.action == pomodoroSetTime {
				p.action = pomodoroView
				p.textInput.SetValue(p.lastTextInputValue)
			}
		}

	case timer.StartStopMsg:
		isStartStopMsg = true

	case timer.TimeoutMsg:
		p.endPomodoro()
	}

	allCmds = append(allCmds, p.setTimeButton.Update(msg))
	allCmds = append(allCmds, p.startButton.Update(msg))
	allCmds = append(allCmds, p.stopButton.Update(msg))

	var cmd tea.Cmd
	p.textInput, cmd = p.textInput.Update(msg)
	allCmds = append(allCmds, cmd)

	p.timer, cmd = p.timer.Update(msg)
	allCmds = append(allCmds, cmd)

	if isStartStopMsg && !p.timer.Running() {
		p.timer.Timeout = p.timeout()
	}

	return tea.Batch(allCmds...)
}

func (p *pomodoroPanel) Render() string {
	if p.action == pomodoroSetTime {
		return lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.NewStyle().Width(30).Align(lipgloss.Center).Render("Set Time"),
			"",
			"",
			"",
			p.textInput.View(),
		)
	}

	var seconds = int(p.timer.Timeout.Seconds())

	var timerDisplay = "" +
		fmt.Sprintf("%02d", seconds/60) +
		":" +
		fmt.Sprintf("%02d", seconds%60)

	var buttons = []string{p.setTimeButton.Render()}

	if p.timer.Running() {
		buttons = append(buttons, p.stopButton.Render())
	} else {
		buttons = append(buttons, p.startButton.Render())
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

func (p *pomodoroPanel) startPomodoro() {
	p.activePomodoro = &models.Pomodoro{
		StartDate: time.Now().UTC(),
		SessionType: models.PomodoroType{
			ID: 1,
		},
	}
}

func (p *pomodoroPanel) endPomodoro() error {
	p.activePomodoro.EndDate = time.Now().UTC()

	var err = database.PomodoroInsert(p.activePomodoro)

	if err != nil {
		log.Printf("Failed to save pomodoro %s", err.Error())
		return err
	}

	return nil
}

func (p *pomodoroPanel) timeout() time.Duration {
	var val, err = strconv.Atoi(p.textInput.Value())

	if err != nil {
		log.Println("Failed to process the time value", err.Error(), p.textInput.Value())
		val = 25
	}

	return time.Duration(val) * time.Second
}
