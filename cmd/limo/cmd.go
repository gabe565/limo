package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/gabe565/limo/internal/server"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"path/filepath"
)

//go:embed description.txt
var description string

var Command = &cobra.Command{
	Use:   "limo",
	Short: "Upload files with style",
	Long:  description,
	Args:  cobra.ExactArgs(1),
	RunE:  run,
}

var conf Config

func init() {
	Command.Flags().Var((*URLFlag)(&conf.Address), "addr", "Server address. If not given, scheme will default to https.")
	Command.Flags().BoolVarP(&conf.Random, "random", "r", false, "Random filename")
	Command.Flags().VarP(&conf.Output, "output", "o", "Output format (text|t|json|j)")
}

func run(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true

	filename := args[0]

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	u := conf.Address
	u.Path = "/api/files/" + filepath.Base(filename)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", u.String(), f)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")

	if conf.Random {
		req.Header.Set("Random", "1")
	}

	raw, err := client.Do(req)
	if err != nil {
		return err
	}

	var resp server.PutFileResponse
	if err = json.NewDecoder(raw.Body).Decode(&resp); err != nil {
		return err
	}

	switch conf.Output {
	case OutputText:
		fmt.Println(resp.URL)
	case OutputJson:
		b, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	}

	return nil
}

func main() {
	if err := Command.Execute(); err != nil {
		os.Exit(1)
	}
}
