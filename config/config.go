package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	config := viper.New()

	config.AddConfigPath("./config/")
	config.SetConfigName("conf")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	return config
}
