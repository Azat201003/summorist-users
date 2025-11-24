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
	"github.com/Azat201003/summorist-users/internal/config"
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
	jwtToken, err := tokens.GenerateToken(users[0].Id)
	//	key, err1 := tokens.GetPublicKey()
	if err != nil {
		return &pb.SignInResponse{Code: 3}, err
	}
	return &pb.SignInResponse{JwtToken: jwtToken, Code: 0, RefreshToken: users[0].RefreshToken}, nil
}

func (s *userServer) Authorize(ctx context.Context, request *pb.AuthRequest) (*pb.AuthResponse, error) {
	userId, err := tokens.ValidateToken(request.JwtToken)

	if err != nil {
		return &pb.AuthResponse{Code: 1}, err
	}

	return &pb.AuthResponse{UserId: userId, Code: 0}, nil
}

func (s *userServer) SignUp(ctx context.Context, request *common.User) (*pb.SignUpResponse, error) {
	request.RefreshToken = tokens.GenerateRefreshToken()
	_, err := s.dbc.CreateUser(request)

	if err != nil {
		return &pb.SignUpResponse{Code: 1}, err
	}

	return &pb.SignUpResponse{Code: 0}, nil
}

func (s *userServer) RefreshTokens(ctx context.Context, request *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	users, err := s.dbc.FindUsers(&common.User{
		Username:     request.Username,
		RefreshToken: request.RefreshToken,
	})

	if err != nil {
		return &pb.RefreshResponse{Code: 1}, err
	}
	if len(users) != 1 {
		return &pb.RefreshResponse{Code: 2}, errors.New("No records found")
	}

	jwtToken, err := tokens.GenerateToken(users[0].Id)

	newRefreshToken := tokens.GenerateRefreshToken()
	users[0].RefreshToken = newRefreshToken
	s.dbc.UpdateUser(&users[0])

	return &pb.RefreshResponse{
		RefreshToken: newRefreshToken,
		JwtToken:     jwtToken,
		Code:         0,
	}, nil
}

func (s *userServer) GetFiltered(request *common.User, stream pb.Users_GetFilteredServer) error {
	users, err := s.dbc.FindUsers(request)
	if err != nil {
		return err
	}

	for _, user := range users {
		if err := stream.Send(&user); err != nil {
			return err
		}
	}
	return nil
}

func newServer() *userServer {
	conf := config.GetConfig()
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		conf.DBHost,
		conf.DBUser,
		conf.DBPassword,
		conf.DBName,
		conf.DBPort,
	)
    fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("%v", err)
	}

	return &userServer{dbc: &database.DBController{
		DB: db,
	}}
}

func StartServer() {
	conf := config.GetConfig()
	lis, _ := net.Listen("tcp", fmt.Sprintf("%v:%v", conf.Host, conf.Port))
	grpcServer := grpc.NewServer()
	pb.RegisterUsersServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
