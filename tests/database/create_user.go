package database_tests

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/Azat201003/summorist-shared/gen/go/common"
	"github.com/Azat201003/summorist-users/internal/passwords"
	"github.com/Azat201003/summorist-users/internal/tokens"
)

func (s *databaseSuite) TestCreateUserOk() {
	_, err := s.dbc.CreateUser(&common.User{
		Username:       "test-" + generateRandomString(10),
		PasswordHash:   passwords.Hash(generateRandomString(16)),
		RefreshToken:	tokens.GenerateRefreshToken(),
	})
	s.NoError(err)
}

func generateRandomString(n int) string {
	bytes := make([]byte, (n+1)/2)
	if _, err := rand.Read(bytes); err != nil {
		return "error"
	}
	return hex.EncodeToString(bytes)[:n]
}
