package test

import (
	"data-loader/internal/database"
	"data-loader/internal/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestMigrations(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")
	db, dbErr := database.Connect()
	db.Config.Logger = logger.Default.LogMode(logger.Silent)

	migrationErr := database.Migrate(db)
	hasAnalystRatings := db.Migrator().HasTable(&models.AnalystRating{})
	hasDataSource := db.Migrator().HasTable(&models.DataSource{})

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.NoError(t, migrationErr, "Database migration should not return an error")
	assert.True(t, hasAnalystRatings, "rating_history table should exist")
	assert.True(t, hasDataSource, "data_source table should exist")

	db.Migrator().DropTable(&models.AnalystRating{}, &models.DataSource{})
}
