package env

import (
	"fast-menu-api/logger"

	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		logger.Use().Error().
			Msg(err.Error())
	}
}

func Get(key string) string {
	return viper.GetString(key)
}
