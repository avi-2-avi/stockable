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

func TestPostAuthRegister(t *testing.T) {
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
	req, _ := http.NewRequest("POST", "/auth/register", nil)
	router.ServeHTTP(w, req)

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Expected HTTP 500 - Invalid Request")
	assert.NotNil(t, w.Body, "Response body should not be nil")
}

func TestPostAuthLogin(t *testing.T) {
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
	req, _ := http.NewRequest("POST", "/auth/login", nil)
	router.ServeHTTP(w, req)

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Expected HTTP 500 - Invalid Request")
	assert.NotNil(t, w.Body, "Response body should not be nil")
}

func TestPostAuthLogout(t *testing.T) {
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
	req, _ := http.NewRequest("POST", "/auth/logout", nil)
	router.ServeHTTP(w, req)

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200")
	assert.NotNil(t, w.Body, "Response body should not be nil")
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
	assert.NotNil(t, w.Body, "Response body should not be nil")

	os.Remove(".env")
	os.Unsetenv("DATABASE_URL")
	os.Unsetenv("ALLOWED_ORIGIN")
	os.Unsetenv("TRU_ADAPTER_URL")
	os.Unsetenv("TRU_ADAPTER_TOKEN")
}

func TestGetAnalystRatingsIndicators(t *testing.T) {
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
	req, _ := http.NewRequest("GET", "/ratings/indicators", nil)
	router.ServeHTTP(w, req)

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200")
	assert.NotNil(t, w.Body, "Response body should not be nil")

	os.Remove(".env")
	os.Unsetenv("DATABASE_URL")
	os.Unsetenv("ALLOWED_ORIGIN")
	os.Unsetenv("TRU_ADAPTER_URL")
	os.Unsetenv("TRU_ADAPTER_TOKEN")
}

func TestGetSources(t *testing.T) {
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
	req, _ := http.NewRequest("GET", "/sources", nil)
	router.ServeHTTP(w, req)

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200")
	assert.NotNil(t, w.Body, "Response body should not be nil")

	os.Remove(".env")
	os.Unsetenv("DATABASE_URL")
	os.Unsetenv("ALLOWED_ORIGIN")
	os.Unsetenv("TRU_ADAPTER_URL")
	os.Unsetenv("TRU_ADAPTER_TOKEN")
}
