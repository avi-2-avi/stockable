package test

import (
	"backend/internal/adapters"
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/services"
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

	repositories.NewAdapterLogRepository(db)
	dataSourceRepo := repositories.NewDataSourceRepository(db)
	analystRatingRepo := repositories.NewAnalystRatingRepository(db)
	analystRatingService := services.NewAnalystRatingService(analystRatingRepo)
	companyService := services.NewCompanyService(repositories.NewCompanyRepository(db))

	dataSource := models.DataSource{Name: "Tru"}
	dataSourceRepoErr := dataSourceRepo.Create(&dataSource)

	pageCount := 0
	mockToken := "mock-bearer-token"

	mockServer := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		authHeader := request.Header.Get("Authorization")
		expectedAuth := "Bearer " + mockToken
		assert.Equal(t, expectedAuth, authHeader, "Authorization header should contain the correct Bearer token")

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

	adapter := adapters.NewTruAdapter(mockServer.URL, mockToken, analystRatingService, companyService, dataSource.ID)

	_, adapterErr := adapter.FetchData()

	var savedRatings []models.AnalystRating
	db.Find(&savedRatings)

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.NoError(t, migrationErr, "Database migration should not return an error")
	assert.NoError(t, dataSourceRepoErr, "Should create data source without error")
	assert.NoError(t, adapterErr, "Fetching ratings should not return an error")

	assert.GreaterOrEqual(t, 2, len(savedRatings), "There should be 2 analyst ratings saved")

	db.Migrator().DropTable(&models.AnalystRating{}, &models.DataSource{}, &models.AdapterLog{}, &models.User{})
	os.Unsetenv("DATABASE_URL")
}
