package main

import (
	"data-loader/config"
	"data-loader/internal/adapters"
	"data-loader/internal/database"
	"data-loader/internal/models"
	"data-loader/internal/repositories"
	"data-loader/internal/services"
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

	dataSourceRepo := repositories.NewDataSourceRepository(db)
	dataSource := models.DataSource{Name: "TruAdapter"}

	err = dataSourceRepo.Create(&dataSource)
	if err != nil {
		fmt.Println("Failed to register data source:", err)
		return
	}

	analystRatingsRepo := repositories.NewAnalystRatingsRepository(db)
	analystRatingsService := services.NewAnalystRatingsService(analystRatingsRepo)

	adapter := adapters.NewTruAdapter(config.TruAdapterURL, config.TruAdapterToken, analystRatingsService, dataSource.ID)

	fmt.Println("Fetching analyst ratings from TruAdapter...")
	_, err = adapter.FetchData()
	if err != nil {
		fmt.Printf("Error fetching ratings: %v\n", err)
	} else {
		fmt.Println("Successfully loaded analyst ratings into the database.")
	}
}
