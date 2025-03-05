package dtos

import "time"

type AnalystRatingDTO struct {
	ID         uint      `json:"id"`
	Ticker     string    `json:"ticker"`
	TargetFrom float64   `json:"target_from"`
	TargetTo   float64   `json:"target_to"`
	Company    string    `json:"company"`
	Action     string    `json:"action"`
	Brokerage  string    `json:"brokerage"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	RatedAt    time.Time `json:"rated_at"`
}
