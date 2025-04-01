package test

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/repositories"
	"os"
	"strings"
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
	dataSource := models.DataSource{Name: "API-TEST"}
	dataSourceRepoErr := dataSourceRepo.Create(&dataSource)
	companyRepo := repositories.NewCompanyRepository(db)
	company := models.Company{Name: "Company-TEST"}
	_ = companyRepo.Create(&company)

	ratingRepo := repositories.NewAnalystRatingsRepository(db)
	rating := models.AnalystRating{
		TargetFrom:   120.5,
		TargetTo:     130.0,
		Action:       "Upgrade",
		Brokerage:    "JP Morgan",
		RatingFrom:   "Hold",
		RatingTo:     "Buy",
		RatedAt:      time.Now(),
		DataSourceID: dataSource.ID,
		CompanyID:    company.ID,
	}
	ratingRepoErr := ratingRepo.Create(&rating)
	fetchedRating, fetchErr := ratingRepo.GetByID(rating.ID)

	t.Cleanup(func() {
		ratingRepo.Delete(rating.ID)
		companyRepo.Delete(company.ID)
		dataSourceRepo.Delete(dataSource.ID)

		tables, err := db.Migrator().GetTables()
		if err == nil {
			for _, table := range tables {
				_ = db.Migrator().DropTable(table)
			}
		}

		for _, env := range os.Environ() {
			key := env[:strings.Index(env, "=")]
			os.Unsetenv(key)
		}

		os.Remove(".env")
	})

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.NoError(t, migrationErr, "Database migration should not return an error")
	assert.NoError(t, dataSourceRepoErr, "Should create data source without error")
	assert.NoError(t, ratingRepoErr, "Should create rating history entry")
	assert.NotZero(t, dataSource.ID, "DataSource ID should be set")
	assert.NoError(t, fetchErr, "Should fetch rating history without error")
	assert.Equal(t, dataSource.ID, fetchedRating.DataSourceID, "Should have correct foreign key")
}
