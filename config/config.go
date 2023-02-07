package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
	}

	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}

	Log struct {
		Level string `yaml:"log_level"`
	}

	PG struct {
		PoolMax int    `yaml:"pool_max"`
		URL     string `yaml:"url"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	yamlFile, err := os.ReadFile("./config/config.yml")
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
