package cockroach

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/campaign"
	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
	admindivcdb "github.com/aqaurius6666/citizen-v/src/internal/db/cockroach/admindiv"
	campaigncdb "github.com/aqaurius6666/citizen-v/src/internal/db/cockroach/campaign"
	citizencdb "github.com/aqaurius6666/citizen-v/src/internal/db/cockroach/citizen"
	rolecdb "github.com/aqaurius6666/citizen-v/src/internal/db/cockroach/role"
	usercdb "github.com/aqaurius6666/citizen-v/src/internal/db/cockroach/user"
	"github.com/aqaurius6666/citizen-v/src/internal/db/role"
	"github.com/aqaurius6666/citizen-v/src/internal/db/user"
	"github.com/aqaurius6666/go-utils/database/cockroach"
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
	CampaignRepo *campaigncdb.CampaignCDBRepo
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

func (s *ServerCDBRepo) ListUser(u *user.Search) ([]*user.User, error) {
	return s.UserRepo.ListUser(u)
}

func (s *ServerCDBRepo) CountAdminDiv(u *admindiv.Search) (*int64, error) {
	return s.AdminDivRepo.CountAdminDiv(u)
}

func (s *ServerCDBRepo) UpdateAdminDiv(u *admindiv.Search, v *admindiv.AdminDiv) error {
	return s.AdminDivRepo.UpdateAdminDiv(u, v)
}

func (s *ServerCDBRepo) UpdateCitizen(u *citizen.Search, v *citizen.Citizen) error {
	return s.CitizenRepo.UpdateCitizen(u, v)
}
func (s *ServerCDBRepo) UpdateUser(u *user.Search, v *user.User) error {
	return s.UserRepo.UpdateUser(u, v)
}

func (s *ServerCDBRepo) CountUser(u *user.Search) (*int64, error) {
	return s.UserRepo.CountUser(u)
}

func (s *ServerCDBRepo) CountCampaign(u *campaign.Search) (*int64, error) {
	return s.CampaignRepo.CountCampaign(u)
}

func (s *ServerCDBRepo) UpdateCampaign(u *campaign.Search, v *campaign.Campaign) error {
	return s.CampaignRepo.UpdateCampaign(u, v)
}

func (s *ServerCDBRepo) ListCampaign(u *campaign.Search) ([]*campaign.Campaign, error) {
	return s.CampaignRepo.ListCampaign(u)
}

func (s *ServerCDBRepo) InsertCampaign(u *campaign.Campaign) (*campaign.Campaign, error) {
	return s.CampaignRepo.InsertCampaign(u)
}

func (s *ServerCDBRepo) SelectCampaign(search *campaign.Search) (*campaign.Campaign, error) {
	return s.CampaignRepo.SelectCampaign(search)
}
func (s *ServerCDBRepo) TotalCampaignRecord(search *campaign.Search) (*campaign.Campaign, error) {
	return s.CampaignRepo.TotalCampaignRecord(search)
}

func (s *ServerCDBRepo) DeleteCitizen(search *citizen.Search) error {
	return s.CitizenRepo.DeleteCitizen(search)
}
