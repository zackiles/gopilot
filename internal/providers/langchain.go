package providers

import (
	"fmt"
)

type Langchain struct {
	apiKey string
	model  string
}

func NewLangchain(apiKey, model string) (*Langchain, error) {
	return &Langchain{
		apiKey: apiKey,
		model:  model,
	}, nil
}

func (l *Langchain) Send(history []Message, message interface{}, stream bool) (string, error) {
	return "", fmt.Errorf("langchain provider not yet implemented - requires langchain-go implementation")
}

func (l *Langchain) SupportsStreaming() bool {
	return false
}

func (l *Langchain) HandleRateLimiting(err error) error {
	return err
}
