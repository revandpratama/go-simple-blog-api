package config

import "github.com/spf13/viper"

type Config struct {
}

var ENV *Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		panic(err)
	}
}
