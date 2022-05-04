package main

import (
	_ "embed"
	"github.com/spf13/cobra"
	"os"
)

//go:embed description.txt
var description string

var Command = &cobra.Command{
	Use:   "limo",
	Short: "Upload files with style",
	Long:  description,
}

func main() {
	if err := Command.Execute(); err != nil {
		os.Exit(1)
	}
}
