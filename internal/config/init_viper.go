package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func InitViper(name string) func() {
	return func() {
		viper.SetConfigName(name)
		viper.SetConfigType("yaml")
		viper.AddConfigPath("$HOME/.config/")
		viper.AddConfigPath("$HOME/")
		viper.AddConfigPath("/etc/" + name + "/")

		viper.AutomaticEnv()
		viper.SetEnvPrefix(name)
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error
			} else {
				// Config file was found but another error was produced
				panic(fmt.Errorf("Fatal error reading config file: %w \n", err))
			}
		}
	}
}
