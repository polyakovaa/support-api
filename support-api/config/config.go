package config

import (
	"os"

	"github.com/caarlos0/env/v8"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DBHost     string `yaml:"db_host"`
	DBPort     int    `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password" env:"DB_PASSWORD"`
	DBName     string `yaml:"db_name"`
	APIPort    int    `yaml:"api_port"`
}

func Load(path string) (*Config, error) {
	cfg := Config{}

	if _, err := os.Stat(path); err == nil {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return nil, err
		}
	}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
