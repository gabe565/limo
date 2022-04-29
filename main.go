package main

import (
	"github.com/gabe565/limo/internal/database"
	"github.com/gabe565/limo/internal/server"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//go:generate sqlboiler sqlite3

func main() {
	db, err := database.Open()
	if err != nil {
		log.Panic(err)
	}

	if err = database.Migrate(db); err != nil {
		log.Panic(err)
	}

	r := server.New()
	address := "127.0.0.1:3000"
	go func() {
		if err := http.ListenAndServe(address, r); err != nil {
			log.Panic(err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs
	log.Info("exiting")
}
