package models

import "time"

type AdapterLog struct {
	ID          uint   `gorm:"primaryKey"`
	AdapterName string `gorm:"not null"`
	RunAt       time.Time
}
