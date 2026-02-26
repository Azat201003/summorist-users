package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"github.com/Azat201003/summorist-users/internal/database"
	"github.com/Azat201003/summorist-users/internal/tokens"
)

type UserServer struct {
	pb.UnimplementedUsersServer
	DBC *database.DBController
}

func (s *UserServer) getUserByJwt(jwtToken string) (*pb.User, error) {
	userId, err := tokens.ValidateToken(jwtToken)
	if err != nil {
		return nil, err
	}
	user, err := s.DBC.FindUser(&pb.User{UserId: userId})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServer) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	user, err := s.DBC.FindUser(&pb.User{
		Username:     request.Username,
		PasswordHash: request.PasswordHash,
	})
	
	if err != nil {
		return &pb.SignInResponse{Code: 1}, err
	}
	
	jwtToken, err := tokens.GenerateToken(user.UserId)
	
	if err != nil {
		return &pb.SignInResponse{Code: 2}, err
	}
	
	return &pb.SignInResponse{JwtToken: jwtToken, Code: 0, RefreshToken: user.RefreshToken}, nil
}

func (s *UserServer) Authorize(ctx context.Context, request *pb.AuthRequest) (*pb.AuthResponse, error) {
	userId, err := tokens.ValidateToken(request.JwtToken)

	if err != nil {
		return &pb.AuthResponse{Code: 1}, err
	}

	return &pb.AuthResponse{UserId: userId, Code: 0}, nil
}

func (s *UserServer) SignUp(ctx context.Context, user *pb.User) (*pb.StatusResponse, error) {
	user.RefreshToken = tokens.GenerateRefreshToken()
	_, err := s.DBC.CreateUser(user)

	if err != nil {
		return &pb.StatusResponse{Code: 1}, err
	}

	return &pb.StatusResponse{Code: 0}, nil
}

func (s *UserServer) RefreshTokens(ctx context.Context, request *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	jwtToken, err := tokens.GenerateToken(request.UserId)

	if err != nil {
		return &pb.RefreshResponse{Code: 1}, err
	}

	user, err := s.DBC.FindUser(&pb.User{
		UserId:     	request.UserId,
		RefreshToken: request.RefreshToken,
	})

	if err != nil {
		return &pb.RefreshResponse{Code: 2}, err
	}

	return &pb.RefreshResponse{
		RefreshToken:	user.RefreshToken,
		JwtToken:     jwtToken,
		Code:         0,
	}, nil
}

func (s *UserServer) GetFiltered(request *pb.GetFilteredRequest, stream pb.Users_GetFilteredServer) error {
	requester, err := s.getUserByJwt(request.JwtToken)

	var users []pb.User

	if err != nil || !requester.IsAdmin {
		users, err = s.DBC.FindUsersBriefly(request.Filter)
	} else {
		users, err = s.DBC.FindUsers(request.Filter)
	}

	for i := range users {
		if err := stream.Send(&users[i]); err != nil {
			return err
		}
	}
	return nil
}

func (s *UserServer) UpdateUser(ctx context.Context, request *pb.UpdateRequest) (*pb.StatusResponse, error) {
	user, err := s.getUserByJwt(request.JwtToken)
	
	if (err != nil) {
		return &pb.StatusResponse{Code: 1}, err
	}

	if (user == nil || user.UserId != request.User.UserId && !user.IsAdmin) {
		return &pb.StatusResponse{Code: 2}, ErrNotPermitted // Permission denieded
	}

	err = s.DBC.UpdateUser(request.User)
	if err != nil {
		return &pb.StatusResponse{Code: 3}, err
	}
	return &pb.StatusResponse{Code: 0}, nil
}

func (s *UserServer) RemoveUser(ctx context.Context, request *pb.RemoveRequest) (*pb.StatusResponse, error) {
	user, err := s.getUserByJwt(request.JwtToken)
	
	if (request.UserId == 0 || err != nil) {
		return &pb.StatusResponse{Code: 2}, ErrNotPermitted
	}

	if (user == nil || user.UserId != request.UserId && !user.IsAdmin) {
		return &pb.StatusResponse{Code: 3}, ErrNotPermitted // Permission denieded
	}

	err = s.DBC.DeleteUser(request.UserId)

	if err != nil {
		return &pb.StatusResponse{Code: 3}, err
	}

	return &pb.StatusResponse{Code: 0}, nil
}

func newServer() *UserServer {
	DBC := database.DBController{}
	err := DBC.InitDB()

	if err != nil {
		log.Fatalf("%v", err)
	}

	return &UserServer{DBC: &DBC}
}

func StartServer() {
	lis, _ := net.Listen("tcp", fmt.Sprintf("%v:%v", os.Getenv("USERS_HOST"), os.Getenv("USERS_PORT")))
	grpcServer := grpc.NewServer()
	pb.RegisterUsersServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
