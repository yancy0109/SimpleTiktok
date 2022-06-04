package repository

import (
	"fmt"
	"time"
)

type User struct {
        ID      uint
        User_id int `gorm:"default:(-)"`
        User_name string `gorm:"default:(-)"`
        Password string `gorm:"default:(-)"`
        Salt    string `gorm:"default:(-)"`
        Create_date time.Time
}

func (User) TableName() string{
	return "user"
}

func CreateUser(username string, password string, salt string) (int, error) {
	user0 := User{User_id:time.Now().Nanosecond(), User_name:username, Password:password, Salt:salt, Create_date: time.Now()}

	result := db.Create(&user0)

	if result.Error != nil {
		fmt.Printf("failed to insert new user, err: %e\n", result.Error)
		return 0, result.Error
	}
	// userid deprecated, primary key used instead
	// return int(time.Now().UnixNano()), result.Error
	return int(user0.ID), nil
}

func FindUser(username string) (User, error) {
	user0 := User{User_name: username}
	result := db.Where("user_name = ?", username).First(&user0)
	err := result.Error
	if err != nil {
		fmt.Printf("failed to find user, err: %e\n", err)
		return User{}, err
	}
	return user0, err
}
//func FindUsers(username string) (User[], error) {
//	return make([]User, 0), nil
//}

