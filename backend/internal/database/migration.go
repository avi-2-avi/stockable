package database

import (
	"backend/internal/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Company{},
		&models.DataSource{},
		&models.Portafolio{},
		&models.PortafolioHolding{},
		&models.AnalystRating{},
		&models.AdapterLog{},
	)
	if err != nil {
		return err
	}

	AddAdminRole(db)
	AddUserRole(db)
	return AddDefaultUser(db)
}

func AddAdminRole(db *gorm.DB) error {
	var count int64
	db.Model(&models.Role{}).Where("name = ?", "admin").Count(&count)

	if count == 0 {
		adminRole := models.Role{
			ID:   1,
			Name: "admin",
		}
		if err := db.Create(&adminRole).Error; err != nil {
			return fmt.Errorf("failed to create admin role: %v", err)
		} else {
			fmt.Println("Admin role created successfully")
		}
	}
	return nil
}

func AddUserRole(db *gorm.DB) error {
	var count int64
	db.Model(&models.Role{}).Where("name = ?", "user").Count(&count)

	if count == 0 {
		userRole := models.Role{
			ID:   2,
			Name: "user",
		}
		if err := db.Create(&userRole).Error; err != nil {
			return fmt.Errorf("failed to create user role: %v", err)
		} else {
			fmt.Println("User role created successfully")
		}
	}
	return nil
}

func AddDefaultUser(db *gorm.DB) error {
	var count int64
	db.Model(&models.User{}).Where("email = ?", "admin@mail.com").Count(&count)

	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Admin123!"), bcrypt.DefaultCost)
		admin := models.User{
			FullName: "Admin",
			Email:    "admin@mail.com",
			Password: string(hashedPassword),
			RoleID:   1,
		}
		if err := db.Create(&admin).Error; err != nil {
			return fmt.Errorf("failed to create admin user: %v", err)
		} else {
			fmt.Println("Admin user created successfully")
		}
	}
	return nil
}
