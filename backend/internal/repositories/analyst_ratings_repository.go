package repositories

import (
	"backend/internal/dtos"
	"backend/internal/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnalystRatingsRepository struct {
	db *gorm.DB
}

func NewAnalystRatingsRepository(db *gorm.DB) *AnalystRatingsRepository {
	return &AnalystRatingsRepository{db: db}
}

func (r *AnalystRatingsRepository) Create(rating *models.AnalystRating) error {
	return r.db.Create(rating).Error
}

func (r *AnalystRatingsRepository) CreateBatch(ratings []models.AnalystRating) error {
	if len(ratings) == 0 {
		return nil
	}
	return r.db.Create(&ratings).Error
}

func (r *AnalystRatingsRepository) GetByID(id uuid.UUID) (*models.AnalystRating, error) {
	var rating models.AnalystRating
	err := r.db.First(&rating, id).Error
	if err != nil {
		return nil, err
	}
	return &rating, nil
}

func (r *AnalystRatingsRepository) GetAll(sortOrder, sortBy, sourceID string, filters map[string]string, page, limit int) ([]models.AnalystRating, int64, error) {
	var ratings []models.AnalystRating
	var total int64
	query := r.db

	query = applyFilters(query, sourceID, filters)
	query.Model(&models.AnalystRating{}).Count(&total)

	sortBy = cleanSortColumn(sortBy)
	query = applySorting(query, sortBy, sortOrder)
	query = applyPagination(query, page, limit)

	err := query.Find(&ratings).Error
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
		"ticker":                       true,
		"target_from":                  true,
		"target_to":                    true,
		"company":                      true,
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

	query = query.Where("rated_at = (SELECT MAX(rated_at) FROM analyst_ratings AS ar WHERE ar.ticker = analyst_ratings.ticker)")

	return query
}

func applySorting(query *gorm.DB, sortBy, sortOrder string) *gorm.DB {
	return query.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))
}

func applyPagination(query *gorm.DB, page, limit int) *gorm.DB {
	offset := (page - 1) * limit
	return query.Limit(limit).Offset(offset)
}

func (r *AnalystRatingsRepository) GetIndicators(sourceID string) (dtos.AnalystRatingIndicatorsDTO, error) {
	var dto dtos.AnalystRatingIndicatorsDTO

	totalCount, err := r.getTotalRatingsCount(sourceID)
	if err != nil || totalCount == 0 {
		return dto, err
	}

	dto.BuyNowPercentage, err = r.getBuyNowPercentage(sourceID, totalCount)
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
	dto.HighestIncrementInTargetPriceTicker = highestIncrementRating.Ticker

	return dto, nil
}

func (r *AnalystRatingsRepository) getTotalRatingsCount(sourceID string) (int64, error) {
	var totalCount int64
	query := r.db.Model(&models.AnalystRating{})
	if sourceID != "" {
		query = query.Where("data_source_id = ?", sourceID)
	}
	err := query.Count(&totalCount).Error
	return totalCount, err
}

func (r *AnalystRatingsRepository) getBuyNowPercentage(sourceID string, totalCount int64) (float64, error) {
	var buyNowCount int64
	query := r.db.Model(&models.AnalystRating{}).Where("combined_prediction_index > ?", 70)
	if sourceID != "" {
		query = query.Where("data_source_id = ?", sourceID)
	}
	err := query.Count(&buyNowCount).Error
	if err != nil {
		return 0, err
	}
	return float64(buyNowCount) / float64(totalCount) * 100, nil
}

func (r *AnalystRatingsRepository) getPositiveTargetAdjustmentPercentage(sourceID string, totalCount int64) (float64, error) {
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

func (r *AnalystRatingsRepository) getHighestIncrementInTargetPrice(sourceID string) (models.AnalystRating, error) {
	var rating models.AnalystRating

	subQuery := r.db.Model(&models.AnalystRating{}).
		Select("id").
		Order("(target_to - target_from) DESC").
		Limit(1)

	if sourceID != "" {
		subQuery = subQuery.Where("data_source_id = ?", sourceID)
	}

	err := r.db.Where("id = (?)", subQuery).First(&rating).Error
	if err != nil {
		return models.AnalystRating{}, err
	}

	return rating, nil
}

func (r *AnalystRatingsRepository) GetRecommendations() ([]models.AnalystRating, error) {
	// TODO: Change how to get recommendations
	var ratings []models.AnalystRating
	err := r.db.Where("action = ?", "Buy").Find(&ratings).Error
	if err != nil {
		return nil, err
	}
	return ratings, nil
}
