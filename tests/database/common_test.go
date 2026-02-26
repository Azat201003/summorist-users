package database_tests

import (
	"github.com/Azat201003/summorist-users/internal/database"
	"github.com/stretchr/testify/suite"
	"testing"
)

type databaseSuite struct {
	suite.Suite
	dbc *database.DBController
}

func (s *databaseSuite) SetupTest() {
	s.dbc = &database.DBController{}
	err := s.dbc.InitDB()
	s.NoError(err)
}

func TestDatabase(t *testing.T) {
	suite.Run(t, new(databaseSuite))
}
