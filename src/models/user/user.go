package models

import (
	"Backend/backend-api/src/config"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Address  string
	PhoneNum int
}

func CreateUser(newUser *User) *gorm.DB {
	return config.DB.Create(&newUser)
}

func FindEmail(input *User) []User {
	item := []User{}
	config.DB.Raw("SELECT * FROM users WHERE email = ?", input.Email).Scan(&item)
	return item
}
