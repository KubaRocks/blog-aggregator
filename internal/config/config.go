package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	configPath, err := getConfigPathFilePath()

	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(configPath)

	if err != nil {
		return Config{}, err
	}

	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c *Config) SetUser(u string) error {
	c.CurrentUsername = u
	return write(*c)
}

func getConfigPathFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		// If we can't even find HOME, we're stuck. Wrap the error for context.
		return "", fmt.Errorf("cannot find user home directory: %w", err)
	}
	return filepath.Join(home, configFileName), nil
}

func write(cfg Config) error {
	configPath, err := getConfigPathFilePath()

	if err != nil {
		return fmt.Errorf("cannot find user home directory: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")

	if err != nil {
		return fmt.Errorf("failed to marshal config to JSON: %w", err)
	}

	err = os.WriteFile(configPath, data, 0600)

	if err != nil {
		return fmt.Errorf("failed to write config file at %s, %w", configPath, err)
	}

	return nil
}
