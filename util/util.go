package util

import (
	"github.com/spf13/viper"
	"log"
)

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error when reading config : %s", err)
	}
	return viper.GetString(key)
}
