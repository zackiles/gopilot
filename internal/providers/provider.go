package providers

type Provider interface {
	Send(message interface{}, stream bool) (string, error)
}

func New(providerName, apiKey, model string) (Provider, error) {
	switch providerName {
	case "openai":
		return NewOpenAI(apiKey, model)
	case "anthropic":
		return NewAnthropic(apiKey, model)
	case "cohere-ai":
		return NewCohere(apiKey, model)
	case "openrouter":
		return NewOpenRouter(apiKey, model)
	default:
		return NewOpenAI(apiKey, model) // Default to OpenAI
	}
}
