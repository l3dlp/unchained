package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type AccountConfig struct {
	Email    string `yaml:"email"`
	ClientID string `yaml:"client_id"`
	Secret   string `yaml:"secret"`
}

func LoadAccount(path string) (*AccountConfig, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var ac AccountConfig
	if err := yaml.Unmarshal(b, &ac); err != nil {
		return nil, err
	}
	return &ac, nil
}

func GetConfigDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, "config"), nil
}
