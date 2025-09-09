package database

import (
	"gorm.io/gorm"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"crypto/rand"
	"github.com/Azat201003/summorist-shared/gen/go/common"
)

type DBController struct {
	DB *gorm.DB
}

func (dbc *DBController) CreateUser(user *common.User) (uint64, error) {
	result := dbc.DB.Create(user)
	return user.Id, result.Error
}

func (dbc *DBController) FindUsers(filter *common.User) ([]common.User, error) {
	var users []common.User
	result := dbc.DB.Where(filter).Find(&users)
	return users, result.Error
}

// Finding user to change by newUser.Id
func (dbc *DBController) UpdateUser(newUser *common.User) error {
	result := dbc.DB.Save(newUser)
	return result.Error
}



func (dbc *DBController) CreateTokenKeys(tokenKeys *common.TokenKeys) (uint64, error) {
	if tokenKeys.PrivateKey == nil {
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return 0, err
		}
		der, _ := x509.MarshalPKCS8PrivateKey(privateKey)
		tokenKeys.PrivateKey = pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: der,
		})
		publicKey := &privateKey.PublicKey
		der, _ = x509.MarshalPKCS8PrivateKey(publicKey)
		tokenKeys.PublicKey = pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: der,
		})
	}
	result := dbc.DB.Create(tokenKeys)
	return tokenKeys.Id, result.Error
}

func (dbc *DBController) FindTokenKeys(filter *common.TokenKeys) ([]common.TokenKeys, error) {
	var tokenKeys []common.TokenKeys
	result := dbc.DB.Where(filter).Find(&tokenKeys)
	return tokenKeys, result.Error
}

// Finding token keys to change by newTokenKeys.Id
func (dbc *DBController) UpdateTokenKeys(newTokenKeys *common.TokenKeys) error {
	result := dbc.DB.Save(newTokenKeys)
	return result.Error
}
