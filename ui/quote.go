package ui

import (
	"app/provider"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type QuotePanelTickMsg struct{}

type quotePanel struct {
	provider provider.QuoteProvider
}

func NewQuotePanel() *quotePanel {
	return &quotePanel{}
}

func (p *quotePanel) Init() tea.Cmd {
	p.provider = *provider.NewQuoteProvider("https://uselessfacts.jsph.pl/random.json?language=en")
	return func() tea.Msg {
		p.provider.Refresh()
		return QuotePanelTickMsg{}
	}
}

func (p *quotePanel) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case QuotePanelTickMsg:
		return tea.Tick(10*time.Second, func(now time.Time) tea.Msg {
			p.provider.Refresh()
			return QuotePanelTickMsg{}
		})
	}
	return nil
}

func (p *quotePanel) Render() string {
	return p.provider.Message()
}
