package main

import (
	"data-loader/config"
	"data-loader/internal/database"
	"fmt"
)

func main() {
	fmt.Println("Starting data loader...")

	_, err := config.LoadConfig()
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

	// Load data from API adapters
}
