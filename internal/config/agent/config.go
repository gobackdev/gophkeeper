// Пакет config предоставляет конфигурацию для приложения.
package config

import (
	"encoding/json"
	"errors"
	"os"
)

// VersionInfo содержит информацию о версии приложения
type VersionInfo struct {
	Version string `json:"version"`
	Date    string `json:"date"`
}

// AgentConfig конфигурация агента
type AgentConfig struct {
	Version VersionInfo `json:"version"`
}

// NewAgentConfig создает новый экземпляр AgentConfig
func NewAgentConfig() *AgentConfig {
	return &AgentConfig{}
}

// LoadAgentConfig загружает конфигурацию агента из JSON-файла по указанному пути
func LoadAgentConfig(path string) (*AgentConfig, error) {
	if path == "" {
		return NewAgentConfig(), nil
	}

	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return NewAgentConfig(), nil
		}
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	cfg := NewAgentConfig()
	if err := dec.Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
