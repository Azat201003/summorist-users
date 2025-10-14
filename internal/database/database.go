package database

import (
	"gorm.io/gorm"
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



func (dbc *DBController) CreateRefreshToken(token *common.RefreshToken) (uint64, error) {
	result := dbc.DB.Create(token)
	return token.Id, result.Error
}

func (dbc *DBController) FindRefreshToken(filter *common.RefreshToken) ([]common.RefreshToken, error) {
	var tokens []common.RefreshToken
	result := dbc.DB.Where(filter).Find(&tokens)
	return tokens, result.Error
}

// Finding token keys to change by newTokenKeys.Id
func (dbc *DBController) UpdateTokenKeys(newRefreshToken *common.RefreshToken) error {
	result := dbc.DB.Save(newRefreshToken)
	return result.Error
}
