package provider

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type QuoteProvider struct {
	url  string
	data *quote
}

type quote struct {
	Text   string `json:"text"`
	Source string `json:"source"`
}

func NewQuoteProvider(url string) *QuoteProvider {
	return &QuoteProvider{
		url: url,
		data: &quote{
			Text: "Loading...",
		},
	}
}

func (p *QuoteProvider) Message() string {
	var msg = p.data.Text

	if p.data.Source != "" {
		msg += " - " + p.data.Source
	}

	return msg
}

func (p *QuoteProvider) SetUrl(val string) {
	p.url = val
}

func (p *QuoteProvider) Refresh() error {
	p.data.Text = ""

	var client = &http.Client{
		Timeout: 10 * time.Second,
	}

	var res, err = client.Get(p.url)

	if err != nil {
		p.data.Text = err.Error()
		return err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		p.data.Text = err.Error()
		return err
	}

	if err = json.Unmarshal(data, p.data); err != nil {
		p.data.Text = err.Error()
		return err
	}

	return nil
}
