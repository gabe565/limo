package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const DSN = "file:data/limo.db?cache=shared"

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DSN)
	if err != nil {
		return db, err
	}

	boil.SetDB(db)
	boil.DebugMode = true

	return db, err
}
