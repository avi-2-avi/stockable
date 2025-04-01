package test

import (
	"backend/internal/database"
	"backend/internal/models"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestMigrations(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")
	db, dbErr := database.Connect()
	db.Config.Logger = logger.Default.LogMode(logger.Silent)

	t.Cleanup(func() {
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
	})

	migrationErr := database.Migrate(db)
	hasRole := db.Migrator().HasTable(&models.Role{})
	hasUser := db.Migrator().HasTable(&models.User{})
	hasCompany := db.Migrator().HasTable(&models.Company{})
	hasDataSource := db.Migrator().HasTable(&models.DataSource{})
	hasPortafolio := db.Migrator().HasTable(&models.Portafolio{})
	hasPortafolioHolding := db.Migrator().HasTable(&models.PortafolioHolding{})
	hasAnalystRating := db.Migrator().HasTable(&models.AnalystRating{})
	hasAdapterLog := db.Migrator().HasTable(&models.AdapterLog{})
	hasAdminRole := db.Where(&models.Role{Name: "admin"}).First(&models.Role{}).Error == nil
	hasUserRole := db.Where(&models.Role{Name: "user"}).First(&models.Role{}).Error == nil
	hasAdminUser := db.Where(&models.User{RoleID: 1}).First(&models.User{}).Error == nil

	assert.NoError(t, dbErr, "Database connection should not return an error")
	assert.NotNil(t, db, "Database connection should not be nil")
	assert.NoError(t, migrationErr, "Database migration should not return an error")
	assert.True(t, hasRole, "role table should exist")
	assert.True(t, hasCompany, "company table should exist")
	assert.True(t, hasPortafolio, "portafolio table should exist")
	assert.True(t, hasPortafolioHolding, "portafolio_holding table should exist")
	assert.True(t, hasAnalystRating, "rating_history table should exist")
	assert.True(t, hasDataSource, "data_source table should exist")
	assert.True(t, hasAdapterLog, "adapter_log table should exist")
	assert.True(t, hasUser, "user table should exist")
	assert.True(t, hasAdminRole, "admin role should exist")
	assert.True(t, hasUserRole, "user role should exist")
	assert.True(t, hasAdminUser, "admin user should exist")

}
