package user

import (
	"github.com/aqaurius6666/boilerplate-server-go/src/internal/db/user"
	"github.com/aquarius6666/go-utils/database/cockroach"
)

var (
	_ user.UserRepo = (*UserCDBRepo)(nil)
)

type UserCDBRepo struct {
	cockroach.CDBRepository
}

func (u *UserCDBRepo) SelectOne(_ *user.Search) (*user.User, error) {
	panic("not implemented") // TODO: Implement
}
