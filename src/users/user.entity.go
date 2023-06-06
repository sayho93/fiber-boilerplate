package users

import (
	"fiber/src/common/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func init() {
	database.Connection.AutoMigrate(User{})
}
