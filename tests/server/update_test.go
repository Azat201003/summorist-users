package server_tests

import (
	"context"

	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"github.com/Azat201003/summorist-users/internal/tokens"
	"github.com/Azat201003/summorist-users/internal/server"
	"github.com/DATA-DOG/go-sqlmock"
)

func (s *serverSuite) TestUpdateUserMyselfOk() {
	username := "test-" + generateRandomString(10)
	userId := uint64(1)

	jwtToken, err := tokens.GenerateToken(userId)
	s.NoError(err)

	newUsername := "newUsername"
	s.dbmock.ExpectQuery(`SELECT \* FROM "users"`).
		WithArgs(userId, 1).
		WillReturnRows(sqlmock.NewRows([]string{"is_admin", "password_hash", "refresh_token", "username", "user_id"}).AddRow(false, []byte("..."), "...", username, userId))
	s.dbmock.ExpectBegin()
	s.dbmock.ExpectExec(`UPDATE`).
		WithArgs(newUsername, userId, userId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.dbmock.ExpectCommit()


	updateResponse, err := (*s.usersClient).UpdateUser(context.Background(), &pb.UpdateRequest{
		JwtToken: jwtToken,
		User: &pb.User{
			UserId: userId,
			Username: newUsername,
		},
	})

	s.NoError(err)
	s.Equal(updateResponse.Code, int32(0))
	s.NoError(s.dbmock.ExpectationsWereMet())
}

// TODO add admin test ok

func (s *serverSuite) TestUpdateUserInvalidToken() {
	_, err := (*s.usersClient).UpdateUser(context.Background(), &pb.UpdateRequest{
		JwtToken: "invalid",
		User: &pb.User{
			UserId:       1,
			Username: "test",
		},
	})
	s.Error(err)
}

func (s *serverSuite) TestUpdateUserPermissionDeniedUndefinedUserId() {
	token, err := tokens.GenerateToken(1)
	s.NoError(err)
	definedUserId := uint64(1)
	undefinedUserId := uint64(999999999999999)
	
	s.dbmock.ExpectQuery(`SELECT \* FROM "users"`).
		WithArgs(definedUserId, 1).
		WillReturnRows(sqlmock.NewRows([]string{"is_admin", "password_hash", "refresh_token", "username", "user_id"}).AddRow(false, []byte("..."), "...", "...", definedUserId))


	_, err = (*s.usersClient).UpdateUser(context.Background(), &pb.UpdateRequest{
		JwtToken: token,
		User: &pb.User{
			UserId:   undefinedUserId,
		},
	})
	

	s.Error(err, server.ErrNotPermitted)
	s.NoError(s.dbmock.ExpectationsWereMet())
}
