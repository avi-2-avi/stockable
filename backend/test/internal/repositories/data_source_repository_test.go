package test

import (
	"data-loader/internal/database"
	"data-loader/internal/models"
	"data-loader/internal/repositories"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestDataSourceRepository(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")
	db, dbErr := database.Connect()
	db.Config.Logger = logger.Default.LogMode(logger.Silent)
	migrationErr := database.Migrate(db)

	dataSourceRepo := repositories.NewDataSourceRepository(db)
	dataSource := models.DataSource{Name: "API-TEST"}
	repoErr := dataSourceRepo.Create(&dataSource)
	fetchedSource, fetchErr := dataSourceRepo.GetByID(dataSource.ID)

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.NoError(t, migrationErr, "Database migration should not return an error")
	assert.NoError(t, repoErr, "Should create data source without error")
	assert.NotZero(t, dataSource.ID, "DataSource ID should be set")
	assert.NoError(t, fetchErr, "Should fetch data source without error")
	assert.Equal(t, "API-TEST", fetchedSource.Name, "Fetched data source should match")

	db.Migrator().DropTable(&models.AnalystRating{}, &models.DataSource{})
	os.Unsetenv("DATABASE_URL")
}
