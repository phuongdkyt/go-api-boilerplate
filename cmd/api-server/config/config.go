package config

import (
	"fmt"

	"github.com/anhnmt/go-api-boilerplate/internal/pkg/config"
)

type Config struct {
	Log      config.Log      `mapstructure:"log"`
	Postgres config.Postgres `mapstructure:"postgres"`
	Redis    config.Redis    `mapstructure:"redis"`
	Server   config.Server   `mapstructure:"server"`
	JWT      config.JWT      `mapstructure:"jwt"`
	Crypto   config.Crypto   `mapstructure:"crypto"`
}

func New() (Config, error) {
	cfg := Config{}

	configFile, err := config.FilePath()
	if err != nil {
		return cfg, err
	}

	err = config.Load(configFile, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("read config error: %w", err)
	}

	return cfg, nil
}
