package server_tests

import (
	"context"
	
	"github.com/DATA-DOG/go-sqlmock"

	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"github.com/Azat201003/summorist-users/internal/tokens"
)

func (s *serverSuite) TestRefreshTokenOk() {
	s.dbmock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"\."refresh_token" = \$1 AND "users"\."user_id" = \$2 ORDER BY "users"\."user_id" LIMIT \$3`).
		WithArgs("SomeRefreshToken", 1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"user_id","username","is_admin"}).AddRow("1", "Some username", "true"))
	
	user := &pb.User{
		UserId: 1,
		RefreshToken: "SomeRefreshToken",
	}



	response, err := (*s.usersClient).RefreshTokens(context.Background(), &pb.RefreshRequest{
		UserId:     	user.UserId,
		RefreshToken: user.RefreshToken,
	})



	s.NoError(err)
	s.Equal(response.Code, int32(0))

	id, err := tokens.ValidateToken(response.JwtToken)
	s.NoError(err)
	s.Equal(uint64(1), id, "Another id")
	s.NoError(s.dbmock.ExpectationsWereMet())
}
