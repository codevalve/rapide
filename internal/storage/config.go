package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	RemoteURL    string `json:"remote_url,omitempty"`
	AutoSync     bool   `json:"auto_sync"`
	AutoHideDays int    `json:"auto_hide_days"`
}

func GetConfigPath() (string, error) {
	if custom := os.Getenv("RAPIDE_FILE"); custom != "" {
		return filepath.Join(filepath.Dir(custom), "config.json"), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".rapide", "config.json"), nil
}

func LoadConfig() (*Config, error) {
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{AutoSync: false, AutoHideDays: 14}, nil
		}
		return nil, err
	}

	cfg := Config{
		AutoHideDays: 14, // Default value as per Issue #16
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func SaveConfig(cfg *Config) error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
