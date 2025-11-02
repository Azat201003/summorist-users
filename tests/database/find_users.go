package database_tests

import (
	"github.com/Azat201003/summorist-shared/gen/go/common"
	"github.com/Azat201003/summorist-users/internal/passwords"
)

func (s *databaseSuite) TestFindUsersOk() {
	users, err := s.dbc.FindUsers(&common.User{
		Username: "Abeme",
	})
	s.NoError(err)

	s.True(passwords.Verify(users[0].PasswordHash, "1234"))
}
