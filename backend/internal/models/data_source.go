package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DataSource struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string         `gorm:"unique"`
	IsVisible bool           `gorm:"default:false"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
