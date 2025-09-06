package tests

import (
    "testing"
    "github.com/stretchr/testify/suite"
//    "github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/Azat201003/summorist-shared/gen/go/user-service"
)

type BaseSuite struct {
	suite.Suite
	usersClient *pb.UsersClient
}

func (s *BaseSuite) SetupTest() {	
	conn, err := grpc.NewClient("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	
	if err != nil {
		s.T().Fatalf("%v%v", "Not connected to client ", err)
	}

	defer conn.Close()
	client := pb.NewUsersClient(conn)
	s.usersClient = &client
}

func TestSuit(t *testing.T) {
	suite.Run(t, new(BaseSuite))
}

