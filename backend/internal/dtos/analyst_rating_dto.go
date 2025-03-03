package dtos

type AnalystRatingDTO struct {
	ID         uint    `json:"id"`
	Ticker     string  `gorm:"not null"`
	TargetFrom float64 `gorm:"not null"`
	TargetTo   float64 `gorm:"not null"`
	Company    string  `gorm:"not null"`
	Action     string  `gorm:"not null"`
	Brokerage  string  `gorm:"not null"`
	RatingFrom string  `gorm:"not null"`
	RatingTo   string  `gorm:"not null"`
}
