package main

import (
	"backend/config"
	"backend/internal/database"
	"backend/internal/routes"
	"fmt"
	"log"
)

func main() {
	config.LoadConfig()

	db, err := database.Connect()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return
	}

	router := routes.SetupRouter()
	routes.RegisterRoutes(router, db)

	log.Println("Server is running on port 8085")
	if err := router.Run("0.0.0.0:8085"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
