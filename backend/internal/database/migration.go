package database

import (
	"data-loader/internal/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.DataSource{}, &models.AnalystRating{})
	if err != nil {
		return err
	}

	return nil
}
