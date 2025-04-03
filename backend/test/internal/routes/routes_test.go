package test

import (
	"backend/config"
	"backend/internal/database"
	"backend/internal/repositories"
	"backend/internal/routes"
	"bytes"
	"encoding/json"
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

	reqBody := `{
		"email": "test@example.com",
		"password": "password123",
		"full_name": "Test User",
		"role_name": "user"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.Equal(t, http.StatusCreated, w.Code, "Expected HTTP 201 - Created")
	assert.NotNil(t, w.Body, "Response body should not be nil")

	defer func() {
		db.Exec("DELETE FROM users WHERE email = ?", "test@example.com")
	}()

	os.Remove(".env")
}

func TestPostAuthLogin(t *testing.T) {
	mockEnvContent := "DATABASE_URL=postgresql://root@localhost:26257/defaultdb?sslmode=disable\n" +
		"ALLOWED_ORIGIN=http://localhost:5173\n" +
		"TRU_ADAPTER_URL=http://localhost:8080\n" +
		"TRU_ADAPTER_TOKEN=1234567890"

	os.WriteFile(".env", []byte(mockEnvContent), 0644)
	config.LoadConfig()
	db, dbErr := database.Connect()

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")

	router := routes.SetupRouter()
	routes.RegisterRoutes(router, db)

	reqBody := `{
		"email": "test@example.com",
		"password": "password123",
		"full_name": "Test User",
		"role_name": "user"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	reqBody = `{
		"email": "test@example.com",
		"password": "password123"
	}`

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 - OK")
	assert.NotNil(t, w.Body, "Response body should not be nil")

	authRepo := repositories.NewAuthRepository(db)
	user, err := authRepo.GetUserByEmail("test@example.com")
	if err != nil || user == nil {
		t.Fatalf("User retrieval failed: %v", err)
	}

	defer func() {
		db.Exec("DELETE FROM users WHERE email = ?", "test@example.com")
	}()

	os.Remove(".env")
}

func TestPostAuthLogout(t *testing.T) {
	mockEnvContent := "DATABASE_URL=postgresql://root@localhost:26257/defaultdb?sslmode=disable\n" +
		"ALLOWED_ORIGIN=http://localhost:5173\n" +
		"TRU_ADAPTER_URL=http://localhost:8080\n" +
		"TRU_ADAPTER_TOKEN=1234567890"

	os.WriteFile(".env", []byte(mockEnvContent), 0644)
	config.LoadConfig()
	db, dbErr := database.Connect()

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")

	router := routes.SetupRouter()
	routes.RegisterRoutes(router, db)

	reqBody := `{
		"email": "test@example.com",
		"password": "password123",
		"full_name": "Test User",
		"role_name": "user"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	reqBody = `{
		"email": "test@example.com",
		"password": "password123"
	}`

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/auth/logout", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200")
	assert.NotNil(t, w.Body, "Response body should not be nil")

	defer func() {
		db.Exec("DELETE FROM users WHERE email = ?", "test@example.com")
	}()

	os.Remove(".env")
}

func TestGetUsers(t *testing.T) {
	mockEnvContent := "DATABASE_URL=postgresql://root@localhost:26257/defaultdb?sslmode=disable\n" +
		"ALLOWED_ORIGIN=http://localhost:5173\n" +
		"TRU_ADAPTER_URL=http://localhost:8080\n" +
		"TRU_ADAPTER_TOKEN=1234567890"

	os.WriteFile(".env", []byte(mockEnvContent), 0644)
	config.LoadConfig()
	db, dbErr := database.Connect()

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")

	router := routes.SetupRouter()
	routes.RegisterRoutes(router, db)

	reqBody := `{
		"email": "test@example.com",
		"password": "password123",
		"full_name": "Test User",
		"role_name": "admin"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	reqBody = `{
		"email": "test@example.com",
		"password": "password123"
	}`

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	cookies := w.Result().Cookies()
	var authCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "auth_token" {
			authCookie = cookie
			break
		}
	}

	assert.NotNil(t, authCookie, "Auth cookie should be present in login response")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/auth/list", nil)
	req.AddCookie(authCookie)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200")
	assert.NotNil(t, w.Body, "Response body should not be nil")

	defer func() {
		db.Exec("DELETE FROM users WHERE email = ?", "test@example.com")
	}()

	os.Remove(".env")
}

func TestPatchUser(t *testing.T) {
	mockEnvContent := "DATABASE_URL=postgresql://root@localhost:26257/defaultdb?sslmode=disable\n" +
		"ALLOWED_ORIGIN=http://localhost:5173\n" +
		"TRU_ADAPTER_URL=http://localhost:8080\n" +
		"TRU_ADAPTER_TOKEN=1234567890"

	os.WriteFile(".env", []byte(mockEnvContent), 0644)
	config.LoadConfig()
	db, dbErr := database.Connect()

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")

	router := routes.SetupRouter()
	routes.RegisterRoutes(router, db)

	reqBody := `{
		"email": "test@example.com",
		"password": "password123",
		"full_name": "Test User",
		"role_name": "admin"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	var registerResponse map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&registerResponse)
	assert.NoError(t, err, "Error decoding register response")

	bodyField, bodyExists := registerResponse["body"]
	assert.True(t, bodyExists, "'body' field should be present in the response")

	userID := bodyField.(map[string]interface{})["id"].(string)
	assert.NotNil(t, userID, "User ID should not be nil")

	oldReqBody := `{
		"email": "test@example.com",
		"password": "password123"
	}`

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(oldReqBody)))
	router.ServeHTTP(w, req)

	cookies := w.Result().Cookies()
	var authCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "auth_token" {
			authCookie = cookie
			break
		}
	}

	assert.NotNil(t, authCookie, "Auth cookie should be present in login response")

	reqBody = `{
		"email": "test@example123.com",
		"password": "password123123",
		"full_name": "Updated User"
	}`

	w = httptest.NewRecorder()
	t.Log("User id", userID)
	req, _ = http.NewRequest("PATCH", "/api/auth/update/"+userID, bytes.NewBuffer([]byte(reqBody)))
	req.AddCookie(authCookie)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 - OK")
	assert.NotNil(t, w.Body, "Response body should not be nil")

	defer func() {
		db.Exec("DELETE FROM users WHERE email = ?", "test@example.com")
	}()

	os.Remove(".env")
}

func TestDeleteUser(t *testing.T) {
	mockEnvContent := "DATABASE_URL=postgresql://root@localhost:26257/defaultdb?sslmode=disable\n" +
		"ALLOWED_ORIGIN=http://localhost:5173\n" +
		"TRU_ADAPTER_URL=http://localhost:8080\n" +
		"TRU_ADAPTER_TOKEN=1234567890"

	os.WriteFile(".env", []byte(mockEnvContent), 0644)
	config.LoadConfig()
	db, dbErr := database.Connect()

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")

	router := routes.SetupRouter()
	routes.RegisterRoutes(router, db)

	reqBody := `{
		"email": "test@example.com",
		"password": "password123",
		"full_name": "Test User",
		"role_name": "admin"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	var registerResponse map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&registerResponse)
	assert.NoError(t, err, "Error decoding register response")

	bodyField, bodyExists := registerResponse["body"]
	assert.True(t, bodyExists, "'body' field should be present in the response")

	userID := bodyField.(map[string]interface{})["id"].(string)
	assert.NotNil(t, userID, "User ID should not be nil")

	oldReqBody := `{
	"email": "test@example.com",
	"password": "password123"
}`

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(oldReqBody)))
	router.ServeHTTP(w, req)

	cookies := w.Result().Cookies()
	var authCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "auth_token" {
			authCookie = cookie
			break
		}
	}

	assert.NotNil(t, authCookie, "Auth cookie should be present in login response")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/api/auth/delete/"+userID, nil)
	req.AddCookie(authCookie)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 - OK")
	assert.NotNil(t, w.Body, "Response body should not be nil")

	defer func() {
		db.Exec("DELETE FROM users WHERE email = ?", "test@example.com")
	}()

	os.Remove(".env")
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

	reqBody := `{
		"email": "test@example.com",
		"password": "password123",
		"full_name": "Test User",
		"role_name": "admin"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	var registerResponse map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&registerResponse)
	assert.NoError(t, err, "Error decoding register response")

	bodyField, bodyExists := registerResponse["body"]
	assert.True(t, bodyExists, "'body' field should be present in the response")

	userID := bodyField.(map[string]interface{})["id"].(string)
	assert.NotNil(t, userID, "User ID should not be nil")

	reqBody = `{
	"email": "test@example.com",
	"password": "password123"
	}`

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	cookies := w.Result().Cookies()
	var authCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "auth_token" {
			authCookie = cookie
			break
		}
	}

	assert.NotNil(t, authCookie, "Auth cookie should be present in login response")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/ratings", nil)
	req.AddCookie(authCookie)
	router.ServeHTTP(w, req)

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200")
	assert.NotNil(t, w.Body, "Response body should not be nil")

	defer func() {
		db.Exec("DELETE FROM users WHERE email = ?", "test@example.com")
	}()

	os.Remove(".env")
}

func TestGetAnalystRatingsIndicators(t *testing.T) {
	mockEnvContent := "DATABASE_URL=postgresql://root@localhost:26257/defaultdb?sslmode=disable\n" +
		"ALLOWED_ORIGIN=http://localhost:5173\n" +
		"TRU_ADAPTER_URL=http://localhost:8080\n" +
		"TRU_ADAPTER_TOKEN=1234567890"

	os.WriteFile(".env", []byte(mockEnvContent), 0644)
	config.LoadConfig()
	db, _ := database.Connect()

	router := routes.SetupRouter()
	routes.RegisterRoutes(router, db)

	reqBody := `{
		"email": "test@example.com",
		"password": "password123",
		"full_name": "Test User",
		"role_name": "admin"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	var registerResponse map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&registerResponse)
	assert.NoError(t, err, "Error decoding register response")

	bodyField, bodyExists := registerResponse["body"]
	assert.True(t, bodyExists, "'body' field should be present in the response")

	userID := bodyField.(map[string]interface{})["id"].(string)
	assert.NotNil(t, userID, "User ID should not be nil")

	reqBody = `{
	"email": "test@example.com",
	"password": "password123"
	}`

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	cookies := w.Result().Cookies()
	var authCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "auth_token" {
			authCookie = cookie
			break
		}
	}

	assert.NotNil(t, authCookie, "Auth cookie should be present in login response")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/ratings/indicators", nil)
	req.AddCookie(authCookie)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200")
	assert.NotNil(t, w.Body, "Response body should not be nil")

	defer func() {
		db.Exec("DELETE FROM users WHERE email = ?", "test@example.com")
	}()

	os.Remove(".env")
}

