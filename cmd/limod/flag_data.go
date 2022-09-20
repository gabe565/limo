package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	Command.Flags().StringP("data-dir", "d", "./data", "Data directory")
	if err := Command.RegisterFlagCompletionFunc("data-dir", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveFilterDirs
	}); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("data", Command.Flags().Lookup("data-dir")); err != nil {
		panic(err)
	}
}
