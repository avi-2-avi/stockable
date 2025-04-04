package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnalystRating struct {
	ID                         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TargetFrom                 float64   `gorm:"not null"`
	TargetTo                   float64   `gorm:"not null"`
	Action                     string    `gorm:"not null"`
	Brokerage                  string    `gorm:"not null"`
	RatingFrom                 string    `gorm:"not null"`
	RatingTo                   string    `gorm:"not null"`
	ActionImpactScore          float64   `gorm:"null"`
	RatingChangeImpact         float64   `gorm:"null"`
	TargetAdjustmentPercentage float64   `gorm:"null"`
	CombinedPredictionIndex    float64   `gorm:"null"`
	RatedAt                    time.Time `gorm:"not null"`
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
	DeletedAt                  gorm.DeletedAt `gorm:"index"`

	DataSourceID uuid.UUID  `gorm:"not null"`
	DataSource   DataSource `gorm:"foreignKey:DataSourceID"`

	CompanyID uuid.UUID `gorm:"not null"`
	Company   Company   `gorm:"foreignKey:CompanyID"`
}
