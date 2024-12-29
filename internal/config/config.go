package config

import (
	"encoding/json"
	"os"
	"path/filepath"

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
