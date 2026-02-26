package tokens

import (
	"crypto/rand"
	"github.com/golang-jwt/jwt/v5"
	"math/big"
	"time"
	"os"
)

func GenerateToken(userId uint64) (string, error) {
	privateKey := os.Getenv("USERS_PRIVATE_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
	})

	parsed, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString(parsed)

	return tokenString, err
}

func ValidateToken(tokenString string) (uint64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		key := os.Getenv("USERS_PUBLIC_KEY")
		parsed, err := jwt.ParseRSAPublicKeyFromPEM([]byte(key))
		return parsed, err
	}, jwt.WithValidMethods([]string{jwt.SigningMethodRS512.Alg()}))
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return uint64(claims["sub"].(float64)), nil
	} else {
		return 0, err
	}
}

func GenerateRefreshToken() string {
	const (
		tokenLength = 256
		charset     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	)
	token := make([]byte, tokenLength)

	for i := range token {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		token[i] = charset[randomIndex.Int64()]
	}

	return string(token)
}
