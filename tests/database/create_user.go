package database_tests

import (
	"github.com/Azat201003/summorist-shared/gen/go/common"
	"github.com/Azat201003/summorist-users/internal/passwords"
    "crypto/rand"
    "encoding/hex"
	"fmt"
)

func (s *databaseSuite) TestCreateUserOk() {
	tokenId, err := s.dbc.CreateTokenKeys(new(common.TokenKeys))
	if err != nil {
		fmt.Println("Cant create token")
		return
	}
	_, err = s.dbc.CreateUser(&common.User{
		Username: "test-" + generateRandomString(10),
		PasswordHash: passwords.Hash(generateRandomString(16)),
		TokenId: tokenId,
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

