package manager

import (
	"backend/config"
	"backend/internal/adapters"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/services"
	"fmt"

	"gorm.io/gorm"
)

type AdapterFactory struct {
	config             *config.Config
	dataSourceRepo     *repositories.DataSourceRepository
	analystRatingsRepo *repositories.AnalystRatingsRepository
}

func NewAdapterFactory(config *config.Config, db *gorm.DB) *AdapterFactory {
	return &AdapterFactory{
		config:             config,
		dataSourceRepo:     repositories.NewDataSourceRepository(db),
		analystRatingsRepo: repositories.NewAnalystRatingsRepository(db),
	}
}

func (f *AdapterFactory) CreateAdapter(name string) adapters.RatingAdapter {
	switch name {
	case "TruAdapter":
		truSource, _ := f.dataSourceRepo.GetByName("TruAdapter")
		if truSource == nil {
			fmt.Printf("Creating new data source: %s\n", "TruAdapter")
			truSource = &models.DataSource{Name: "TruAdapter"}
			if err := f.dataSourceRepo.Create(truSource); err != nil {
				fmt.Println("Error creating DataSource:", err)
				return nil
			}
			fmt.Printf("New DataSource created with ID: %s\n", truSource.ID)
		} else {
			fmt.Printf("Existing DataSource found with ID: %s\n", truSource.ID)
		}

		analystService := services.NewAnalystRatingsService(f.analystRatingsRepo)
		return adapters.NewTruAdapter(
			f.config.TruAdapterURL,
			f.config.TruAdapterToken,
			analystService,
			truSource.ID,
		)
	case "DummyAdapter":
		dummySource, _ := f.dataSourceRepo.GetByName("DummyAdapter")
		if dummySource == nil {
			fmt.Printf("Creating new data source: %s\n", "DummyAdapter")
			dummySource = &models.DataSource{Name: "DummyAdapter"}
			if err := f.dataSourceRepo.Create(dummySource); err != nil {
				fmt.Println("Error creating DataSource:", err)
				return nil
			}
			fmt.Printf("New DataSource created with ID: %s\n", dummySource.ID)
		} else {
			fmt.Printf("Existing DataSource found with ID: %s\n", dummySource.ID)
		}

		analystService := services.NewAnalystRatingsService(f.analystRatingsRepo)
		return adapters.NewDummyAdapter(analystService, dummySource.ID)
	default:
		return nil
	}
}
