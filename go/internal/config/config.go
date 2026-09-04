package config

import (
	"os"
)

type Config struct {
	Logs LogConfig
	DB   PostgresConfig
	Port string
}

type LogConfig struct {
	Style string
	Level string
}

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

func LoadConfig() *Config {
	cfg := &Config{
		Port: os.Getenv("PORT"),
		Logs: LogConfig{
			Style: os.Getenv("LOG_STYLE"),
			Level: os.Getenv("LOG_LEVEL"),
		},
		DB: PostgresConfig{
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
		},
	}

	return cfg
}
