package repositories

import (
	"backend/internal/dtos"
	"backend/internal/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnalystRatingRepository struct {
	db *gorm.DB
}

func NewAnalystRatingRepository(db *gorm.DB) *AnalystRatingRepository {
	return &AnalystRatingRepository{db: db}
}

func (r *AnalystRatingRepository) Create(rating *models.AnalystRating) error {
	return r.db.Create(rating).Error
}

func (r *AnalystRatingRepository) CreateBatch(ratings []models.AnalystRating) error {
	if len(ratings) == 0 {
		return nil
	}
	return r.db.Create(&ratings).Error
}

func (r *AnalystRatingRepository) GetByID(id uuid.UUID) (*models.AnalystRating, error) {
	var rating models.AnalystRating
	err := r.db.First(&rating, id).Error
	if err != nil {
		return nil, err
	}
	return &rating, nil
}

func (r *AnalystRatingRepository) GetAll(sortOrder, sortBy, sourceID string, filters map[string]string, page, limit int) ([]models.AnalystRating, int64, error) {
	var ratings []models.AnalystRating
	var total int64
	query := r.db

	query = applyFilters(query, sourceID, filters)
	query.Model(&models.AnalystRating{}).Count(&total)

	sortBy = cleanSortColumn(sortBy)
	query = applySorting(query, sortBy, sortOrder)
	query = applyPagination(query, page, limit)

	err := query.Preload("Company").Find(&ratings).Error
	if err != nil {
		return nil, 0, err
	}

	return ratings, total, nil
}

func cleanSortColumn(sortBy string) string {
	allowedColumns := allowedActionColumns()
	if _, exists := allowedColumns[sortBy]; !exists {
		return "combined_prediction_index"
	}
	return sortBy
}

func allowedActionColumns() map[string]bool {
	return map[string]bool{
		"target_from":                  true,
		"target_to":                    true,
		"action":                       true,
		"brokerage":                    true,
		"rating_from":                  true,
		"rating_to":                    true,
		"created_at":                   true,
		"combined_prediction_index":    true,
		"rating_increase_percentage":   true,
		"action_impact_score":          true,
		"rating_change_impact":         true,
		"target_adjustment_percentage": true,
	}
}

func applyFilters(query *gorm.DB, sourceID string, filters map[string]string) *gorm.DB {
	if sourceID != "" {
		query = query.Where("data_source_id = ?", sourceID)
	}

	allowedColumns := allowedActionColumns()
	for column, value := range filters {
		if _, exists := allowedColumns[column]; exists {
			query = query.Where(fmt.Sprintf("%s LIKE ?", column), "%"+value+"%")
		}
	}

	query = query.Where("rated_at = (SELECT MAX(rated_at) FROM analyst_ratings AS ar WHERE ar.company_id = analyst_ratings.company_id)")

	return query
}

func applySorting(query *gorm.DB, sortBy, sortOrder string) *gorm.DB {
	return query.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))
}

func applyPagination(query *gorm.DB, page, limit int) *gorm.DB {
	offset := (page - 1) * limit
	return query.Limit(limit).Offset(offset)
}

func (r *AnalystRatingRepository) GetTotalRatingsCount(sourceID string) (int64, error) {
	var totalCount int64
	query := r.db.Model(&models.AnalystRating{})
	if sourceID != "" {
		query = query.Where("data_source_id = ?", sourceID)
	}
	err := query.Count(&totalCount).Error
	return totalCount, err
}

func (r *AnalystRatingRepository) CalculateIndicators(sourceID string, totalCount int64, minCPI, maxCPI float64) (dtos.AnalystRatingIndicatorsDTO, error) {
	var dto dtos.AnalystRatingIndicatorsDTO

	var err error
	dto.BuyNowPercentage, err = r.getBuyNowPercentage(sourceID, totalCount, minCPI, maxCPI)
	if err != nil {
		return dto, err
	}

	dto.PositiveTargetAdjustmentPercentage, err = r.getPositiveTargetAdjustmentPercentage(sourceID, totalCount)
	if err != nil {
		return dto, err
	}

	highestIncrementRating, err := r.getHighestIncrementInTargetPrice(sourceID)
	if err != nil {
		return dto, err
	}

	dto.HighestIncrementInTargetPrice = highestIncrementRating.TargetTo - highestIncrementRating.TargetFrom
	dto.HighestIncrementInTargetPriceTicker = highestIncrementRating.Company.Ticker
	dto.HighestIncrementInTargetPriceName = highestIncrementRating.Company.Name

	return dto, nil
}

