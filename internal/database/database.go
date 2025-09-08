package database

import (
	"gorm.io/gorm"
	//	"gorm.io/driver/postgres"
	// "context"
	//"fmt"
	"github.com/Azat201003/summorist-shared/gen/go/common"
)

type DBController struct {
	DB *gorm.DB
}

func (dbc *DBController) CreateUser(user *common.User) error {
	result := dbc.DB.Create(user)
	return result.Error
}

func (dbc *DBController) FindUsers(filter *common.User) ([]common.User, error) {
	var users []common.User
	result := dbc.DB.Where(filter).Find(&users)
	return users, result.Error
}

