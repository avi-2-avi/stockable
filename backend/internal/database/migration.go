package database

import (
	"backend/internal/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.DataSource{}, &models.AnalystRating{}, &models.AdapterLog{})
	if err != nil {
		return err
	}

	return AddDefaultUser(db)
}

func AddDefaultUser(db *gorm.DB) error {
	var count int64
	db.Model(&models.User{}).Where("email = ?", "admin@mail.com").Count(&count)

	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123!"), bcrypt.DefaultCost)
		admin := models.User{
			FullName: "Admin",
			Email:    "admin@mail.com",
			Password: string(hashedPassword),
		}
		if err := db.Create(&admin).Error; err != nil {
			return fmt.Errorf("failed to create admin user: %v", err)
		} else {
			fmt.Println("Admin user created successfully")
		}
	}
	return nil
}
