package manager

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AdapterManager struct {
	adapterFactory *AdapterFactory
	dataSourceRepo *repositories.DataSourceRepository
	adapterLogRepo *repositories.AdapterLogRepository
}

func NewAdapterManager(factory *AdapterFactory, db *gorm.DB) *AdapterManager {
	return &AdapterManager{
		adapterFactory: factory,
		dataSourceRepo: repositories.NewDataSourceRepository(db),
		adapterLogRepo: repositories.NewAdapterLogRepository(db),
	}
}

func (m *AdapterManager) RunAdapters(adapterName string) error {
	adapterNames := []string{"TruAdapter", "DummyAdapter"} // Add new adapters here

	if adapterName != "" {
		fmt.Printf("Running adapter: %s...\n", adapterName)
		adapterNames = []string{adapterName}
	} else {
		fmt.Println("Running all adapters...")
	}

	for _, name := range adapterNames {
		adapter := m.adapterFactory.CreateAdapter(name)
		if adapter == nil {
			fmt.Printf("ERROR: No adapter found for %s\n", name)
			continue
		}

		_, err := adapter.FetchData()
		if err != nil {
			fmt.Printf("Error fetching ratings for %s: %v\n", name, err)
			continue
		}

		m.logAdapterRun(name)
		fmt.Printf("Successfully loaded analyst ratings from %s.\n", name)
	}

	return nil
}

func (m *AdapterManager) logAdapterRun(name string) {
	log := &models.AdapterLog{
		AdapterName: name,
		RunAt:       time.Now(),
	}
	m.adapterLogRepo.Create(log)
}
