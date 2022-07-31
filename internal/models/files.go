package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name      string
	ExpiresAt sql.NullTime
}
