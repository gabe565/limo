package database

import (
	"github.com/gabe565/limo/internal/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	log.Info("running migrations")

	if err := db.AutoMigrate(&models.File{}); err != nil {
		return err
	}

	return nil
}
