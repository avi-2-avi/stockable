package repositories

import (
	"backend/internal/models"
	"fmt"

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

func (r *AnalystRatingsRepository) GetByID(id uint) (*models.AnalystRating, error) {
	var rating models.AnalystRating
	err := r.db.First(&rating, id).Error
	if err != nil {
		return nil, err
	}
	return &rating, nil
}

func (r *AnalystRatingsRepository) GetAll(sortOrder, sortBy, sourceID string, filters map[string]string, page, limit int) ([]models.AnalystRating, error) {
	var ratings []models.AnalystRating
	query := r.db

	sortBy = cleanSortColumn(sortBy)
	query = applyFilters(query, sourceID, filters)
	query = applySorting(query, sortBy, sortOrder)
	query = applyPagination(query, page, limit)

	err := query.Find(&ratings).Error
	if err != nil {
		return nil, err
	}

	return ratings, nil
}

func cleanSortColumn(sortBy string) string {
	allowedColumns := allowedActionColumns()
	if _, exists := allowedColumns[sortBy]; !exists {
		return "rated_at"
	}
	return sortBy
}

func allowedActionColumns() map[string]bool {
	return map[string]bool{
		"ticker":      true,
		"target_from": true,
		"target_to":   true,
		"company":     true,
		"action":      true,
		"brokerage":   true,
		"rating_from": true,
		"rating_to":   true,
		"created_at":  true,
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

	return query
}

func applySorting(query *gorm.DB, sortBy, sortOrder string) *gorm.DB {
	return query.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))
}

func applyPagination(query *gorm.DB, page, limit int) *gorm.DB {
	offset := (page - 1) * limit
	return query.Limit(limit).Offset(offset)
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