func (r *AnalystRatingRepository) getBuyNowPercentage(sourceID string, totalCount int64, minCPI, maxCPI float64) (float64, error) {
	if minCPI == maxCPI {
		return 0, nil
	}

	var buyNowCount int64
	query := r.db.Model(&models.AnalystRating{}).
		Where("(CAST(combined_prediction_index AS NUMERIC) - CAST(? AS NUMERIC)) / (CAST(? AS NUMERIC) - CAST(? AS NUMERIC)) * 100 > ?",
			minCPI, maxCPI, minCPI, 70)

	if sourceID != "" {
		query = query.Where("data_source_id = ?", sourceID)
	}

	err := query.Count(&buyNowCount).Error
	if err != nil {
		return 0, err
	}

	return float64(buyNowCount) / float64(totalCount) * 100, nil
}

func (r *AnalystRatingRepository) getPositiveTargetAdjustmentPercentage(sourceID string, totalCount int64) (float64, error) {
	var positiveCount int64
	query := r.db.Model(&models.AnalystRating{}).Where("target_adjustment_percentage > 0")
	if sourceID != "" {
		query = query.Where("data_source_id = ?", sourceID)
	}
	err := query.Count(&positiveCount).Error
	if err != nil {
		return 0, err
	}
	return float64(positiveCount) / float64(totalCount) * 100, nil
}

func (r *AnalystRatingRepository) getHighestIncrementInTargetPrice(sourceID string) (models.AnalystRating, error) {
	var rating models.AnalystRating

	subQuery := r.db.Model(&models.AnalystRating{}).
		Select("id").
		Order("(target_to - target_from) DESC").
		Limit(1)

	if sourceID != "" {
		subQuery = subQuery.Where("data_source_id = ?", sourceID)
	}

	err := r.db.Preload("Company").Where("id = (?)", subQuery).First(&rating).Error
	if err != nil {
		return models.AnalystRating{}, err
	}

	return rating, nil
}

func (r *AnalystRatingRepository) GetMinMaxCPI() (float64, float64, error) {
	var result struct {
		Min float64 `gorm:"column:min"`
		Max float64 `gorm:"column:max"`
	}

	err := r.db.Raw("SELECT MIN(combined_prediction_index) as min, MAX(combined_prediction_index) as max FROM analyst_ratings").
		Scan(&result).Error

	if err != nil {
		return 0, 0, err
	}

	return result.Min, result.Max, nil
}

func (r *AnalystRatingRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.AnalystRating{}, id).Error
}

func (r *AnalystRatingRepository) GetLatestRatings(sourceID string) ([]models.AnalystRating, error) {
	var latestRatings []models.AnalystRating

	query := r.db.Model(&models.AnalystRating{}).Order("rated_at DESC").Limit(5)
	if sourceID != "" {
		query = query.Where("data_source_id = ?", sourceID)
	}
	err := query.Preload("Company").Find(&latestRatings).Error
	if err != nil {
		return nil, err
	}

	return latestRatings, nil
}

func (r *AnalystRatingRepository) GetDonutRatingChart(sourceID string) ([]dtos.DonutChartDTO, error) {
	var donutChart []dtos.DonutChartDTO

	query := r.db.Model(&models.AnalystRating{}).
		Select("rating_to, COUNT(*) as count, MAX(rated_at) as latest_rated_at").
		Group("rating_to").
		Order("count DESC")
	if sourceID != "" {
		query = query.Where("data_source_id = ?", sourceID)
	}

	err := query.Scan(&donutChart).Error
	if err != nil {
		return nil, err
	}

	return donutChart, nil
}

func (r *AnalystRatingRepository) GetRawCPIData(sourceID string) ([]models.AnalystRating, error) {
	var ratings []models.AnalystRating

	query := r.db.Model(&models.AnalystRating{}).Select("combined_prediction_index")
	if sourceID != "" {
		query = query.Where("data_source_id = ?", sourceID)
	}
	query = query.Order("combined_prediction_index")

	err := query.Find(&ratings).Error
	if err != nil {
		return nil, err
	}

	return ratings, nil
}
