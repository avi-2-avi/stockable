package test

import (
	"data-loader/internal/database"
	"data-loader/internal/models"
	"data-loader/internal/repositories"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestAnalystRatingsRepository(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")
	db, dbErr := database.Connect()
	db.Config.Logger = logger.Default.LogMode(logger.Silent)
	migrationErr := database.Migrate(db)
	dataSourceRepo := repositories.NewDataSourceRepository(db)

	ratingRepo := repositories.NewAnalystRatingsRepository(db)
	dataSource := models.DataSource{Name: "API-TEST"}
	dataSourceRepoErr := dataSourceRepo.Create(&dataSource)
	rating := models.AnalystRating{
		Ticker:       "AAPL",
		TargetFrom:   120.5,
		TargetTo:     130.0,
		Company:      "Apple Inc.",
		Action:       "Upgrade",
		Brokerage:    "JP Morgan",
		RatingFrom:   "Hold",
		RatingTo:     "Buy",
		RatedAt:      time.Now(),
		DataSourceID: dataSource.ID,
	}
	ratingRepoErr := ratingRepo.Create(&rating)
	fetchedRating, fetchErr := ratingRepo.GetByID(rating.ID)

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.NoError(t, migrationErr, "Database migration should not return an error")
	assert.NoError(t, dataSourceRepoErr, "Should create data source without error")
	assert.NoError(t, ratingRepoErr, "Should create rating history entry")
	assert.NotZero(t, dataSource.ID, "DataSource ID should be set")
	assert.NoError(t, fetchErr, "Should fetch rating history without error")
	assert.Equal(t, "AAPL", fetchedRating.Ticker, "Fetched rating should match")
	assert.Equal(t, dataSource.ID, fetchedRating.DataSourceID, "Should have correct foreign key")

	db.Migrator().DropTable(&models.AnalystRating{}, &models.DataSource{})
	os.Unsetenv("DATABASE_URL")
}
