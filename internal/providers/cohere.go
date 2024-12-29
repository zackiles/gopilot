package providers

import (
	"fmt"
)

type Cohere struct {
	apiKey string
	model  string
}

func NewCohere(apiKey, model string) (*Cohere, error) {
	if model == "" {
		model = "command"
	}
	return &Cohere{
		apiKey: apiKey,
		model:  model,
	}, nil
}

func (c *Cohere) Send(history []Message, message interface{}, stream bool) (string, error) {
	return "", fmt.Errorf("cohere provider not yet implemented - requires cohere-go client")
}

func (c *Cohere) SupportsStreaming() bool {
	return false
}

func (c *Cohere) HandleRateLimiting(err error) error {
	return err
}
