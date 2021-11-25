package cockroach

import (
	"github.com/aquarius6666/citizen-v/src/internal/db/admindiv"
	admindivcdb "github.com/aquarius6666/citizen-v/src/internal/db/cockroach/admindiv"
	rolecdb "github.com/aquarius6666/citizen-v/src/internal/db/cockroach/role"
	usercdb "github.com/aquarius6666/citizen-v/src/internal/db/cockroach/user"
	"github.com/aquarius6666/citizen-v/src/internal/db/role"
	"github.com/aquarius6666/citizen-v/src/internal/db/user"
	"github.com/aquarius6666/go-utils/database/cockroach"
)

// var (
// 	_ db.ServerRepo = (*ServerCDBRepo)(nil)
// )

type ServerCDBRepo struct {
	cockroach.CDBRepository
	UserRepo     *usercdb.UserCDBRepo
	RoleRepo     *rolecdb.RoleCDBRepo
	AdminDivRepo *admindivcdb.AdminDivCDBRepo
}

func (s *ServerCDBRepo) SelectUser(search *user.Search) (*user.User, error) {
	return s.UserRepo.SelectUser(search)
}

func (s *ServerCDBRepo) InsertUser(u *user.User) (*user.User, error) {
	return s.UserRepo.InsertUser(u)
}

func (s *ServerCDBRepo) SelectRole(search *role.Search) (*role.Role, error) {
	return s.RoleRepo.SelectRole(search)
}

func (s *ServerCDBRepo) InsertRole(u *role.Role) (*role.Role, error) {
	return s.RoleRepo.InsertRole(u)
}

func (s *ServerCDBRepo) SelectAdminDiv(search *admindiv.Search) (*admindiv.AdminDiv, error) {
	return s.AdminDivRepo.SelectAdminDiv(search)
}

func (s *ServerCDBRepo) InsertAdminDiv(u *admindiv.AdminDiv) (*admindiv.AdminDiv, error) {
	return s.AdminDivRepo.InsertAdminDiv(u)
}
