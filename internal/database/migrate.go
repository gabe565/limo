package database

import (
	"database/sql"
	"github.com/gabe565/limo/internal/migrations"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
)

func Migrate(db *sql.DB) error {
	goose.SetBaseFS(migrations.Embed)
	goose.SetLogger(log.WithField("package", "database"))

	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}

	if err := goose.Up(db, "."); err != nil {
		return err
	}

	return nil
}
