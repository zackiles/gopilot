package providers

import (
	"context"
	"encoding/json"

	"github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	client *openai.Client
	model  string
}

func NewOpenAI(apiKey, model string) (*OpenAI, error) {
	if model == "" {
		model = "gpt-3.5-turbo"
	}

	client := openai.NewClient(apiKey)
	return &OpenAI{
		client: client,
		model:  model,
	}, nil
}

func (o *OpenAI) Send(message interface{}, stream bool) (string, error) {
	content := ""
	switch v := message.(type) {
	case string:
		content = v
	default:
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		content = string(jsonBytes)
	}

	resp, err := o.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: o.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
