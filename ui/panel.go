package ui

import tea "github.com/charmbracelet/bubbletea"

type panel interface {
	Init() tea.Cmd
	Update(msg tea.Msg) tea.Cmd
	Render() string
}
