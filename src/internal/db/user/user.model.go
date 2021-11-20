package user

import "github.com/aquarius6666/go-utils/database"

type User struct {
	database.BaseModel
}

type Search struct {
	database.DefaultSearchModel
	User
}
