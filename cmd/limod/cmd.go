package main

import (
	"database/sql"
	_ "embed"
	"github.com/gabe565/limo/internal/completion"
	"github.com/gabe565/limo/internal/database"
	"github.com/gabe565/limo/internal/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

//go:embed description.txt
var description string

var Command = &cobra.Command{
	Use:               "limod [address]",
	Short:             "Limo server",
	Long:              description,
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: validArgs,
	PreRunE:           preRun,
	RunE:              run,
}

var completionFlag string

func init() {
	completion.CompletionFlag(Command, &completionFlag)
}

var address = "127.0.0.1:3000"

func validArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var validArgs []string
	if len(args) == 0 {
		validArgs = []string{":80", ":8080", ":3000"}
		for _, port := range validArgs {
			validArgs = append(validArgs, "127.0.0.1"+port)
		}
	}
	return validArgs, cobra.ShellCompDirectiveNoFileComp
}

func preRun(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		address = args[0]
	}

	return nil
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

	db, err := database.Open()
	if err != nil {
		log.Panic(err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	if err = database.Migrate(db); err != nil {
		log.Panic(err)
	}

	s := server.Server{}
	go func() {
		if err := s.ListenAndServe(address); err != nil {
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
