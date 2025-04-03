package test

import (
	"backend/config"
	"backend/internal/adapters"
	"backend/internal/database"
	"backend/internal/manager"
	"backend/internal/services"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestAdapterFactory_CreateAdapter(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"items": [], "next_page": ""}`))
	}))
	defer mockServer.Close()

	mockEnvContent := "DATABASE_URL=postgresql://root@localhost:26257/defaultdb?sslmode=disable\n" +
		"TRU_ADAPTER_URL=" + mockServer.URL + "\n" +
		"TRU_ADAPTER_TOKEN=1234567890"
	os.WriteFile(".env", []byte(mockEnvContent), 0644)

	config, _ := config.LoadConfig()
	db, _ := database.Connect()
	db.Config.Logger = logger.Default.LogMode(logger.Silent)
	database.Migrate(db)

	t.Cleanup(func() {
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

	adapterFactory := manager.NewAdapterFactory(config, db)
	adapterFactory.RegisterAdapter("TruAdapter", func(factory *manager.AdapterFactory) adapters.RatingAdapter {
		source := factory.CreateDataSource("TruAdapter", false)
		if source == nil {
			return nil
		}

		analystService := services.NewAnalystRatingService(factory.GetAnalystRatingRepository())
		companyService := services.NewCompanyService(factory.GetCompanyRepository())
		return adapters.NewTruAdapter(
			config.TruAdapterURL,
			config.TruAdapterToken,
			analystService,
			companyService,
			source.ID,
		)
	})
	adapter, err := adapterFactory.CreateAdapter("TruAdapter")

	assert.NotNil(t, adapterFactory)
	assert.NoError(t, err)
	assert.NotNil(t, adapter)

}
