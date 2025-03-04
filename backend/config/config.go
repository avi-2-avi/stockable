package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL     string
	AllowedOrigin   string
	TruAdapterURL   string
	TruAdapterToken string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("no .env file found")
	}

	DatabaseURL := os.Getenv("DATABASE_URL")
	AllowedOrigin := os.Getenv("ALLOWED_ORIGIN")

	// Adapters
	TruAdapterURL := os.Getenv("TRU_ADAPTER_URL")
	TruAdapterToken := os.Getenv("TRU_ADAPTER_TOKEN")

	return &Config{
		DatabaseURL:     DatabaseURL,
		AllowedOrigin:   AllowedOrigin,
		TruAdapterURL:   TruAdapterURL,
		TruAdapterToken: TruAdapterToken,
	}, nil
}
