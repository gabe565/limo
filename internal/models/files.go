package models

import (
	"database/sql"
	"gorm.io/gorm"
	"path"
)

type File struct {
	gorm.Model
	Name      string
	ExpiresAt sql.NullTime
}

func (f File) PublicPath() string {
	return path.Join("/f", f.Name)
}

func (f File) RawPath() string {
	return path.Join("/raw", f.Name)
}
