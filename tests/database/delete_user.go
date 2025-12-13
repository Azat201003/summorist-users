package database_tests

import (
	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"github.com/Azat201003/summorist-users/internal/passwords"
	"github.com/Azat201003/summorist-users/internal/tokens"
)

func (s *databaseSuite) TestDeleteUserOk() {
	// Create a user first
	user := &pb.User{
		Username:     "test-" + generateRandomString(10),
		PasswordHash: passwords.Hash(generateRandomString(16)),
		RefreshToken: tokens.GenerateRefreshToken(),
	}
	userId, err := s.dbc.CreateUser(user)
	s.NoError(err)

	// Delete the user
	err = s.dbc.DeleteUser(userId)
	s.NoError(err)

	// Try to find the user and verify it's deleted
	foundUsers, err := s.dbc.FindUsers(&pb.User{Id: userId})
	s.NoError(err)
	s.Len(foundUsers, 0)
}
