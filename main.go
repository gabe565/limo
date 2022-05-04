package main

import (
	"database/sql"
	"github.com/gabe565/limo/internal/database"
	"github.com/gabe565/limo/internal/server"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

// Reset data dir
//go:generate rm -rf data
//go:generate mkdir data
// Run Goose migrations
//go:generate goose -dir=internal/migrations sqlite3 data/limo.db up
// Create SQLBoiler models
//go:generate sqlboiler sqlite3

func main() {
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
		if err := s.ListenAndServe("127.0.0.1:3000"); err != nil {
			log.Panic(err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs
	log.Info("exiting")
}
