package server_tests

import (
	"github.com/stretchr/testify/suite"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Azat201003/summorist-shared/gen/go/user-service"

	"github.com/Azat201003/summorist-users/internal/database"
)

type serverSuite struct {
	suite.Suite
	usersClient *pb.UsersClient
	dbc         *database.DBController
}

func (s *serverSuite) SetupTest() {
	// service client
	conn, err := grpc.NewClient("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	s.NoError(err)

	client := pb.NewUsersClient(conn)
	s.usersClient = &client

	// database
	dsn := "host=localhost user=smrt_users password=1234 dbname=smrt_users port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	s.NoError(err)
	s.dbc = &database.DBController{
		DB: db,
	}
}

func TestServer(t *testing.T) {
	suite.Run(t, new(serverSuite))
}
