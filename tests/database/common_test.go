package database_tests

import (
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
	dsn := "host=localhost user=smrt_users password=1234 dbname=smrt_users port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	s.NoError(err)
	s.dbc = &database.DBController{
		DB: db,
	}
}

func TestDatabase(t *testing.T) {
	suite.Run(t, new(databaseSuite))
}
