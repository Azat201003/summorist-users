package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	pb "github.com/Azat201003/summorist-shared/gen/go/user-service"
	"google.golang.org/grpc"

	"github.com/Azat201003/summorist-shared/gen/go/common"
	"github.com/Azat201003/summorist-users/internal/database"
	"github.com/Azat201003/summorist-users/internal/tokens"
)

type userServer struct {
	pb.UnimplementedUsersServer
	dbc *database.DBController
}

func (s *userServer) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	users, err := s.dbc.FindUsers(&common.User{
		Username:     request.Username,
		PasswordHash: request.PasswordHash,
	})
	if err != nil {
		return &pb.SignInResponse{Code: 1}, err
	}
	if len(users) != 1 {
		return &pb.SignInResponse{Code: 2}, errors.New("No records found")
	}
	jwtToken, err := tokens.GenerateToken(users[0].Id, "./")
//	key, err1 := tokens.GetPublicKey()
	if err != nil {
		return &pb.SignInResponse{Code: 3}, err
	}
	return &pb.SignInResponse{JwtToken: jwtToken, Code: 0}, nil
}

func (s *userServer) Authorize(ctx context.Context, request *pb.AuthRequest) (*pb.AuthResponse, error) {
	userId, err := tokens.ValidateToken(request.JwtToken, "./")

	if (err != nil) {
		return &pb.AuthResponse{Code: 1}, err
	}

	return &pb.AuthResponse{UserId: userId, Code: 0}, nil
}

func newServer() *userServer {
	dsn := "host=localhost user=smrt_users password=1234 dbname=smrt_users port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("%v", err)
	}

	return &userServer{dbc: &database.DBController{
		DB: db,
	}}
}

func StartServer(host string, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUsersServer(grpcServer, newServer())
	log.Printf("Server starting on port %v", port)
	grpcServer.Serve(lis)
}
