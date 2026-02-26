package server_tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Azat201003/summorist-shared/gen/go/users"

	"github.com/Azat201003/summorist-users/internal/database"
)

type serverSuite struct {
	suite.Suite
	usersClient *pb.UsersClient
	dbc         *database.DBController
}

func (s *serverSuite) SetupTest() {
	// service client
	conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", os.Getenv("USERS_HOST"), os.Getenv("USERS_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))

	s.NoError(err)

	client := pb.NewUsersClient(conn)
	s.usersClient = &client

	// database
	s.dbc = &database.DBController{}
	err = s.dbc.InitDB()
	s.NoError(err)
}

func TestServer(t *testing.T) {
	suite.Run(t, new(serverSuite))
}
