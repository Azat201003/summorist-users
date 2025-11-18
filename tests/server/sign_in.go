package server_tests

import (
	"context"

	"github.com/Azat201003/summorist-shared/gen/go/common"
	pb "github.com/Azat201003/summorist-shared/gen/go/user-service"
	"github.com/Azat201003/summorist-users/internal/passwords"
	"github.com/Azat201003/summorist-users/internal/tokens"
)

func (s *serverSuite) TestSignInOk() {
	response, err := (*s.usersClient).SignIn(context.Background(), &pb.SignInRequest{
		PasswordHash: passwords.Hash("1234"),
		Username:     "Abeme",
	})
	s.NoError(err)

	users, err := s.dbc.FindUsers(&common.User{
		Username: "Abeme",
	})
	if err == nil && len(users) == 1 {
		s.Equal(users[0].RefreshToken, response.RefreshToken)
	}

	id, err := tokens.ValidateToken(response.JwtToken)
	s.NoError(err)

	users, err = s.dbc.FindUsers(
		&common.User{Id: id},
	)
	s.NoError(err)
	s.Equal(len(users), 1)
	s.Equal(users[0].Username, "Abeme")
	s.Equal(users[0].PasswordHash, passwords.Hash("1234"))
}
