package database

import (
	"github.com/Azat201003/summorist-shared/gen/go/common"
	"gorm.io/gorm"
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
