package config

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	AppEnv                   string `mapstructure:"APP_ENV" validate:"required"`
	PublicApiAddress         string `mapstructure:"PUBLIC_API_ADDRESS" validate:"required"`
	RedisHost                string `mapstructure:"REDIS_HOST" validate:"required"`
	RedisPort                string `mapstructure:"REDIS_PORT" validate:"required"`
	DbDriver                 string `mapstructure:"DB_DRIVER" validate:"required"`
	DbSource                 string `mapstructure:"DB_SOURCE" validate:"required"`
	RedisDb                  int    `validate:"required"`
	KafkaBrokers             string `mapstructure:"KAFKA_BROKERS" validate:"required"`
	KafkaCreateUserSendEmail string `mapstructure:"KAFKA_CREATE_USER_SEND_EMAIL" validate:"required"`
}

func GetConfig(validator *validator.Validate) (*Config, error) {
	c := &Config{
		AppEnv:                   viper.GetString("APP_ENV"),
		PublicApiAddress:         viper.GetString("PUBLIC_API_ADDRESS"),
		RedisHost:                viper.GetString("REDIS_HOST"),
		RedisPort:                viper.GetString("REDIS_PORT"),
		RedisDb:                  viper.GetInt("REDIS_DB"),
		DbDriver:                 viper.GetString("DB_DRIVER"),
		DbSource:                 viper.GetString("DB_SOURCE"),
		KafkaBrokers:             viper.GetString("KAFKA_BROKERS"),
		KafkaCreateUserSendEmail: viper.GetString("KAFKA_CREATE_USER_SEND_EMAIL"),
	}
	if err := validator.StructCtx(context.Background(), c); err != nil {
		return nil, err
	}
	return c, nil
}
