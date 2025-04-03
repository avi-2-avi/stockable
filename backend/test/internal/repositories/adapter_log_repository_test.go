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

	t.Cleanup(func() {
		adapterLogRepo.Delete(adapterLog.ID)

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
	assert.NoError(t, adapterLogRepoErr, "Should create data source without error")
}
