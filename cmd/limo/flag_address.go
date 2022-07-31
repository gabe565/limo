package main

import "github.com/spf13/viper"

func init() {
	Command.Flags().Var(&URLFlag{}, "addr", "Server address. If not given, scheme will default to https.")
	if err := viper.BindPFlag("address", Command.Flags().Lookup("addr")); err != nil {
		panic(err)
	}
}
