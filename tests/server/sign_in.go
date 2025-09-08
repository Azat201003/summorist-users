package server_tests

import (
	"context"
	pb "github.com/Azat201003/summorist-shared/gen/go/user-service"
	"github.com/Azat201003/summorist-users/internal/passwords"
)

func (s *serverSuite) TestSignInOk() {
	_, err := (*s.usersClient).SignIn(context.Background(), &pb.SignInRequest{
		PasswordHash: passwords.Hash("1234"),
		Username:     "abeme",
	})
	s.NoError(err)
}
