package server_tests

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"github.com/Azat201003/summorist-users/internal/passwords"
	"github.com/DATA-DOG/go-sqlmock"
)

func (s *serverSuite) TestRegisterUserOk() {
	password := generateRandomString(16)

	s.dbmock.ExpectBegin()
	s.dbmock.ExpectQuery(`INSERT INTO "users"`).
		WithArgs("test", passwords.Hash(password), sqlmock.AnyArg(), false).
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).
			AddRow(1))
	s.dbmock.ExpectCommit()


	response, err := (*s.usersClient).SignUp(context.Background(), &pb.User{
		Username:     "test",
		PasswordHash: passwords.Hash(password),
	})


	s.NoError(err)
	s.Equal(response.Code, int32(0))
	s.NoError(s.dbmock.ExpectationsWereMet())
}

func generateRandomString(n int) string {
	bytes := make([]byte, (n+1)/2)
	if _, err := rand.Read(bytes); err != nil {
		return "error"
	}
	return hex.EncodeToString(bytes)[:n]
}
