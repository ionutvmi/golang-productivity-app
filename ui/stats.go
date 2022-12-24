package ui

import (
	"app/database"
	"app/models"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type updatePomodoroStats struct{}

type statsPanel struct {
	stats *models.PomodoroStats
}

func NewStatsPanel() *statsPanel {
	return &statsPanel{}
}

func (p *statsPanel) Init() tea.Cmd {
	return func() tea.Msg {
		p.stats = database.PomodoroStats()
		return updatePomodoroStats{}
	}
}

func (p *statsPanel) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	}
	return nil
}

func (p *statsPanel) Render() string {

	if p.stats == nil {
		return lipgloss.JoinVertical(
			lipgloss.Center,
			"Stats",
			"Loading...",
		)
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		"Stats",
		"Today: "+strconv.Itoa(p.stats.Today),
		"This Week: "+strconv.Itoa(p.stats.Week),
		"This Month: "+strconv.Itoa(p.stats.Month),
		"This Year: "+strconv.Itoa(p.stats.Year),
	)
}
