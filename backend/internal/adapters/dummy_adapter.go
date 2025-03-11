package adapters

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type DummyAdapter struct {
	service      services.AnalystRatingsService
	dataSourceID uuid.UUID
}

func NewDummyAdapter(service services.AnalystRatingsService, dataSourceID uuid.UUID) RatingAdapter {
	return &DummyAdapter{
		service:      service,
		dataSourceID: dataSourceID,
	}
}

func (dummyAdapter *DummyAdapter) FetchData() ([]models.AnalystRating, error) {
	companies := []struct {
		Ticker  string
		Company string
	}{
		{"AAPL", "Apple Inc."},
		{"MSFT", "Microsoft Corp."},
		{"GOOGL", "Alphabet Inc."},
		{"AMZN", "Amazon.com Inc."},
		{"TSLA", "Tesla Inc."},
		{"NVDA", "NVIDIA Corp."},
		{"META", "Meta Platforms Inc."},
		{"BRK.B", "Berkshire Hathaway Inc."},
		{"V", "Visa Inc."},
		{"JNJ", "Johnson & Johnson"},
		{"WMT", "Walmart Inc."},
		{"PG", "Procter & Gamble Co."},
		{"JPM", "JPMorgan Chase & Co."},
		{"UNH", "UnitedHealth Group Inc."},
		{"MA", "Mastercard Inc."},
		{"HD", "The Home Depot Inc."},
		{"XOM", "Exxon Mobil Corp."},
		{"BAC", "Bank of America Corp."},
		{"PFE", "Pfizer Inc."},
		{"KO", "Coca-Cola Co."},
	}

	ratingOptions := []string{"Sell", "Hold", "Buy", "Neutral"}

	var ratings []models.AnalystRating
	for _, company := range companies {
		targetFrom := rand.Float64()*100 + 50
		targetTo := targetFrom + rand.Float64()*20
		ratingFrom := ratingOptions[rand.Intn(len(ratingOptions))]
		ratingTo := ratingOptions[rand.Intn(len(ratingOptions))]

		rating := models.AnalystRating{
			Ticker:                     company.Ticker,
			Company:                    company.Company,
			TargetFrom:                 targetFrom,
			TargetTo:                   targetTo,
			Action:                     "upgraded by",
			Brokerage:                  "Dummy Brokerage",
			RatingFrom:                 ratingFrom,
			RatingTo:                   ratingTo,
			RatedAt:                    time.Now(),
			DataSourceID:               dummyAdapter.dataSourceID,
			ActionImpactScore:          utils.CalculateActionImpactScore("upgraded by"),
			RatingChangeImpact:         utils.CalculateRatingChangeImpact(ratingFrom, ratingTo),
			TargetAdjustmentPercentage: utils.CalculateTargetAdjustment(targetFrom, targetTo),
		}
		rating.CombinedPredictionIndex = utils.CalculateRawCPI(rating.ActionImpactScore, rating.RatingChangeImpact, rating.TargetAdjustmentPercentage)
		ratings = append(ratings, rating)
	}

	err := dummyAdapter.service.SaveAnalystRatingsBatch(ratings)
	if err != nil {
		return nil, err
	}

	return ratings, nil
}
