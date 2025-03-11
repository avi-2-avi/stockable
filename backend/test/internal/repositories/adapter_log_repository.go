package test

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/repositories"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestAdapterLogRepository(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")
	db, dbErr := database.Connect()
	db.Config.Logger = logger.Default.LogMode(logger.Silent)
	migrationErr := database.Migrate(db)
	adapterLogRepo := repositories.NewAdapterLogRepository(db)
	adapterLog := models.AdapterLog{
		AdapterName: "API-TEST",
		RunAt:       time.Now(),
	}

	adapterLogRepoErr := adapterLogRepo.Create(&adapterLog)

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.NoError(t, migrationErr, "Database migration should not return an error")
	assert.NoError(t, adapterLogRepoErr, "Should create data source without error")

	adapterLogRepo.Delete(adapterLog.ID)
	db.Migrator().DropTable(&models.AnalystRating{}, &models.DataSource{}, &models.AdapterLog{}, &models.User{})
	os.Unsetenv("DATABASE_URL")
}
