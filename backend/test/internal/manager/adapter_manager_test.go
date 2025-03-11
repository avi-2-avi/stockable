package test

import (
	"backend/config"
	"backend/internal/database"
	"backend/internal/manager"
	"backend/internal/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestAdapterManager_RunAdapters_All(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"items": [], "next_page": ""}`))
	}))
	defer mockServer.Close()

	mockEnvContent := "DATABASE_URL=postgresql://root@localhost:26257/defaultdb?sslmode=disable\n" +
		"TRU_ADAPTER_URL=" + mockServer.URL + "\n" + // Use mock server
		"TRU_ADAPTER_TOKEN=1234567890"
	os.WriteFile(".env", []byte(mockEnvContent), 0644)

	config, _ := config.LoadConfig()
	db, _ := database.Connect()
	db.Config.Logger = logger.Default.LogMode(logger.Silent)
	database.Migrate(db)

	adapterFactory := manager.NewAdapterFactory(config, db)
	adapterManager := manager.NewAdapterManager(adapterFactory, db)
	runAdaptersErr := adapterManager.RunAdapters("")

	assert.NotNil(t, adapterManager)
	assert.NoError(t, runAdaptersErr)

	db.Migrator().DropTable(&models.AnalystRating{}, &models.DataSource{}, &models.AdapterLog{}, &models.User{})
	os.Remove(".env")
	os.Unsetenv("DATABASE_URL")
	os.Unsetenv("TRU_ADAPTER_TOKEN")
	os.Unsetenv("TRU_ADAPTER_URL")
}

func TestAdapterManager_RunAdapters_Specific(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"items": [], "next_page": ""}`))
	}))
	defer mockServer.Close()

	mockEnvContent := "DATABASE_URL=postgresql://root@localhost:26257/defaultdb?sslmode=disable\n" +
		"TRU_ADAPTER_URL=" + mockServer.URL + "\n" + // Use mock server
		"TRU_ADAPTER_TOKEN=1234567890"
	os.WriteFile(".env", []byte(mockEnvContent), 0644)

	config, _ := config.LoadConfig()
	db, _ := database.Connect()
	db.Config.Logger = logger.Default.LogMode(logger.Silent)
	database.Migrate(db)

	adapterFactory := manager.NewAdapterFactory(config, db)
	adapterManager := manager.NewAdapterManager(adapterFactory, db)
	runAdaptersErr := adapterManager.RunAdapters("TruAdapter")

	assert.NotNil(t, adapterManager)
	assert.NoError(t, runAdaptersErr)

	db.Migrator().DropTable(&models.AnalystRating{}, &models.DataSource{}, &models.AdapterLog{}, &models.User{})
	os.Remove(".env")
	os.Unsetenv("DATABASE_URL")
	os.Unsetenv("TRU_ADAPTER_TOKEN")
	os.Unsetenv("TRU_ADAPTER_URL")
}
