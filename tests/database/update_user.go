package database_tests

import (
	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"github.com/Azat201003/summorist-users/internal/passwords"
	"github.com/Azat201003/summorist-users/internal/tokens"
)

func (s *databaseSuite) TestUpdateUserOk() {
	// Create a user first
	user := &pb.User{
		Username:     "test-" + generateRandomString(10),
		PasswordHash: passwords.Hash(generateRandomString(16)),
		RefreshToken: tokens.GenerateRefreshToken(),
	}
	userId, err := s.dbc.CreateUser(user)
	s.NoError(err)

	// Update the user
	newUsername := "updated-" + generateRandomString(10)
	updatedUser := &pb.User{
		Id:       userId,
		Username: newUsername,
	}
	err = s.dbc.UpdateUser(updatedUser)
	s.NoError(err)

	// Find the user and verify update
	foundUsers, err := s.dbc.FindUsers(&pb.User{Id: userId})
	s.NoError(err)
	s.Len(foundUsers, 1)
	s.Equal(newUsername, foundUsers[0].Username)
	// PasswordHash and RefreshToken should remain unchanged
	s.Equal(user.PasswordHash, foundUsers[0].PasswordHash)
	s.Equal(user.RefreshToken, foundUsers[0].RefreshToken)
}
