package dtos

import (
	"time"

	"github.com/google/uuid"
)

type AnalystRatingDTO struct {
	ID                       uuid.UUID `json:"id"`
	Ticker                   string    `json:"ticker"`
	TargetFrom               float64   `json:"target_from"`
	TargetTo                 float64   `json:"target_to"`
	Company                  string    `json:"company"`
	Action                   string    `json:"action"`
	Brokerage                string    `json:"brokerage"`
	RatingFrom               string    `json:"rating_from"`
	RatingTo                 string    `json:"rating_to"`
	RatingIncreasePercentage float64   `json:"rating_increase_percentage"`
	CombinedPredictionIndex  float64   `json:"combined_prediction_index"`
	RatedAt                  time.Time `json:"rated_at"`
}

type AnalystRatingIndicatorsDTO struct {
	BuyNowPercentage                    float64 `json:"buy_now_percentage"`
	PositiveTargetAdjustmentPercentage  float64 `json:"positive_target_adjustment_percentage"`
	HighestIncrementInTargetPrice       float64 `json:"highest_increment_in_target_price"`
	HighestIncrementInTargetPriceTicker string  `json:"highest_increment_in_target_price_ticker"`
	HighestIncrementInTargetPriceName   string  `json:"highest_increment_in_target_price_name"`
}
