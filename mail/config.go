package mail

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds the mail settings loaded from config.yaml.
type Config struct {
	SMTPUsername string `yaml:"smtp_username"`
	SMTPPassword string `yaml:"smtp_password"`
	FromName     string `yaml:"from_name"`
	FromAddress  string `yaml:"from_address"`
	ToAddress    string `yaml:"to_address"`
}

// LoadConfig reads and parses the YAML config file at path.
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &cfg, nil
}
