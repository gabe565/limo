package database

import (
	"github.com/onrik/gorm-logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(config Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.BuildDsn()), &gorm.Config{
		Logger: gorm_logrus.New(),
	})
	if err != nil {
		return db, err
	}

	return db, err
}
