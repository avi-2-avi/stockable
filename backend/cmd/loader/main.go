package main

import (
	"backend/config"
	"backend/internal/database"
	"backend/internal/manager"
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

	adapterManager := manager.NewAdapterManager(db, config)
	err = adapterManager.RunAdapters(*adapterName)
	if err != nil {
		fmt.Println("Error running adapters:", err)
	}
}
