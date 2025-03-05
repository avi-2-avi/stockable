package models

import (
	"time"

	"gorm.io/gorm"
)

type DataSource struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"unique"`
	CreateAt  time.Time      `gorm:"autoCreateTime"`
	UpdateAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
