package chat

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopilot/internal/config"
	"gopilot/internal/providers"
)

type Options struct {
	Stream  bool
	NewChat bool
	OneShot bool
}

type Message struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
}

type Session struct {
	provider    providers.Provider
	history     []Message
	historyFile string
}

func NewSession(cfg *config.Config) *Session {
	provider, err := providers.New(cfg.Provider, cfg.APIKey, cfg.Model)
	if err != nil {
		fmt.Printf("Warning: Failed to initialize provider: %v\n", err)
	}

	s := &Session{
		provider:    provider,
		historyFile: getHistoryFilePath(),
	}

	s.loadHistory()
	return s
}

func (s *Session) AddContext(context string) {
	s.history = append(s.history, Message{
		Role:    "system",
		Content: context,
	})
}

func (s *Session) Send(input interface{}, opts Options) (string, error) {
	if opts.NewChat {
		s.history = nil
	}

	if !opts.OneShot {
		s.history = append(s.history, Message{
			Role:    "user",
			Content: input,
		})
	}

	response, err := s.provider.Send(input, opts.Stream)
	if err != nil {
		return "", err
	}

	if !opts.OneShot {
		s.history = append(s.history, Message{
			Role:    "assistant",
			Content: response,
		})
		s.saveHistory()
	}

	return response, nil
}

func getHistoryFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ".gopilot_history.json"
	}
	return filepath.Join(homeDir, ".gopilot_history.json")
}

func (s *Session) loadHistory() {
	data, err := os.ReadFile(s.historyFile)
	if err != nil {
		return
	}

	json.Unmarshal(data, &s.history)
}

func (s *Session) saveHistory() {
	data, err := json.Marshal(s.history)
	if err != nil {
		return
	}

	os.WriteFile(s.historyFile, data, 0644)
}