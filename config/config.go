package config

import (
	"github.com/spf13/viper"
)

var _config *viper.Viper

func GetConfig() *viper.Viper {
	if _config == nil {
		_config = loadConfig()
	}
	return _config
}

func loadConfig() *viper.Viper {
	config := viper.New()

	config.AddConfigPath("./config/")
	config.SetConfigName("conf")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	return config
}
