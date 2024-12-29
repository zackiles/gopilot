package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Provider string `json:"PROVIDER" yaml:"PROVIDER"`
	APIKey   string `json:"API_KEY" yaml:"API_KEY"`
	Model    string `json:"MODEL" yaml:"MODEL"`
}

func Load(configPath string) (*Config, error) {
	cfg := &Config{}

	// Try loading from specified config file
	if configPath != "" {
		return loadFromFile(configPath)
	}

	// Try loading from .env file first
	if err := loadDotEnv(); err == nil {
		if cfg := loadFromEnv(); isValidConfig(cfg) {
			return cfg, nil
		}
	}

	// Try loading from default config files
	if cfg, err := loadFromDefaultFiles(); err == nil {
		return cfg, nil
	}

	// Fall back to environment variables
	return loadFromEnv(), nil
}

func loadFromFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	ext := filepath.Ext(path)

	switch ext {
	case ".json":
		err = json.Unmarshal(data, cfg)
	case ".yml", ".yaml":
		err = yaml.Unmarshal(data, cfg)
	default:
		err = json.Unmarshal(data, cfg)
	}

	return cfg, err
}

func loadFromDefaultFiles() (*Config, error) {
	files := []string{
		"gopilot.config.json",
		"gopilot.config.yml",
		"gopilot.config.yaml",
	}

	for _, file := range files {
		if cfg, err := loadFromFile(file); err == nil {
			return cfg, nil
		}
	}

	return nil, os.ErrNotExist
}

func loadFromEnv() *Config {
	return &Config{
		Provider: os.Getenv("PROVIDER"),
		APIKey:   os.Getenv("API_KEY"),
		Model:    os.Getenv("MODEL"),
	}
}

func loadDotEnv() error {
	data, err := os.ReadFile(".env")
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, `"'`)
		os.Setenv(key, value)
	}

	return nil
}

func isValidConfig(cfg *Config) bool {
	return cfg.Provider != "" && cfg.APIKey != ""
}
