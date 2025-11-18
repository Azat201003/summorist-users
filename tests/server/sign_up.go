package server_tests

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"github.com/Azat201003/summorist-shared/gen/go/common"
	"github.com/Azat201003/summorist-users/internal/passwords"
)

func (s *serverSuite) TestCreateUserOk() {
	username := "test-" + generateRandomString(10)

	response, err := (*s.usersClient).SignUp(context.Background(), &common.User{
		Username:     username,
		PasswordHash: passwords.Hash(generateRandomString(16)),
	})
	s.NoError(err)
	s.Equal(response.Code, int32(0))
}

func generateRandomString(n int) string {
	bytes := make([]byte, (n+1)/2)
	if _, err := rand.Read(bytes); err != nil {
		return "error"
	}
	return hex.EncodeToString(bytes)[:n]
}
