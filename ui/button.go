package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type ButtonKind int

const (
	ButtonPrimary ButtonKind = iota
	ButtonSecondary
	ButtonDanger
)

var buttonBaseStyle = lipgloss.
	NewStyle().
	// Border(lipgloss.NormalBorder()).
	Padding(0, 2).
	MarginRight(1).
	Margin(1)

var buttonPrimaryStyle = buttonBaseStyle.
	Copy().
	Background(lipgloss.Color("#0c450e")).
	Foreground(lipgloss.Color("#FFFFFF"))

var buttonSecondaryStyle = buttonBaseStyle.
	Copy().
	Background(lipgloss.Color("#222222")).
	Foreground(lipgloss.Color("#FFFFFF"))

var buttonDangerStyle = buttonBaseStyle.
	Copy().
	Background(lipgloss.Color("#450a0a")).
	Foreground(lipgloss.Color("#FFFFFF"))

type Button struct {
	id    string
	label string
	kind  ButtonKind
}

func NewButton(id string, label string, kind ButtonKind) *Button {
	return &Button{
		id:    id,
		label: label,
		kind:  kind,
	}
}

func (b *Button) InBounds(msg tea.MouseMsg) bool {
	return zone.Get(b.id).InBounds(msg)
}

func (b *Button) Render() string {
	var btn = ""

	switch b.kind {
	case ButtonPrimary:
		btn = buttonPrimaryStyle.Render(b.label)
	case ButtonSecondary:
		btn = buttonSecondaryStyle.Render(b.label)
	case ButtonDanger:
		btn = buttonDangerStyle.Render(b.label)
	}

	return zone.Mark(b.id, btn)
}
