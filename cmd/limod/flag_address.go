package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	Command.Flags().String("addr", "127.0.0.1:3000", "Listen address")
	if err := Command.RegisterFlagCompletionFunc(
		"addr",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			ports := []uint16{80, 8080, 3000}

			var validArgs []string
			for _, port := range ports {
				validArgs = append(
					validArgs,
					fmt.Sprintf(":%d\tPublic", port),
					fmt.Sprintf("127.0.0.1:%d\tPrivate", port),
				)
			}
			return validArgs, cobra.ShellCompDirectiveNoFileComp
		},
	); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("address", Command.Flags().Lookup("addr")); err != nil {
		panic(err)
	}
}
