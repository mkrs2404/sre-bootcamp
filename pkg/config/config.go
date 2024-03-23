package config

import (
	"log/slog"

	"github.com/caarlos0/env/v9"
	"github.com/pkg/errors"
)

type Config struct {
	ServerPort    string `env:"SERVER_PORT" envDefault:"9090" json:"server_port"`
	DSN           string `env:"DSN" envDefault:"postgres://postgres:postgres@localhost:5435/bootcamp?sslmode=disable" json:"dsn"`
	AppName       string `env:"APP_NAME" envDefault:"CRUD" json:"app_name"`
	Environment   string `env:"ENVIRONMENT" envDefault:"development" json:"environment"`
	MigrationPath string `env:"MIGRATION_PATH" envDefault:"db/migrations" json:"migration_path"`
	APIKey        string `env:"API_KEY" envDefault:"" json:"api_key"`
}

var cfg *Config

func Get() *Config {
	var err error
	if cfg == nil {
		cfg, err = ReadConfig()
		if err != nil {
			panic("error reading config")
		}
	}
	return cfg
}

func ReadConfig() (*Config, error) {
	opts := env.Options{}

	var cfg Config
	err := env.ParseWithOptions(&cfg, opts)
	if err != nil {
		errW := errors.Wrap(err, "could not parse configuration")
		if !(cfg.Environment == "development") {
			return nil, errW
		}
		slog.Error(errW.Error())
	}
	return &cfg, nil
}
