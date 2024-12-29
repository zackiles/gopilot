package providers

import (
	"fmt"
)

type Anthropic struct {
	apiKey string
	model  string
}

func NewAnthropic(apiKey, model string) (*Anthropic, error) {
	if model == "" {
		model = "claude-2"
	}
	return &Anthropic{
		apiKey: apiKey,
		model:  model,
	}, nil
}

func (a *Anthropic) Send(history []Message, message interface{}, stream bool) (string, error) {
	return "", fmt.Errorf("anthropic provider not yet implemented - requires anthropic-go client")
}

func (a *Anthropic) SupportsStreaming() bool {
	return false
}

func (a *Anthropic) HandleRateLimiting(err error) error {
	return err
}
