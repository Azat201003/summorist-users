package server_tests

import (
	"context"

	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"github.com/Azat201003/summorist-users/internal/passwords"
	"github.com/Azat201003/summorist-users/internal/tokens"
)

func (s *serverSuite) TestSignInOk() {
	response, err := (*s.usersClient).SignIn(context.Background(), &pb.SignInRequest{
		PasswordHash: passwords.Hash("1234"),
		Username:     "Abeme",
	})
	s.NoError(err)

	users, err := s.dbc.FindUsers(&pb.User{
		Username: "Abeme",
	})
	if err == nil && len(users) == 1 {
		s.Equal(users[0].RefreshToken, response.RefreshToken)
	}

	id, err := tokens.ValidateToken(response.JwtToken)
	s.NoError(err)

	users, err = s.dbc.FindUsers(
		&pb.User{UserId: id},
	)
	s.NoError(err)
	s.Equal(len(users), 1)
	s.Equal(users[0].Username, "Abeme")
	s.Equal(users[0].PasswordHash, passwords.Hash("1234"))
}
