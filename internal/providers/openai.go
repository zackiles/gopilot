package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

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

func (o *OpenAI) Send(history []Message, message interface{}, stream bool) (string, error) {
	messages := make([]openai.ChatCompletionMessage, len(history)+1)

	// Convert history to OpenAI format
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

	// Add current message
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

func (o *OpenAI) SupportsStreaming() bool {
	return true
}

func (o *OpenAI) HandleRateLimiting(err error) error {
	// Implement exponential backoff if rate limited
	return err
}

func (o *OpenAI) handleStreamingResponse(messages []openai.ChatCompletionMessage) (string, error) {
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

func (o *OpenAI) handleSingleResponse(messages []openai.ChatCompletionMessage) (string, error) {
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
