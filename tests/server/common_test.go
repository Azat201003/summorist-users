package server_tests

import (
	"testing"
	"github.com/stretchr/testify/suite"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Azat201003/summorist-shared/gen/go/user-service"
)

type serverSuite struct {
	suite.Suite
	usersClient *pb.UsersClient
}

func (s *serverSuite) SetupTest() {
	conn, err := grpc.NewClient("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	s.NoError(err)

	client := pb.NewUsersClient(conn)
	s.usersClient = &client
}

func TestServer(t *testing.T) {
	suite.Run(t, new(serverSuite))
}
