package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const DefaultFile = "limo.db"

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DefaultFile)
	if err != nil {
		return db, err
	}

	boil.SetDB(db)

	return db, err
}
