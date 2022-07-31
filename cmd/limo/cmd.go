package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gabe565/limo/internal/completion"
	"github.com/gabe565/limo/internal/config"
	"github.com/gabe565/limo/internal/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path/filepath"
)

//go:embed description.txt
var description string

var Command = &cobra.Command{
	Use:   "limo file",
	Short: "Upload files with style",
	Long:  description,
	RunE:  run,
}

var completionFlag string

func init() {
	cobra.OnInitialize(config.InitViper("limo"))

	completion.CompletionFlag(Command, &completionFlag)
}

func run(cmd *cobra.Command, args []string) error {
	if completionFlag != "" {
		return completion.Run(cmd, completionFlag)
	}

	if len(args) != 1 {
		return errors.New("file is required")
	}

	cmd.SilenceUsage = true

	filename := args[0]

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	var u URLFlag
	if err := u.Set(viper.GetString("address")); err != nil {
		return err
	}
	u.Path = "/api/files/" + filepath.Base(filename)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", u.String(), f)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")

	if viper.GetBool("random") {
		req.Header.Set("Random", "1")
	}
	if viper.IsSet("expires-in") {
		req.Header.Set("ExpiresIn", viper.GetDuration("expires-in").String())
	}

	raw, err := client.Do(req)
	if err != nil {
		return err
	}

	var resp server.PutFileResponse
	if err = json.NewDecoder(raw.Body).Decode(&resp); err != nil {
		return err
	}

	var output Output
	if err := output.Set(viper.GetString("output")); err != nil {
		return err
	}
	switch output {
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
