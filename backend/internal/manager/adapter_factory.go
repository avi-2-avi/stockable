package manager

import (
	"backend/config"
	"backend/internal/adapters"
	"backend/internal/models"
	"backend/internal/repositories"
	"fmt"

	"gorm.io/gorm"
)

type AdapterConstructor func(*AdapterFactory) adapters.RatingAdapter

type AdapterFactory struct {
	config             *config.Config
	dataSourceRepo     *repositories.DataSourceRepository
	analystRatingsRepo *repositories.AnalystRatingsRepository
	companyRepo        *repositories.CompanyRepository
	adapterRegistry    map[string]AdapterConstructor
}

func NewAdapterFactory(config *config.Config, db *gorm.DB) *AdapterFactory {
	factory := &AdapterFactory{
		config:             config,
		dataSourceRepo:     repositories.NewDataSourceRepository(db),
		analystRatingsRepo: repositories.NewAnalystRatingsRepository(db),
		companyRepo:        repositories.NewCompanyRepository(db),
		adapterRegistry:    make(map[string]AdapterConstructor),
	}
	return factory
}

func (f *AdapterFactory) RegisterAdapter(name string, constructor AdapterConstructor) {
	f.adapterRegistry[name] = constructor
}

func (f *AdapterFactory) CreateAdapter(name string) (adapters.RatingAdapter, error) {
	fmt.Printf("Creating adapter: %s\n", name)
	constructor, exists := f.adapterRegistry[name]
	if !exists {
		return nil, fmt.Errorf("adapter not found: %s", name)
	}
	return constructor(f), nil
}

func (f *AdapterFactory) GetAnalystRatingsRepo() *repositories.AnalystRatingsRepository {
	return f.analystRatingsRepo
}

func (f *AdapterFactory) GetCompanyRepository() *repositories.CompanyRepository {
	return f.companyRepo
}

func (f *AdapterFactory) CreateDataSource(name string, isVisible bool) *models.DataSource {
	fmt.Printf("Creating data source: %s\n", name)
	dataSource, err := f.dataSourceRepo.GetByName(name)
	if err != nil {
		fmt.Printf("Error retrieving DataSource '%s': %v\n", name, err)
		return nil
	}

	if dataSource == nil {
		fmt.Printf("Creating new data source: %s\n", name)
		dataSource = &models.DataSource{
			Name:      name,
			IsVisible: isVisible,
		}
		if err := f.dataSourceRepo.Create(dataSource); err != nil {
			fmt.Println("Error creating DataSource:", err)
			return nil
		}
		fmt.Printf("New DataSource created with ID: %s\n", dataSource.ID)
	} else {
		fmt.Printf("Existing DataSource found with ID: %s\n", dataSource.ID)
	}
	return dataSource
}
