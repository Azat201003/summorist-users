package server_tests

import (
	"github.com/stretchr/testify/suite"
	"testing"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Azat201003/summorist-shared/gen/go/user-service"

	"github.com/Azat201003/summorist-users/internal/config"
	"github.com/Azat201003/summorist-users/internal/database"
)

type serverSuite struct {
	suite.Suite
	usersClient *pb.UsersClient
	dbc         *database.DBController
}

func (s *serverSuite) SetupTest() {
	// service client
	conf := config.GetConfig()
	conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", conf.Host, conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	s.NoError(err)

	client := pb.NewUsersClient(conn)
	s.usersClient = &client

	// database
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		conf.DBHost,
		conf.DBUser,
		conf.DBPassword,
		conf.DBName,
		conf.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	s.NoError(err)
	s.dbc = &database.DBController{
		DB: db,
	}
}

func TestServer(t *testing.T) {
	suite.Run(t, new(serverSuite))
}
