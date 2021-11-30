package cockroach

import (
	"github.com/aquarius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aquarius6666/citizen-v/src/internal/db/citizen"
	admindivcdb "github.com/aquarius6666/citizen-v/src/internal/db/cockroach/admindiv"
	citizencdb "github.com/aquarius6666/citizen-v/src/internal/db/cockroach/citizen"
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
	CitizenRepo  *citizencdb.CitizenCDBRepo
}

func (s *ServerCDBRepo) CountCitizen(search *citizen.Search) (*int64, error) {
	return s.CitizenRepo.CountCitizen(search)
}

func (s *ServerCDBRepo) ListCitizen(search *citizen.Search) ([]*citizen.Citizen, error) {
	return s.CitizenRepo.ListCitizen(search)
}

func (s *ServerCDBRepo) SelectCitizen(search *citizen.Search) (*citizen.Citizen, error) {
	return s.CitizenRepo.SelectCitizen(search)
}

func (s *ServerCDBRepo) InsertCitizen(u *citizen.Citizen) (*citizen.Citizen, error) {
	return s.CitizenRepo.InsertCitizen(u)
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

func (s *ServerCDBRepo) ListAdminDiv(u *admindiv.Search) ([]*admindiv.AdminDiv, error) {
	return s.AdminDivRepo.ListAdminDiv(u)
}

func (s *ServerCDBRepo) CountAdminDiv(u *admindiv.Search) (*int64, error) {
	return s.AdminDivRepo.CountAdminDiv(u)
}
