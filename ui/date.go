package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type DatePanelTickMsg struct{}

type datePanel struct {
	now time.Time
}

func NewDatePanel() *datePanel {
	return &datePanel{}
}

func (d *datePanel) Init() tea.Cmd {
	d.now = time.Now()

	return func() tea.Msg {
		return DatePanelTickMsg{}
	}
}

func (d *datePanel) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case DatePanelTickMsg:
		return tea.Tick(1*time.Second, func(now time.Time) tea.Msg {
			d.now = now
			return DatePanelTickMsg{}
		})
	}
	return nil
}

func (d *datePanel) Render() string {
	return "" +
		d.now.Format("Monday, 02-January-2006") +
		"\n" +
		d.now.Format("15:04:05")
}
