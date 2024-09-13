package config

import (
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	AccesKey   string `mapstructure:"AccesKey"`
	SecretKey  string `mapstructure:"SecretKey"`
	BucketName string `mapstructure:"BucketName"`

	Host       string `mapstructure:"HOST"`
	ServerPort string `mapstructure:"SERVER_PORT"`
}

var envs = []string{
	"AccesKey", "SecretKey", "BucketName", "HOST", "SERVER_PORT",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigFile("./pkg/config/.env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
