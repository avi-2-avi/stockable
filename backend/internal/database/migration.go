package database

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.DataSource{}, &models.AnalystRating{}, &models.AdapterLog{})
	if err != nil {
		return err
	}

	return nil
}
