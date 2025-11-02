package tokens

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"time"
)

func GetPublicKey(base string) (string, error) {
	var m map[string]string
	m, err := godotenv.Read(base + "secrets.env")
	if err != nil {
		return "", err
	}
	if key, exists := m["PUBLIC_KEY"]; exists {
		return key, nil
	} else {
		return "", errors.New("PUBLIC_KEY not found in secrets.env")
	}
}

func GetPrivateKey(base string) (string, error) {
	m, err := godotenv.Read(base + "secrets.env")
	if err != nil {
		return "", err
	}
	if key, exists := m["PRIVATE_KEY"]; exists {
		return key, nil
	} else {
		return "", errors.New("PRIVATE_KEY not found in secrets.env")
	}
}

func GenerateToken(userId uint64, base string) (string, error) {
	privateKey, err := GetPrivateKey(base)

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

func ValidateToken(tokenString, base string) (uint64, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		key, err := GetPublicKey(base)
		if err != nil {
			return nil, err
		}
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
