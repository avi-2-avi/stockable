package manager

import (
	"backend/internal/repositories"
	"fmt"

	"gorm.io/gorm"
)

type AdapterManager struct {
	adapterFactory *AdapterFactory
	dataSourceRepo repositories.DataSourceRepository
	adapterLogger  *AdapterLogger
}

func NewAdapterManager(factory *AdapterFactory, db *gorm.DB) *AdapterManager {
	return &AdapterManager{
		adapterFactory: factory,
		dataSourceRepo: *repositories.NewDataSourceRepository(db),
		adapterLogger:  NewAdapterLogger(repositories.NewAdapterLogRepository(db)),
	}
}

func (m *AdapterManager) RunAdapters(adapterName string) error {
	adapterNames, err := m.dataSourceRepo.GetAllAdapterNames()
	fmt.Println("Adapter names:", adapterNames)
	if err != nil {
		fmt.Println("Failed to get adapter names")
		return err
	}

	if adapterName != "" {
		fmt.Printf("Running adapter: %s...\n", adapterName)
		adapterNames = []string{adapterName}
	} else {
		fmt.Println("Running all adapters...")
	}

	for _, name := range adapterNames {
		adapter, err := m.adapterFactory.CreateAdapter(name)
		fmt.Printf("Creating adapter: %s\n", name)
		if adapter == nil || err != nil {
			fmt.Printf("ERROR: No adapter found for %s\n", name)
			continue
		}

		_, err = adapter.FetchData()
		if err != nil {
			fmt.Printf("Error fetching ratings for %s: %v\n", name, err)
			continue
		}

		m.adapterLogger.LogRun(name)
		fmt.Printf("Successfully loaded analyst ratings from %s.\n", name)
	}

	return nil
}
