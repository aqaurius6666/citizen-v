package db

import (
	"github.com/aqaurius6666/boilerplate-server-go/src/internal/db/user"
	"github.com/aquarius6666/go-utils/database"
)

type ServerRepo interface {
	database.CommonRepository
	user.UserRepo
}

