package server_tests

import (
	"context"

	pb "github.com/Azat201003/summorist-shared/gen/go/user-service"
	"github.com/Azat201003/summorist-shared/gen/go/common"
	"github.com/Azat201003/summorist-users/internal/tokens"
)

func (s *serverSuite) TestRefreshTokenOk() {
	users, err := s.dbc.FindUsers(&common.User{Username: "Abeme"})
	user := users[0]

	response, err := (*s.usersClient).RefreshTokens(context.Background(), &pb.RefreshRequest{
		Username: "Abeme",
		RefreshToken: user.RefreshToken,
	})

	s.NoError(err)
	s.Equal(response.Code, int32(0))

	id, err := tokens.ValidateToken(response.JwtToken, "../../")
	s.NoError(err)
	s.Equal(id, uint64(27))

	users, err = s.dbc.FindUsers(&common.User{Username: "Abeme"})
	s.NotEqual(response.RefreshToken, "")
	s.Equal(users[0].RefreshToken, response.RefreshToken)
}
