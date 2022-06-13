package repository

import (
	"time"
)

type User struct {
	ID          int64
	User_id     int    `gorm:"default:(-)"`
	User_name   string `gorm:"default:(-)"`
	Password    string `gorm:"default:(-)"`
	Salt        string `gorm:"default:(-)"`
	Create_date time.Time
}

func (User) TableName() string {
	return "user"
}

func CreateUser(username string, password string, salt string) (int64, error) {
	user0 := User{User_id: time.Now().Nanosecond(), User_name: username, Password: password, Salt: salt, Create_date: time.Now()}

	result := db.Create(&user0)

	if result.Error != nil {
		return 0, result.Error
	}
	return user0.ID, nil
}

func FindUser(username string) (User, error) {
	user0 := User{User_name: username}
	result := db.Where("user_name = ?", username).First(&user0)
	err := result.Error
	if err != nil {
		return User{}, err
	}
	return user0, err
}

func FindUserById(id int64) (User, error) {
	user := User{ID: id}
	result := db.Model(user).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
