package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL     string
	TruAdapterURL   string
	TruAdapterToken string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("No .env file found")
	}

	DatabaseURL := os.Getenv("DATABASE_URL")

	// Adapters
	TruAdapterURL := os.Getenv("TRU_ADAPTER_URL")
	TruAdapterToken := os.Getenv("TRU_ADAPTER_TOKEN")

	return &Config{
		DatabaseURL:     DatabaseURL,
		TruAdapterURL:   TruAdapterURL,
		TruAdapterToken: TruAdapterToken,
	}, nil
}
