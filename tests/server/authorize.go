package server_tests

import (
	"context"

	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"github.com/Azat201003/summorist-users/internal/tokens"
)

func (s *serverSuite) TestAuthorizeOk() {
	token, err := tokens.GenerateToken(27)

	response, err := (*s.usersClient).Authorize(context.Background(), &pb.AuthRequest{
		JwtToken: token,
	})

	s.NoError(err)
	s.Equal(response.Code, int32(0))
	s.Equal(response.UserId, uint64(27))
}
