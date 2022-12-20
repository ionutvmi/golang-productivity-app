package ui

import tea "github.com/charmbracelet/bubbletea"

type datePanel struct{}

func NewDatePanel() *datePanel {
	return &datePanel{}
}

func (*datePanel) Init() tea.Cmd {
	return nil
}

func (*datePanel) Update(msg tea.Msg) {
}

func (*datePanel) Render() string {
	return "Date panel"
}
