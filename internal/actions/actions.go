package actions

import "gopilot/internal/providers"

type Action interface {
	// PreHook modifies or enhances the input before it's sent to the provider
	PreHook(input interface{}, history []providers.Message) (interface{}, []providers.Message, error)

	// PostHook processes the provider's response before it's returned to the user
	PostHook(response string) (string, error)
}

// Registry stores all available actions
var registry = make(map[string]Action)

// Register adds an action to the registry
func Register(name string, action Action) {
	registry[name] = action
}

// Get retrieves an action from the registry
func Get(name string) (Action, bool) {
	action, exists := registry[name]
	return action, exists
}
