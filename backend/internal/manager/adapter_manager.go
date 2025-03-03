package manager

import (
	"backend/config"
	"backend/internal/adapters"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/services"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type AdapterManager struct {
	db             *gorm.DB
	config         *config.Config
	dataSourceRepo *repositories.DataSourceRepository
	adapterLogRepo *repositories.AdapterLogRepository
}

func NewAdapterManager(db *gorm.DB, config *config.Config) *AdapterManager {
	return &AdapterManager{
		db:             db,
		config:         config,
		dataSourceRepo: repositories.NewDataSourceRepository(db),
		adapterLogRepo: repositories.NewAdapterLogRepository(db),
	}
}

func (adapterManager *AdapterManager) RunAdapters(adapterName string) error {
	adaptersToRun := make(map[string]func() adapters.RatingAdapter)

	adaptersToRun["TruAdapter"] = func() adapters.RatingAdapter {
		truAdapterSource, err := adapterManager.dataSourceRepo.GetByName("TruAdapter")
		if err != nil || truAdapterSource == nil {
			fmt.Println("Creating DataSource entry for TruAdapter...")
			truAdapterSource = &models.DataSource{Name: "TruAdapter"}
			adapterManager.dataSourceRepo.Create(truAdapterSource)
		}

		analystRatingsRepo := repositories.NewAnalystRatingsRepository(adapterManager.db)
		analystRatingsService := services.NewAnalystRatingsService(analystRatingsRepo)

		return adapters.NewTruAdapter(
			adapterManager.config.TruAdapterURL,
			adapterManager.config.TruAdapterToken,
			analystRatingsService,
			truAdapterSource.ID,
		)
	}

	if adapterName == "" {
		fmt.Println("Running all adapters...")
	} else {
		fmt.Printf("Running adapter: %s...\n", adapterName)
	}

	for key, createAdapter := range adaptersToRun {
		if adapterName != "" && !strings.EqualFold(adapterName, key) {
			continue
		}

		dataSource, err := adapterManager.dataSourceRepo.GetByName(key)
		if err != nil || dataSource == nil {
			fmt.Printf("Creating DataSource entry for %s...\n", key)
			dataSource = &models.DataSource{Name: key}
			adapterManager.dataSourceRepo.Create(dataSource)
		}

		adapter := createAdapter()
		if adapter == nil {
			fmt.Printf("ERROR: Adapter factory returned nil for %s\n", key)
			continue
		}
		fmt.Printf("Adapter created: %+v\n", adapter)

		_, err = adapter.FetchData()
		if err != nil {
			fmt.Printf("Error fetching ratings for %s: %v\n", key, err)
			continue
		}

		log := &models.AdapterLog{
			AdapterName: key,
			RunAt:       time.Now(),
		}
		adapterManager.adapterLogRepo.Create(log)
		fmt.Printf("Successfully loaded analyst ratings from %s.\n", key)
	}

	return nil
}
