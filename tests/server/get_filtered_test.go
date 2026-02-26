package server_tests

import (
	"context"
	//"regexp"
	"io"

	"github.com/DATA-DOG/go-sqlmock"

	pb "github.com/Azat201003/summorist-shared/gen/go/users"
)

func (s *serverSuite) TestGetFilteredOk() {
	s.dbmock.ExpectQuery(`SELECT (.+) FROM "users" WHERE "users"\."username" = \$1`).
		WithArgs("Some username").
		WillReturnRows(sqlmock.NewRows([]string{"user_id","username","is_admin"}).AddRow("1", "Some username", "true"))


	stream, err := (*s.usersClient).GetFiltered(context.Background(), &pb.GetFilteredRequest{
		Filter: &pb.User{
			Username: "Some username",
		},
	})
	user, err := stream.Recv()
	

	s.NoError(err)
	s.Equal("Some username", user.Username)
	_, err = stream.Recv()
	s.Error(err, io.EOF)
	s.NoError(s.dbmock.ExpectationsWereMet())
}

func (s *serverSuite) TestGetFilteredNotFound() {
	s.dbmock.ExpectQuery(`.+`).   // matches any query
    WithArgs(sqlmock.AnyArg()).
    WillReturnRows(sqlmock.NewRows([]string{"user_id","username","is_admin"}))


	stream, err := (*s.usersClient).GetFiltered(context.Background(), &pb.GetFilteredRequest{
		Filter: &pb.User{
			Username: "Some username",
		},
	})
	s.NoError(err)
	_, err = stream.Recv()


	s.NoError(s.dbmock.ExpectationsWereMet())
	s.Error(err, io.EOF)
}

