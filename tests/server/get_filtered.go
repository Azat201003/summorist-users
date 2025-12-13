package server_tests

import (
	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"context"
	"io"
)

func (s *serverSuite) TestGetFilteredOk() {
	stream, err := (*s.usersClient).GetFiltered(context.Background(), &pb.User{
		Username: "Abeme",
	})
	s.NoError(err)
	user, err := stream.Recv()
	s.Equal(user.Username, "Abeme")
	_, err = stream.Recv()
	s.Equal(err, io.EOF)
}

func (s *serverSuite) TestGetFilteredNotFound() {
	stream, err := (*s.usersClient).GetFiltered(context.Background(), &pb.User{
		Username: "Name that does not exists so I believe it",
	})
	s.NoError(err)
	_, err = stream.Recv()
	s.Equal(err, io.EOF)
}

