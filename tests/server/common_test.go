package server_tests

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"testing"


	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"


	pb "github.com/Azat201003/summorist-shared/gen/go/users"

	"github.com/Azat201003/summorist-users/internal/database"
	"github.com/Azat201003/summorist-users/internal/server"
)

type serverSuite struct {
	suite.Suite
	usersClient *pb.UsersClient
	dbc         *database.DBController
	dbmock      sqlmock.Sqlmock
	lis 				net.Listener
	db 					*sql.DB
}

func (s *serverSuite) SetupSuite() {
	// mock db
	fmt.Println("mock db setting up")
	db, mock, err := sqlmock.New()
	s.db = db
	s.NoError(err)
	s.dbmock = mock

	// database
	fmt.Println("database setting up")
	dialector := postgres.New(postgres.Config{
		Conn: db,
		DriverName: "postgres",
	})
	gormdb, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	s.NoError(err)
	s.dbc = &database.DBController{DB: gormdb}
	
	// service server
	fmt.Println("service server setting up")
	s.lis, _ = net.Listen("tcp", fmt.Sprintf("%v:%v", "0.0.0.0", os.Getenv("USERS_PORT")))
	grpcServer := grpc.NewServer()
	pb.RegisterUsersServer(grpcServer, &server.UserServer{DBC: s.dbc})
	go grpcServer.Serve(s.lis)
	
	// service client
	fmt.Println("service client setting up")
	conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", "0.0.0.0", os.Getenv("USERS_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))

	s.NoError(err)

	client := pb.NewUsersClient(conn)
	s.usersClient = &client
	fmt.Println("all was set up")
}

func (s *serverSuite) TearDownSuite() {
	s.lis.Close()
	s.db.Close()
}

func TestServer(t *testing.T) {
	suite.Run(t, new(serverSuite))
}
