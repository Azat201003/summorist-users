package server_tests

import (
	"context"

	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"github.com/Azat201003/summorist-users/internal/passwords"
	"github.com/Azat201003/summorist-users/internal/tokens"
)

func (s *serverSuite) TestRemoveUserOk() {
	username := "test-" + generateRandomString(10)
	password := generateRandomString(16)

	// Create user directly using dbc
	userId, err := s.dbc.CreateUser(&pb.User{
		Username:     username,
		PasswordHash: passwords.Hash(password),
		RefreshToken: tokens.GenerateRefreshToken(),
	})
	s.NoError(err)

	// Generate JWT token using tokens lib
	jwtToken, err := tokens.GenerateToken(userId)
	s.NoError(err)

	// Remove user
	removeResponse, err := (*s.usersClient).RemoveUser(context.Background(), &pb.RemoveRequest{
		JwtToken: jwtToken,
		UserId:   userId,
	})
	s.NoError(err)
	s.Equal(removeResponse.Code, int32(0))

	// Verify removal
	deletedUsers, err := s.dbc.FindUsers(&pb.User{Id: userId})
	s.NoError(err)
	s.Equal(len(deletedUsers), 0)
}

func (s *serverSuite) TestRemoveUserInvalidToken() {
	_, err := (*s.usersClient).RemoveUser(context.Background(), &pb.RemoveRequest{
		JwtToken: "invalid",
		UserId:   1,
	})
	s.Error(err)
}

func (s *serverSuite) TestRemoveUserPermissionDenied() {
	// Use existing user token but try to remove different user
	token, err := tokens.GenerateToken(1)
	s.NoError(err)

	removeResponse, err := (*s.usersClient).RemoveUser(context.Background(), &pb.RemoveRequest{
		JwtToken: token,
		UserId:   999,
	})
	s.NoError(err)
	s.Equal(removeResponse.Code, int32(2))
}
