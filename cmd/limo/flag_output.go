package main

import "github.com/spf13/viper"

func init() {
	v := OutputText
	Command.Flags().VarP(&v, "output", "o", "Output format (text|t|json|j)")
	if err := viper.BindPFlag("output", Command.Flags().Lookup("output")); err != nil {
		panic(err)
	}
}
