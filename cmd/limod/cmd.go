package main

import (
	_ "embed"
	"github.com/gabe565/limo/internal/completion"
	"github.com/gabe565/limo/internal/config"
	"github.com/gabe565/limo/internal/database"
	"github.com/gabe565/limo/internal/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

//go:embed description.txt
var description string

var Command = &cobra.Command{
	Use:   "limod [address]",
	Short: "Limo server",
	Long:  description,
	Args:  cobra.MaximumNArgs(1),
	RunE:  run,
}

var completionFlag string

func init() {
	cobra.OnInitialize(config.InitViper("limod"))

	completion.CompletionFlag(Command, &completionFlag)
}

func run(cmd *cobra.Command, args []string) error {
	if completionFlag != "" {
		return completion.Run(cmd, completionFlag)
	}
	cmd.SilenceUsage = true

	var err error
	if err = os.MkdirAll("data/files", 0755); err != nil {
		log.Panic(err)
	}

	db, err := database.Open(database.NewConfig())
	if err != nil {
		log.Panic(err)
	}

	if err = database.Migrate(db); err != nil {
		log.Panic(err)
	}

	s := server.Server{
		DB: db,
	}
	go func() {
		if err := s.ListenAndServe(viper.GetString("address")); err != nil {
			log.Panic(err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs
	log.Info("exiting")
	return nil
}

func main() {
	if err := Command.Execute(); err != nil {
		os.Exit(1)
	}
}
