package db

import (
	"example.com/Sinezx/words-server/util"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account  string `json:"account"`
	Password string `json:"password"`
}

func InsertUser(account string, password string) (uint, error) {
	user := User{Account: account, Password: util.Md5(password)}
	result := gormDB.Create(&user)
	return user.ID, result.Error
}

func QueryUserByAccount(account string) (*User, error) {
	user := User{Account: account}
	result := gormDB.Where("account = ?", user.Account).First(&user)
	return &user, result.Error
}

func HardDeleteUser(userId uint) (int64, error) {
	user := User{}
	user.ID = userId
	result := gormDB.Unscoped().Delete(&user)
	return result.RowsAffected, result.Error
}
