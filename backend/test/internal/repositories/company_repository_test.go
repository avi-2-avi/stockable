package test

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/repositories"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestCompanyRepository(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")
	db, dbErr := database.Connect()
	db.Config.Logger = logger.Default.LogMode(logger.Silent)
	migrationErr := database.Migrate(db)

	companyRepo := repositories.NewCompanyRepository(db)
	company := models.Company{
		Name: "Company-TEST",
	}

	companyRepoErr := companyRepo.Create(&company)

	t.Cleanup(func() {
		companyRepo.Delete(company.ID)

		tables, err := db.Migrator().GetTables()
		if err == nil {
			for _, table := range tables {
				_ = db.Migrator().DropTable(table)
			}
		}

		for _, env := range os.Environ() {
			key := env[:strings.Index(env, "=")]
			os.Unsetenv(key)
		}

		os.Remove(".env")
	})

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.NoError(t, migrationErr, "Database migration should not return an error")
	assert.NoError(t, companyRepoErr, "Should create company without error")
}
