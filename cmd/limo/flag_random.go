package main

import "github.com/spf13/viper"

func init() {
	Command.Flags().BoolP("random", "r", false, "Random filename")
	if err := viper.BindPFlag("random", Command.Flags().Lookup("random")); err != nil {
		panic(err)
	}
}
