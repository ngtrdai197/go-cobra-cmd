package config

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	AppEnv    string `mapstructure:"APP_ENV" validate:"required"`
	RedisHost string `mapstructure:"REDIS_HOST" validate:"required"`
	RedisPort string `mapstructure:"REDIS_PORT" validate:"required"`
	RedisDb   int    `validate:"required"`
}

func GetConfig(validator *validator.Validate) (*Config, error) {
	c := &Config{
		AppEnv:    viper.GetString("APP_ENV"),
		RedisHost: viper.GetString("REDIS_HOST"),
		RedisPort: viper.GetString("REDIS_PORT"),
		RedisDb:   viper.GetInt("REDIS_DB"),
	}
	if err := validator.StructCtx(context.Background(), c); err != nil {
		return nil, err
	}
	return c, nil
}