func TestGetSources(t *testing.T) {
	mockEnvContent := "DATABASE_URL=postgresql://root@localhost:26257/defaultdb?sslmode=disable\n" +
		"ALLOWED_ORIGIN=http://localhost:5173\n" +
		"TRU_ADAPTER_URL=http://localhost:8080\n" +
		"TRU_ADAPTER_TOKEN=1234567890"

	os.WriteFile(".env", []byte(mockEnvContent), 0644)
	config.LoadConfig()
	db, _ := database.Connect()

	router := routes.SetupRouter()
	routes.RegisterRoutes(router, db)

	reqBody := `{
		"email": "test@example.com",
		"password": "password123",
		"full_name": "Test User",
		"role_name": "admin"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	var registerResponse map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&registerResponse)
	assert.NoError(t, err, "Error decoding register response")

	bodyField, bodyExists := registerResponse["body"]
	assert.True(t, bodyExists, "'body' field should be present in the response")

	userID := bodyField.(map[string]interface{})["id"].(string)
	assert.NotNil(t, userID, "User ID should not be nil")

	reqBody = `{
	"email": "test@example.com",
	"password": "password123"
	}`

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	cookies := w.Result().Cookies()
	var authCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "auth_token" {
			authCookie = cookie
			break
		}
	}

	assert.NotNil(t, authCookie, "Auth cookie should be present in login response")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/sources", nil)
	req.AddCookie(authCookie)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200")
	assert.NotNil(t, w.Body, "Response body should not be nil")

	defer func() {
		db.Exec("DELETE FROM users WHERE email = ?", "test@example.com")
	}()

	os.Remove(".env")
}

func TestPatchSource(t *testing.T) {
	mockEnvContent := "DATABASE_URL=postgresql://root@localhost:26257/defaultdb?sslmode=disable\n" +
		"ALLOWED_ORIGIN=http://localhost:5173\n" +
		"TRU_ADAPTER_URL=http://localhost:8080\n" +
		"TRU_ADAPTER_TOKEN=1234567890"

	os.WriteFile(".env", []byte(mockEnvContent), 0644)
	config.LoadConfig()
	db, _ := database.Connect()

	router := routes.SetupRouter()
	routes.RegisterRoutes(router, db)

	reqBody := `{
		"email": "test@example.com",
		"password": "password123",
		"full_name": "Test User",
		"role_name": "admin"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	var registerResponse map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&registerResponse)
	assert.NoError(t, err, "Error decoding register response")

	bodyField, bodyExists := registerResponse["body"]
	assert.True(t, bodyExists, "'body' field should be present in the response")

	userID := bodyField.(map[string]interface{})["id"].(string)
	assert.NotNil(t, userID, "User ID should not be nil")

	reqBody = `{
	"email": "test@example.com",
	"password": "password123"
	}`

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	cookies := w.Result().Cookies()
	var authCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "auth_token" {
			authCookie = cookie
			break
		}
	}

	assert.NotNil(t, authCookie, "Auth cookie should be present in login response")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/sources", nil)
	req.AddCookie(authCookie)
	router.ServeHTTP(w, req)

	var sourcesResponse map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&sourcesResponse)
	assert.NoError(t, err, "Error decoding sources response")

	bodyField, bodyExists = sourcesResponse["body"]
	assert.True(t, bodyExists, "'body' field should be present in the response")

	sourceList := bodyField.([]interface{})
	assert.Greater(t, len(sourceList), 0, "Source list should not be empty")

	sourceID := sourceList[0].(map[string]interface{})["id"].(string)
	assert.NotNil(t, sourceID, "Source ID should not be nil")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PATCH", "/api/sources/"+sourceID, bytes.NewBuffer([]byte(`{"is_visible": false}`)))
	req.AddCookie(authCookie)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200")
	assert.NotNil(t, w.Body, "Response body should not be nil")

	defer func() {
		db.Exec("DELETE FROM users WHERE email = ?", "test@example.com")
	}()

	os.Remove(".env")
}

