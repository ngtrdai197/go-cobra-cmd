package config

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	AppEnv   string `mapstructure:"APP_ENV" validate:"required"`
	RedisUrl string `mapstructure:"REDIS_URL" validate:"required"`
}

func LoadConfig(validator *validator.Validate) (*Config, error) {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w", err))
	}
	c := &Config{
		AppEnv:   viper.GetString("APP_ENV"),
		RedisUrl: viper.GetString("REDIS_URL"),
	}
	if err := validator.StructCtx(context.Background(), c); err != nil {
		return nil, err
	}
	return c, nil
}
