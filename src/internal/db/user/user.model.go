package user

import "github.com/aquarius6666/go-utils/database"

type User struct {
	database.BaseModel
	Username     *string `gorm:"column:username;not null"`
	HashPassword *string `gorm:"column:hash_password;not null"`
}

type Search struct {
	database.DefaultSearchModel
	User
}
