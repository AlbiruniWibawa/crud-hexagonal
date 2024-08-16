package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds the configuration details
type Config struct {
	MongoDB struct {
		DSN    string `yaml:"dsn"`
		DBName string `yaml:"dbname"`
	} `yaml:"mongo"`

	App struct {
		Port string `yaml:"port"`
	} `yaml:"app"`
}

// LoadConfig reads a YAML file and unmarshals it into a Config struct
func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
