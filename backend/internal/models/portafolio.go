package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Portafolio struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Category  string    `json:"category" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID uuid.UUID `gorm:"not null"`
	User   User      `gorm:"foreignKey:UserID"`

	DataSourceID uuid.UUID  `gorm:"not null"`
	DataSource   DataSource `gorm:"foreignKey:DataSourceID"`
}
