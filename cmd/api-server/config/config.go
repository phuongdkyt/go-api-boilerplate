package config

import (
	"github.com/anhnmt/go-api-boilerplate/internal/pkg/casbin"
	"github.com/anhnmt/go-api-boilerplate/internal/pkg/config"
	"github.com/anhnmt/go-api-boilerplate/internal/pkg/logger"
	"github.com/anhnmt/go-api-boilerplate/internal/pkg/otel"
	"github.com/anhnmt/go-api-boilerplate/internal/pkg/postgres"
	"github.com/anhnmt/go-api-boilerplate/internal/pkg/redis"
)

type Config struct {
	App      config.App      `mapstructure:"app"`
	Logger   logger.Config   `mapstructure:"log"`
	Postgres postgres.Config `mapstructure:"postgres"`
	Redis    redis.Config    `mapstructure:"redis"`
	Server   config.Server   `mapstructure:"server"`
	JWT      config.JWT      `mapstructure:"jwt"`
	Crypto   config.Crypto   `mapstructure:"crypto"`
	Otel     otel.Config     `mapstructure:"otel"`
	Casbin   casbin.Config   `mapstructure:"casbin"`
}
