package models

import (
	"time"

	"gorm.io/gorm"
)

type AnalystRating struct {
	gorm.Model
	Ticker     string    `gorm:"not null"`
	TargetFrom float64   `gorm:"not null"`
	TargetTo   float64   `gorm:"not null"`
	Company    string    `gorm:"not null"`
	Action     string    `gorm:"not null"`
	Brokerage  string    `gorm:"not null"`
	RatingFrom string    `gorm:"not null"`
	RatingTo   string    `gorm:"not null"`
	RatedAt    time.Time `gorm:"not null"`

	DataSourceID uint       `gorm:"not null"`
	DataSource   DataSource `gorm:"foreignKey:DataSourceID"`
}