func TestGetRatingsDashboard(t *testing.T) {
	mockEnvContent := "DATABASE_URL=postgresql://root@localhost:26257/defaultdb?sslmode=disable\n" +
		"ALLOWED_ORIGIN=http://localhost:5173\n" +
		"TRU_ADAPTER_URL=http://localhost:8080\n" +
		"TRU_ADAPTER_TOKEN=1234567890"

	os.WriteFile(".env", []byte(mockEnvContent), 0644)
	config.LoadConfig()
	db, _ := database.Connect()

	router := routes.SetupRouter()
	routes.RegisterRoutes(router, db)

	reqBody := `{
		"email": "test@example.com",
		"password": "password123",
		"full_name": "Test User",
		"role_name": "admin"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	var registerResponse map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&registerResponse)
	assert.NoError(t, err, "Error decoding register response")

	bodyField, bodyExists := registerResponse["body"]
	assert.True(t, bodyExists, "'body' field should be present in the response")

	userID := bodyField.(map[string]interface{})["id"].(string)
	assert.NotNil(t, userID, "User ID should not be nil")

	reqBody = `{
	"email": "test@example.com",
	"password": "password123"
	}`

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(reqBody)))
	router.ServeHTTP(w, req)

	cookies := w.Result().Cookies()
	var authCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "auth_token" {
			authCookie = cookie
			break
		}
	}

	assert.NotNil(t, authCookie, "Auth cookie should be present in login response")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/ratings/dashboard", nil)
	req.AddCookie(authCookie)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200")
	assert.NotNil(t, w.Body, "Response body should not be nil")

	defer func() {
		db.Exec("DELETE FROM users WHERE email = ?", "test@example.com")
	}()

	os.Remove(".env")
}
