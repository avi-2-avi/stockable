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

func TestPortafolioRepository(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")
	db, dbErr := database.Connect()
	db.Config.Logger = logger.Default.LogMode(logger.Silent)
	migrationErr := database.Migrate(db)

	authRepo := repositories.NewAuthRepository(db)

	role := models.Role{Name: "API-TEST"}
	authRepo.CreateRole(&role)
	user := models.User{
		FullName: "API-TEST",
		Email:    "api-test@stockable.com",
		Password: "api-test",
		RoleID:   role.ID,
	}
	authRepo.CreateUser(&user)

	dataSourceRepo := repositories.NewDataSourceRepository(db)
	dataSource := models.DataSource{
		Name: "API-TEST",
	}
	dataSourceRepo.Create(&dataSource)

	portafolioRepo := repositories.NewPortafolioRepository(db)
	portafolio := models.Portafolio{
		Name:         "API-TEST",
		Category:     "simulation",
		UserID:       user.ID,
		DataSourceID: dataSource.ID,
	}
	repoErr := portafolioRepo.Create(&portafolio)
	fetchedPortafolio, fetchErr := portafolioRepo.GetByID(portafolio.ID)

	t.Cleanup(func() {
		portafolioRepo.Delete(portafolio.ID)
		dataSourceRepo.Delete(dataSource.ID)
		authRepo.DeleteUser(user.ID)
		authRepo.DeleteRole(role.ID)

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
	assert.NoError(t, repoErr, "Should create portafolio without error")
	assert.NotZero(t, portafolio.ID, "Portafolio ID should be set")
	assert.NoError(t, fetchErr, "Should fetch portafolio without error")
	assert.Equal(t, "API-TEST", fetchedPortafolio.Name, "Fetched portafolio should match")
}
