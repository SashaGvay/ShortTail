package config

import (
	"fmt"

	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Conf struct {
	ENV           string `env:"ENV" envDefault:"PROD"`
	HttpAddr      string `env:"HTTP_ADDR" envDefault:"0.0.0.0:8080"`
	BadgetDataDir string `env:"BADGER_DB_PATH" envDefault:"./badger_data"`
}

func NewConfig() (*Conf, error) {
	cfg := &Conf{}

	godotenv.Load()
	if err := env.Parse(cfg, env.Options{}); err != nil {
		return nil, fmt.Errorf("env.Parse: %w", err)
	}
	return cfg, nil
}
