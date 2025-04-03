package services

import (
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/repositories"
	utils "backend/internal/utils/cpi"
	"sort"
)

type AnalystRatingService interface {
	SaveAnalystRating(rating *models.AnalystRating) error
	SaveAnalystRatingsBatch(ratings []models.AnalystRating) error
	GetAll(sortOrder, sortBy, sourceID string, filters map[string]string, page, limit int) ([]dtos.AnalystRatingDTO, int64, error)
	GetIndicators(sourceID string) (dtos.AnalystRatingIndicatorsDTO, error)
	GetMinMaxCPI() (float64, float64, error)
	GetDashboardRatings(sourceID string) (dtos.AnalystRatingDashboardDTO, error)
}

type analystRatingService struct {
	ratingRepo *repositories.AnalystRatingRepository
}

func NewAnalystRatingService(ratingRepo *repositories.AnalystRatingRepository) AnalystRatingService {
	return &analystRatingService{
		ratingRepo: ratingRepo,
	}
}

func (service *analystRatingService) SaveAnalystRating(rating *models.AnalystRating) error {
	return service.ratingRepo.Create(rating)
}

func (service *analystRatingService) SaveAnalystRatingsBatch(ratings []models.AnalystRating) error {
	return service.ratingRepo.CreateBatch(ratings)
}

func (service *analystRatingService) GetAll(sortOrder, sortBy, sourceID string, filters map[string]string, page, limit int) ([]dtos.AnalystRatingDTO, int64, error) {
	ratings, total, err := service.ratingRepo.GetAll(sortOrder, sortBy, sourceID, filters, page, limit)
	if err != nil {
		return nil, 0, err
	}

	minCPI, maxCPI, err := service.GetMinMaxCPI()
	if err != nil {
		return nil, 0, err
	}

	if minCPI == maxCPI {
		maxCPI += 0.1
	}

	normalizeCPI := func(rawCPI float64) float64 {
		return ((rawCPI - minCPI) / (maxCPI - minCPI)) * 100
	}

	var ratingsDTOs []dtos.AnalystRatingDTO
	for _, rating := range ratings {
		ratingsDTOs = append(ratingsDTOs, dtos.AnalystRatingDTO{
			ID:                       rating.ID,
			Ticker:                   rating.Company.Ticker,
			TargetFrom:               rating.TargetFrom,
			TargetTo:                 rating.TargetTo,
			Company:                  rating.Company.Name,
			CompanyID:                rating.Company.ID.String(),
			Action:                   rating.Action,
			Brokerage:                rating.Brokerage,
			RatingFrom:               rating.RatingFrom,
			RatingTo:                 rating.RatingTo,
			RatedAt:                  rating.RatedAt,
			RatingIncreasePercentage: (rating.TargetTo - rating.TargetFrom) / rating.TargetFrom * 100,
			CombinedPredictionIndex:  normalizeCPI(rating.CombinedPredictionIndex),
		})
	}

	return ratingsDTOs, total, nil
}

func (service *analystRatingService) GetMinMaxCPI() (float64, float64, error) {
	return service.ratingRepo.GetMinMaxCPI()
}

func (service *analystRatingService) GetIndicators(sourceID string) (dtos.AnalystRatingIndicatorsDTO, error) {
	totalCount, err := service.ratingRepo.GetTotalRatingsCount(sourceID)
	if err != nil || totalCount == 0 {
		return dtos.AnalystRatingIndicatorsDTO{}, err
	}

	minCPI, maxCPI, err := service.ratingRepo.GetMinMaxCPI()
	if err != nil {
		return dtos.AnalystRatingIndicatorsDTO{}, err
	}

	return service.ratingRepo.CalculateIndicators(sourceID, totalCount, minCPI, maxCPI)
}

func (service *analystRatingService) GetDashboardRatings(sourceID string) (dtos.AnalystRatingDashboardDTO, error) {
	latestRatings, err := service.ratingRepo.GetLatestRatings(sourceID)
	if err != nil {
		return dtos.AnalystRatingDashboardDTO{}, err
	}

	latestDtoRatings, err := service.formatToLatestRatingsDTO(latestRatings)
	if err != nil {
		return dtos.AnalystRatingDashboardDTO{}, err
	}

	allRatings, err := service.ratingRepo.GetRawCPIData(sourceID)
	if err != nil {
		return dtos.AnalystRatingDashboardDTO{}, err
	}

	minCPI, maxCPI, err := service.GetMinMaxCPI()
	if err != nil {
		return dtos.AnalystRatingDashboardDTO{}, err
	}
	if minCPI == maxCPI {
		maxCPI += 0.1
	}

	categoryCounts := make(map[string]int)
	for _, rating := range allRatings {
		normalizedCPI := ((rating.CombinedPredictionIndex - minCPI) / (maxCPI - minCPI)) * 100
		category := utils.CategorizeCPI(normalizedCPI)
		categoryCounts[category]++
	}

	var donutCPIChart []dtos.DonutChartDTO
	for category, count := range categoryCounts {
		donutCPIChart = append(donutCPIChart, dtos.DonutChartDTO{
			Label: category,
			Count: count,
		})
	}

	sort.Slice(donutCPIChart, func(i, j int) bool {
		return donutCPIChart[i].Count > donutCPIChart[j].Count
	})

	donutRatingChart, err := service.ratingRepo.GetDonutRatingChart(sourceID)
	if err != nil {
		return dtos.AnalystRatingDashboardDTO{}, err
	}

	return dtos.AnalystRatingDashboardDTO{
		LatestRatings:    latestDtoRatings,
		DonutCPIChart:    donutCPIChart,
		DonutRatingChart: donutRatingChart,
	}, nil
}

func (service *analystRatingService) formatToLatestRatingsDTO(ratings []models.AnalystRating) ([]dtos.AnalystRatingDTO, error) {
	var latestRatingsDTO []dtos.AnalystRatingDTO
	for _, rating := range ratings {
		latestRatingsDTO = append(latestRatingsDTO, dtos.AnalystRatingDTO{
			ID:                       rating.ID,
			Company:                  rating.Company.Name,
			Ticker:                   rating.Company.Ticker,
			TargetFrom:               rating.TargetFrom,
			TargetTo:                 rating.TargetTo,
			Action:                   rating.Action,
			Brokerage:                rating.Brokerage,
			RatingFrom:               rating.RatingFrom,
			RatingTo:                 rating.RatingTo,
			RatingIncreasePercentage: (rating.TargetTo - rating.TargetFrom) / rating.TargetFrom * 100,
			CombinedPredictionIndex:  rating.CombinedPredictionIndex,
			RatedAt:                  rating.RatedAt,
		})
	}
	return latestRatingsDTO, nil
}
