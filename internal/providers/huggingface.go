package providers

import (
	"fmt"
)

type HuggingFace struct {
	apiKey string
	model  string
}

func NewHuggingFace(apiKey, model string) (*HuggingFace, error) {
	if model == "" {
		model = "gpt2"
	}
	return &HuggingFace{
		apiKey: apiKey,
		model:  model,
	}, nil
}

func (h *HuggingFace) Send(history []Message, message interface{}, stream bool) (string, error) {
	return "", fmt.Errorf("huggingface provider not yet implemented - requires huggingface-go client")
}

func (h *HuggingFace) SupportsStreaming() bool {
	return false
}

func (h *HuggingFace) HandleRateLimiting(err error) error {
	return err
}
