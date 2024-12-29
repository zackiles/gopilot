package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type OpenRouter struct {
	client *openai.Client
	model  string
}

func NewOpenRouter(apiKey, model string) (*OpenRouter, error) {
	if model == "" {
		model = "openai/gpt-3.5-turbo"
	}

	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://openrouter.ai/api/v1"

	client := openai.NewClientWithConfig(config)
	return &OpenRouter{
		client: client,
		model:  model,
	}, nil
}

func (o *OpenRouter) Send(history []Message, message interface{}, stream bool) (string, error) {
	messages := make([]openai.ChatCompletionMessage, len(history)+1)

	for i, msg := range history {
		content := ""
		switch v := msg.Content.(type) {
		case string:
			content = v
		default:
			jsonBytes, _ := json.Marshal(v)
			content = string(jsonBytes)
		}

		messages[i] = openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: content,
		}
	}

	content := ""
	switch v := message.(type) {
	case string:
		content = v
	default:
		jsonBytes, _ := json.Marshal(v)
		content = string(jsonBytes)
	}
	messages[len(messages)-1] = openai.ChatCompletionMessage{
		Role:    "user",
		Content: content,
	}

	if stream {
		return o.handleStreamingResponse(messages)
	}
	return o.handleSingleResponse(messages)
}

func (o *OpenRouter) SupportsStreaming() bool {
	return true
}

func (o *OpenRouter) HandleRateLimiting(err error) error {
	return err
}

func (o *OpenRouter) handleStreamingResponse(messages []openai.ChatCompletionMessage) (string, error) {
	stream, err := o.client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    o.model,
			Messages: messages,
		},
	)
	if err != nil {
		return "", err
	}
	defer stream.Close()

	var fullResponse strings.Builder
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		content := response.Choices[0].Delta.Content
		fmt.Print(content)
		fullResponse.WriteString(content)
	}
	fmt.Println()
	return fullResponse.String(), nil
}

func (o *OpenRouter) handleSingleResponse(messages []openai.ChatCompletionMessage) (string, error) {
	resp, err := o.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    o.model,
			Messages: messages,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
