package database_tests

import (
	"fmt"
	"github.com/Azat201003/summorist-users/internal/config"
	"github.com/Azat201003/summorist-users/internal/database"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

type databaseSuite struct {
	suite.Suite
	dbc *database.DBController
}

func (s *databaseSuite) SetupTest() {
	conf := config.GetConfig()
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		conf.DBHost,
		conf.DBUser,
		conf.DBPassword,
		conf.DBName,
		conf.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	s.NoError(err)
	s.dbc = &database.DBController{
		DB: db,
	}
}

func TestDatabase(t *testing.T) {
	suite.Run(t, new(databaseSuite))
}
