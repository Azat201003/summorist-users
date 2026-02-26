package server_tests

import (
	"context"

	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	//"github.com/Azat201003/summorist-users/internal/passwords"
	"github.com/Azat201003/summorist-users/internal/server"
	"github.com/Azat201003/summorist-users/internal/tokens"
	"github.com/DATA-DOG/go-sqlmock"
)

func (s *serverSuite) TestRemoveUserMyselfOk() {
	username := "test-" + generateRandomString(10)
	userId := uint64(1)

	jwtToken, err := tokens.GenerateToken(userId)
	s.NoError(err)

	s.dbmock.ExpectQuery(`SELECT \* FROM "users"`).
		WithArgs(userId, 1).
		WillReturnRows(sqlmock.NewRows([]string{"is_admin", "password_hash", "refresh_token", "username", "user_id"}).AddRow(false, []byte("..."), "...", username, userId))
	s.dbmock.ExpectBegin()
	s.dbmock.ExpectExec(`DELETE FROM "users" WHERE user_id = \$1`).
		WithArgs(userId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.dbmock.ExpectCommit()


	removeResponse, err := (*s.usersClient).RemoveUser(context.Background(), &pb.RemoveRequest{
		JwtToken: jwtToken,
		UserId:   userId,
	})
	

	s.NoError(err)
	s.Equal(removeResponse.Code, int32(0))
	s.NoError(s.dbmock.ExpectationsWereMet())
}

// TODO add admin test ok

func (s *serverSuite) TestRemoveUserInvalidToken() {
	_, err := (*s.usersClient).RemoveUser(context.Background(), &pb.RemoveRequest{
		JwtToken: "invalid",
		UserId:   1,
	})
	

	s.Error(err)
}

func (s *serverSuite) TestRemoveUserPermissionDeniedUndefinedUserId() {
	token, err := tokens.GenerateToken(1)
	s.NoError(err)
	definedUserId := uint64(1)
	undefinedUserId := uint64(999999999999999)
	
	s.dbmock.ExpectQuery(`SELECT \* FROM "users"`).
		WithArgs(definedUserId, 1).
		WillReturnRows(sqlmock.NewRows([]string{"is_admin", "password_hash", "refresh_token", "username", "user_id"}).AddRow(false, []byte("..."), "...", "...", definedUserId))


	_, err = (*s.usersClient).RemoveUser(context.Background(), &pb.RemoveRequest{
		JwtToken: token,
		UserId:   undefinedUserId,
	})
	

	s.Error(err, server.ErrNotPermitted)
	s.NoError(s.dbmock.ExpectationsWereMet())
}
