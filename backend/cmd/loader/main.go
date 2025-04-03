package main

import (
	"backend/config"
	"backend/internal/adapters"
	"backend/internal/database"
	"backend/internal/manager"
	"backend/internal/services"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Starting data loader...")

	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load configuration")
		return
	}

	db, err := database.Connect()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return
	}

	err = database.Migrate(db)
	if err != nil {
		fmt.Println("Failed to migrate database")
		return
	}

	adapterName := flag.String("name", "", "Specify adapter name to run individually")
	flag.Parse()

	adapterFactory := manager.NewAdapterFactory(config, db)

	dummySource := adapterFactory.CreateDataSource("DummyAdapter", false)

	adapterFactory.RegisterAdapter("DummyAdapter", func(factory *manager.AdapterFactory) adapters.RatingAdapter {
		analystService := services.NewAnalystRatingService(factory.GetAnalystRatingRepository())
		companyService := services.NewCompanyService(factory.GetCompanyRepository())
		return adapters.NewDummyAdapter(analystService, companyService, dummySource.ID)
	})

	truSource := adapterFactory.CreateDataSource("TruAdapter", true)

	adapterFactory.RegisterAdapter("TruAdapter", func(factory *manager.AdapterFactory) adapters.RatingAdapter {
		analystService := services.NewAnalystRatingService(factory.GetAnalystRatingRepository())
		companyService := services.NewCompanyService(factory.GetCompanyRepository())
		return adapters.NewTruAdapter(
			config.TruAdapterURL,
			config.TruAdapterToken,
			analystService,
			companyService,
			truSource.ID,
		)
	})

	adapterManager := manager.NewAdapterManager(adapterFactory, db)

	err = adapterManager.RunAdapters(*adapterName)

	if err != nil {
		fmt.Println("Error running adapters:", err)
	}
}
