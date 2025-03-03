package models

import "time"

type AdapterLog struct {
	ID          uint   `gorm:"primaryKey"`
	AdapterName string `gorm:"uniqueIndex"`
	RunAt       time.Time
}
