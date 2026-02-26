package database

import (
	pb "github.com/Azat201003/summorist-shared/gen/go/users"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
)

type DBController struct {
	DB *gorm.DB
}

func (dbc *DBController) CreateUser(user *pb.User) (uint64, error) {
	result := dbc.DB.Create(user)
	return user.UserId, result.Error
}

func (dbc *DBController) FindUsers(filter *pb.User) ([]pb.User, error) {
	var users []pb.User
	result := dbc.DB.Where(filter).Find(&users)
	return users, result.Error
}

// Finding user to change by newUser.Id
func (dbc *DBController) UpdateUser(newUser *pb.User) error {
	result := dbc.DB.Where("user_id = ?", newUser.UserId).Updates(newUser)
	return result.Error
}

func (dbc *DBController) DeleteUser(userId uint64) error {
	result := dbc.DB.Where("user_id = ?", userId).Delete(&pb.User{})
	return result.Error
}

func (dbc *DBController) FindUser(filter *pb.User) (*pb.User, error) {
	user := new(pb.User)
	result := dbc.DB.Where(filter).First(user)
	return user, result.Error
}

func (dbc *DBController) FindUsersBriefly(filter *pb.User) ([]pb.User, error) {	
	var users []pb.User
	result := dbc.DB.Where(filter).Select("user_id", "username", "is_admin").Find(&users)
	return users, result.Error
}

func (dbc *DBController) InitDB() error {	
	db, err := gorm.Open(postgres.Open(os.Getenv("USERS_POSTGRES_DSN")), &gorm.Config{})

	if err != nil {
		return err
	}

	dbc.DB = db

	return nil
}

