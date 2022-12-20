package ui

import (
	"app/config"
	"app/provider"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type QuotePanelTickMsg struct{}

type quotePanel struct {
	configNs string
	provider provider.QuoteProvider
}

func NewQuotePanel(configNs string) *quotePanel {
	return &quotePanel{
		configNs: configNs,
	}
}

func (p *quotePanel) url() string {
	return config.GetString(p.configNs + ".url")
}

func (p *quotePanel) Init() tea.Cmd {
	p.provider = *provider.NewQuoteProvider(p.url())

	return func() tea.Msg {
		p.provider.Refresh()
		return QuotePanelTickMsg{}
	}
}

func (p *quotePanel) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case ConfigUpdatedMsg:
		p.provider.SetUrl(p.url())
		p.provider.Refresh()

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
