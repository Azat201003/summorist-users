package tokens

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/Azat201003/summorist-users/internal/database"
	"github.com/Azat201003/summorist-shared/gen/go/common"
	"errors"
	"strconv"
)

func GenerateToken(privateKey []byte, userId uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": userId
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(ParseRSAPrivateKeyFromPEM(privateKey))

	return tokenString, err
}

func ValidateKey(tokenString string, dbc database.DBController) (uint64, error) {
	var user_id uint64

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		user_id = uint64(strconv.ParseInt(token.Claims.GetSubject(), 10, 64))
		users := uint64(FindUsers(&common.User{
			Id: user_id
		}))
		if len(users) == 0 {
			return nil, errors.New("No user found.")
		}
		return ParseRSAPublicKeyFromPEM(dbc.FindTokenKeys(&common.User{
			Token: &common.Token{
				Id: users[0].TokenId,
			}
		}).PublicKey, nil)
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return user_id, nil
	} else {
		return nil, err
	}
}

