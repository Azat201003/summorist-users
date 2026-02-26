package server_tests

import (
	"context"

	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"github.com/Azat201003/summorist-users/internal/passwords"
	"github.com/Azat201003/summorist-users/internal/tokens"
	"github.com/DATA-DOG/go-sqlmock"
)

func (s *serverSuite) TestSignInOk() {
	username := "Abeme"
	passwordHash := passwords.Hash("1234")
	userId := uint64(1)

	s.dbmock.ExpectQuery(`SELECT \* FROM "users"`).
		WithArgs(username, passwordHash, 1).
		WillReturnRows(sqlmock.NewRows([]string{"is_admin", "password_hash", "refresh_token", "username", "user_id"}).AddRow(false, passwordHash, "...", username, userId))
	

	response, err := (*s.usersClient).SignIn(context.Background(), &pb.SignInRequest{
		PasswordHash: passwordHash,
		Username:     username,
	})
	

	s.NoError(err)

	tokenUserId, err := tokens.ValidateToken(response.JwtToken)
	s.NoError(err)

	s.Equal(userId, tokenUserId)
}
