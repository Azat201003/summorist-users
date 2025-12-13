package database

import (
	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"gorm.io/gorm"
)

type DBController struct {
	DB *gorm.DB
}

func (dbc *DBController) CreateUser(user *pb.User) (uint64, error) {
	result := dbc.DB.Create(user)
	return user.Id, result.Error
}

func (dbc *DBController) FindUsers(filter *pb.User) ([]pb.User, error) {
	var users []pb.User
	result := dbc.DB.Where(filter).Find(&users)
	return users, result.Error
}

// Finding user to change by newUser.Id
func (dbc *DBController) UpdateUser(newUser *pb.User) error {
	result := dbc.DB.Updates(newUser)
	return result.Error
}

func (dbc *DBController) DeleteUser(userId uint64) error {
	user := &pb.User{Id: userId}
	result := dbc.DB.Delete(user)
	return result.Error
}
