package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"backend/internal/database"
)

func TestDatabaseConnection(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")

	db, dbErr := database.Connect()
	sqlDB, sqlErr := db.DB()

	assert.NoError(t, dbErr, "Database should not return an error")
	assert.NotNil(t, db, "Database should not be nil")
	assert.NoError(t, sqlErr, "Getting underlying SQL Database should not return an error")
	assert.NotNil(t, sqlDB, "Underlying SQL Database should not be nil")

	sqlDB.Close()
	os.Unsetenv("DATABASE_URL")
}
