package actions

import (
	"gopilot/internal/providers"
)

func init() {
	Register("edit-code", &EditCodeAction{})
}

type EditCodeAction struct{}

func (a *EditCodeAction) PreHook(input interface{}, history []providers.Message) (interface{}, []providers.Message, error) {
	// Add system message to guide the model's response format
	systemMsg := providers.Message{
		Role: "system",
		Content: `You are a code editor. When asked to make changes to code:
1. Analyze the requested changes
2. Respond with only the modified code sections
3. Use ... to indicate unchanged code
4. Include brief comments explaining the changes`,
	}

	newHistory := append([]providers.Message{systemMsg}, history...)
	return input, newHistory, nil
}

func (a *EditCodeAction) PostHook(response string) (string, error) {
	// Could add additional formatting or validation here
	return response, nil
}
