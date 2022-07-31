package main

import "github.com/spf13/viper"

func init() {
	Command.Flags().DurationP("expires", "e", 0, "Set expiration time")
	if err := viper.BindPFlag("expires", Command.Flags().Lookup("expires")); err != nil {
		panic(err)
	}
}
