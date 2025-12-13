package server_tests

import (
	"context"

	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"github.com/Azat201003/summorist-users/internal/passwords"
	"github.com/Azat201003/summorist-users/internal/tokens"
)

func (s *serverSuite) TestUpdateUserOk() {
	username := "test-" + generateRandomString(10)
	newUsername := "updated-" + generateRandomString(10)
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

	// Update user
	updateResponse, err := (*s.usersClient).UpdateUser(context.Background(), &pb.UpdateRequest{
		JwtToken: jwtToken,
		User: &pb.User{
			Id:           userId,
			Username:     newUsername,
		},
	})
	s.NoError(err)
	s.Equal(updateResponse.Code, int32(0))

	// Verify update
	updatedUsers, err := s.dbc.FindUsers(&pb.User{Username: newUsername})
	s.NoError(err)
	s.Equal(len(updatedUsers), 1)
	s.Equal(updatedUsers[0].Username, newUsername)
}

func (s *serverSuite) TestUpdateUserInvalidToken() {
	_, err := (*s.usersClient).UpdateUser(context.Background(), &pb.UpdateRequest{
		JwtToken: "invalid",
		User: &pb.User{
			Id:       1,
			Username: "test",
		},
	})
	s.Error(err)
}

func (s *serverSuite) TestUpdateUserPermissionDenied() {
	// Use existing user token but try to update different user
	token, err := tokens.GenerateToken(1)
	s.NoError(err)

	updateResponse, err := (*s.usersClient).UpdateUser(context.Background(), &pb.UpdateRequest{
		JwtToken: token,
		User: &pb.User{
			Id:       999,
			Username: "test",
		},
	})
	s.NoError(err)
	s.Equal(updateResponse.Code, int32(2))
}
