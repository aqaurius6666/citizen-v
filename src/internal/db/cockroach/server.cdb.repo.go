package cockroach

import (
	usercdb "github.com/aqaurius6666/boilerplate-server-go/src/internal/db/cockroach/user"
	"github.com/aqaurius6666/boilerplate-server-go/src/internal/db/user"
	"github.com/aquarius6666/go-utils/database/cockroach"
)

// var (
// 	_ db.ServerRepo = (*ServerCDBRepo)(nil)
// )

type ServerCDBRepo struct {
	cockroach.CDBRepository
	UserRepo *usercdb.UserCDBRepo
}

func (s *ServerCDBRepo) SelectOne(search *user.Search) (*user.User, error) {
	return s.UserRepo.SelectOne(search)
}
