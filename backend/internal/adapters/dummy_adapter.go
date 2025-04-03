package adapters

import (
	"backend/internal/models"
	"backend/internal/services"
	cpi "backend/internal/utils/cpi"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type DummyAdapter struct {
	ratingService  services.AnalystRatingService
	companyService services.CompanyService
	dataSourceID   uuid.UUID
}

func NewDummyAdapter(ratingService services.AnalystRatingService, companyService services.CompanyService, dataSourceID uuid.UUID) RatingAdapter {
	return &DummyAdapter{
		ratingService:  ratingService,
		companyService: companyService,
		dataSourceID:   dataSourceID,
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
		existingCompany, err := dummyAdapter.companyService.CreateCompanyByTicker(company.Ticker, company.Company)
		if err != nil {
			return nil, err
		}

		targetFrom := rand.Float64()*100 + 50
		targetTo := targetFrom + rand.Float64()*20
		ratingFrom := ratingOptions[rand.Intn(len(ratingOptions))]
		ratingTo := ratingOptions[rand.Intn(len(ratingOptions))]

		rating := models.AnalystRating{
			TargetFrom:                 targetFrom,
			TargetTo:                   targetTo,
			Action:                     "upgraded by",
			Brokerage:                  "Dummy Brokerage",
			RatingFrom:                 ratingFrom,
			RatingTo:                   ratingTo,
			RatedAt:                    time.Now(),
			ActionImpactScore:          cpi.CalculateActionImpactScore("upgraded by"),
			RatingChangeImpact:         cpi.CalculateRatingChangeImpact(ratingFrom, ratingTo),
			TargetAdjustmentPercentage: cpi.CalculateTargetAdjustment(targetFrom, targetTo),
			DataSourceID:               dummyAdapter.dataSourceID,
			CompanyID:                  existingCompany.ID,
		}
		rating.CombinedPredictionIndex = cpi.CalculateRawCPI(rating.ActionImpactScore, rating.RatingChangeImpact, rating.TargetAdjustmentPercentage)
		ratings = append(ratings, rating)
	}

	err := dummyAdapter.ratingService.SaveAnalystRatingsBatch(ratings)
	if err != nil {
		return nil, err
	}

	return ratings, nil
}
