package database_tests

import (
	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"github.com/Azat201003/summorist-users/internal/passwords"
)

func (s *databaseSuite) TestFindUsersOk() {
	users, err := s.dbc.FindUsers(&pb.User{
		Username: "Abeme",
	})
	s.NoError(err)

	s.True(passwords.Verify(users[0].PasswordHash, "1234"))
}
