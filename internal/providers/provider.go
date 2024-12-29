package providers

type Message struct {
	Role    string
	Content interface{}
}

type Provider interface {
	Send(history []Message, message interface{}, stream bool) (string, error)
	SupportsStreaming() bool
	HandleRateLimiting(error) error
}

func New(providerName, apiKey, model string) (Provider, error) {
	switch providerName {
	case "openai":
		return NewOpenAI(apiKey, model)
	case "anthropic":
		return NewAnthropic(apiKey, model)
	case "cohere-ai":
		return NewCohere(apiKey, model)
	case "huggingface":
		return NewHuggingFace(apiKey, model)
	case "langchain":
		return NewLangchain(apiKey, model)
	case "openrouter":
		return NewOpenRouter(apiKey, model)
	default:
		return NewOpenAI(apiKey, model) // Default to OpenAI
	}
}
