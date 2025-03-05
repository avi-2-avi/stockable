package test

import (
	"backend/config"
	"backend/internal/database"
	"backend/internal/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerHealthCheck(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200")
	assert.Equal(t, `{"status":"ok"}`, w.Body.String(), "Expected JSON response")
}

func TestGetAnalystRatings(t *testing.T) {
	mockEnvContent := "DATABASE_URL=postgresql://root@localhost:26257/defaultdb?sslmode=disable\n" +
		"ALLOWED_ORIGIN=http://localhost:5173\n" +
		"TRU_ADAPTER_URL=http://localhost:8080\n" +
		"TRU_ADAPTER_TOKEN=1234567890"

	os.WriteFile(".env", []byte(mockEnvContent), 0644)
	config.LoadConfig()
	db, dbErr := database.Connect()

	router := routes.SetupRouter()
	routes.RegisterRoutes(router, db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ratings", nil)
	router.ServeHTTP(w, req)

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200")

	os.Remove(".env")
	os.Unsetenv("DATABASE_URL")
	os.Unsetenv("ALLOWED_ORIGIN")
	os.Unsetenv("TRU_ADAPTER_URL")
	os.Unsetenv("TRU_ADAPTER_TOKEN")
}
