package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"data-loader/config"
)

func TestLoadConfigWithEnvFile(t *testing.T) {
	mockEnvValue := "postgres://test_user:test_password@localhost:26257/testdb?sslmode=disable"
	mockEnvContent := "DATABASE_URL=" + mockEnvValue
	fileErr := os.WriteFile(".env", []byte(mockEnvContent), 0644)

	loaded_config, configErr := config.LoadConfig()

	assert.NoError(t, fileErr, "Failed to create mock .env file")
	assert.NoError(t, configErr, "LoadConfig should not return an error")
	assert.Equal(t, mockEnvValue, loaded_config.DatabaseURL,
		"DATABASE_URL should be correctly loaded from .env")
	os.Remove(".env")
	os.Unsetenv("DATABASE_URL")
}

func TestLoadConfigWithoutEnvFile(t *testing.T) {
	os.Remove(".env")

	_, err := config.LoadConfig()

	assert.Errorf(t, err, "LoadConfig should return an error when .env is missing")
}
