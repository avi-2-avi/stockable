package test

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/repositories"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestPortafolioHoldingRepository(t *testing.T) {
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
		Name: "DataSource-TEST",
	}
	dataSourceRepoErr := dataSourceRepo.Create(&dataSource)

	portafolioRepo := repositories.NewPortafolioRepository(db)
	portafolio := models.Portafolio{
		Name:         "Portafolio-TEST",
		Category:     "simulation",
		UserID:       user.ID,
		DataSourceID: dataSource.ID,
	}
	portafolioRepoErr := portafolioRepo.Create(&portafolio)

	companyRepo := repositories.NewCompanyRepository(db)
	company := models.Company{
		Name: "Company-TEST",
	}
	companyRepoErr := companyRepo.Create(&company)

	holdingRepo := repositories.NewPortafolioHoldingRepository(db)
	holding := models.PortafolioHolding{
		Quantity:      100,
		PurchasePrice: 100,
		PurchasedAt:   time.Now(),
		PortafolioID:  portafolio.ID,
		CompanyID:     company.ID,
	}
	holdingRepoErr := holdingRepo.Create(&holding)

	t.Cleanup(func() {
		holdingRepo.Delete(holding.ID)
		companyRepo.Delete(company.ID)
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
	assert.NoError(t, dataSourceRepoErr, "Should create data source without error")
	assert.NoError(t, portafolioRepoErr, "Should create portafolio without error")
	assert.NoError(t, companyRepoErr, "Should create company without error")
	assert.NoError(t, holdingRepoErr, "Should create holding without error")
}
