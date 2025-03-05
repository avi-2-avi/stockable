package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DataSource struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string         `gorm:"unique"`
	CreateAt  time.Time      `gorm:"autoCreateTime"`
	UpdateAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (r *DataSource) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}
