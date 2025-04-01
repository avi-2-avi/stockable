package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PortafolioHolding struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Quantity      float64   `json:"quantity" gorm:"not null"`
	PurchasePrice float64   `json:"purchase_price" gorm:"not null"`
	PurchasedAt   time.Time `json:"purchased_at" gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`

	PortafolioID uuid.UUID  `gorm:"not null"`
	Portafolio   Portafolio `gorm:"foreignKey:PortafolioID"`

	CompanyID uuid.UUID `gorm:"not null"`
	Company   Company   `gorm:"foreignKey:CompanyID"`
}
