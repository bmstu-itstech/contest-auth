package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"

	"github.com/bmstu-itstech/contest-auth/config/yaml"
)

// Config holds the configuration for the application.
type Config struct {
	GRPC     yaml.Server   `validate:"required" yaml:"grpc"`
	Postgres yaml.Postgres `validate:"required" yaml:"postgres"`
	Redis    yaml.Redis    `validate:"required" yaml:"redis"`
}

// LoadConfig reads and parses the configuration from a file specified by the CONFIG_PATH environment variable.
func LoadConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.Wrap(err, "cannot read config")
	}

	return &cfg, nil
}
