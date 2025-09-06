// cmd/server.go
package main

import (
	"context"
	"log"
	"net"

	//"github.com/Azat201003/summorist-shared/gen/go/common"
	pb "github.com/Azat201003/summorist-shared/gen/go/user-service"
	"google.golang.org/grpc"
)

type userServer struct {
	pb.UnimplementedUsersServer
}

func (s *userServer) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	return &pb.SignInResponse{Token: "aoe"}, nil
}

func newServer() *userServer {
	return &userServer{}
}

func startServer() {
	lis, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUsersServer(grpcServer, newServer())
	log.Printf("Server starting on port 8001")
	grpcServer.Serve(lis)
}
