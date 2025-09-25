package tokens

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/Azat201003/summorist-users/internal/database"
	"github.com/Azat201003/summorist-shared/gen/go/common"
	"errors"
	"strconv"
	"time"
)

func GenerateToken(privateKey []byte, userId uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
	})

	parsed, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString(parsed)

	return tokenString, err
}

func ValidateKey(tokenString string, dbc database.DBController) (uint64, error) {
	var user_id uint64

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		str_id, _ := token.Claims.GetSubject()
		user_id, _ = strconv.ParseUint(str_id, 10, 64)
		users, _ := dbc.FindUsers(&common.User{
			Id: uint64(user_id),
		})
		if len(users) == 0 {
			return nil, errors.New("No user found.")
		}
		found, _ := dbc.FindTokenKeys(&common.TokenKeys{
				Id: users[0].TokenId,
		})
		res, _ := jwt.ParseRSAPublicKeyFromPEM(found[0].PublicKey)
		return res, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return 0, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok {
		return user_id, nil
	} else {
		return 0, err
	}
}

