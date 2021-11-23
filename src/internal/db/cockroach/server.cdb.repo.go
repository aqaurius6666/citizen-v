package cockroach

import (
	usercdb "github.com/aquarius6666/citizen-v/src/internal/db/cockroach/user"
	"github.com/aquarius6666/citizen-v/src/internal/db/user"
	"github.com/aquarius6666/go-utils/database/cockroach"
)

// var (
// 	_ db.ServerRepo = (*ServerCDBRepo)(nil)
// )

type ServerCDBRepo struct {
	cockroach.CDBRepository
	UserRepo *usercdb.UserCDBRepo
}

func (s *ServerCDBRepo) SelectUser(search *user.Search) (*user.User, error) {
	return s.UserRepo.SelectUser(search)
}

func (s *ServerCDBRepo) InsertUser(u *user.User) (*user.User, error) {
	return s.UserRepo.InsertUser(u)
}
