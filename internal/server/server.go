package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Azat201003/summorist-shared/gen/go/user-service"
	"google.golang.org/grpc"
)

type userServer struct {
	pb.UnimplementedUsersServer
}

func (s *userServer) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	return &pb.SignInResponse{JwtToken: "aoe"}, nil
}

func newServer() *userServer {
	return &userServer{}
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
