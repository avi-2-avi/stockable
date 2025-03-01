package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("No .env file found")
	}

	DatabaseURL := os.Getenv("DATABASE_URL")

	return &Config{DatabaseURL: DatabaseURL}, nil
}
