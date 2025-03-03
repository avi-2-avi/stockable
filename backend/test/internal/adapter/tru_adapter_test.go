package test

import (
	"data-loader/internal/adapters"
	"data-loader/internal/database"
	"data-loader/internal/models"
	"data-loader/internal/repositories"
	"data-loader/internal/services"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func mockAPIResponse(nextPage string) string {
	data := map[string]interface{}{
		"items": []map[string]string{
			{
				"ticker":      "SDGR",
				"target_from": "$50.00",
				"target_to":   "$45.00",
				"company":     "Schrödinger",
				"action":      "target lowered by",
				"brokerage":   "Piper Sandler",
				"rating_from": "Overweight",
				"rating_to":   "Overweight",
				"time":        "2025-02-28T00:30:09.297349502Z",
			},
		},
		"next_page": nextPage,
	}
	response, _ := json.Marshal(data)
	return string(response)
}

func TestTruAdapter_FetchData(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")
	db, dbErr := database.Connect()
	db.Config.Logger = logger.Default.LogMode(logger.Silent)
	migrationErr := database.Migrate(db)

	dataSourceRepo := repositories.NewDataSourceRepository(db)
	analystRatingsRepo := repositories.NewAnalystRatingsRepository(db)

	analystRatingsService := services.NewAnalystRatingsService(analystRatingsRepo)

	dataSource := models.DataSource{Name: "TruAdapter"}
	dataSourceRepoErr := dataSourceRepo.Create(&dataSource)

	pageCount := 0
	mockServer := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)

		if pageCount == 0 {
			_, _ = responseWriter.Write([]byte(mockAPIResponse("NextPageToken")))
		} else {
			_, _ = responseWriter.Write([]byte(mockAPIResponse("")))
		}
		pageCount++
	}))
	defer mockServer.Close()

	adapter := adapters.NewTruAdapter(mockServer.URL, analystRatingsService, dataSource.ID)

	_, adapterErr := adapter.FetchData()

	var savedRatings []models.AnalystRating
	db.Find(&savedRatings)

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.NoError(t, migrationErr, "Database migration should not return an error")
	assert.NoError(t, dataSourceRepoErr, "Should create data source without error")
	assert.NoError(t, adapterErr, "Fetching ratings should not return an error")

	assert.Equal(t, 2, len(savedRatings), "There should be 2 analyst ratings saved")
	assert.Equal(t, "SDGR", savedRatings[0].Ticker, "Ticker should match")
	assert.Equal(t, "Schrödinger", savedRatings[0].Company, "Company name should match")
	assert.Equal(t, dataSource.ID, savedRatings[0].DataSourceID, "Should have correct DataSourceID")

	db.Migrator().DropTable(&models.AnalystRating{}, &models.DataSource{})
	os.Unsetenv("DATABASE_URL")
}
