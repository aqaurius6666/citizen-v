package model

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db/campaign"
	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/aqaurius6666/go-utils/utils"
	"golang.org/x/xerrors"
)

type Citizen interface {
	InsertCitizen(*citizen.Citizen) (*citizen.Citizen, error)
	UpdateCitizen(*citizen.Search, *citizen.Citizen) error
}

func (s *ServerModel) InsertCitizen(c *citizen.Citizen) (*citizen.Citizen, error) {
	var err error
	c.CurrentPlace, err = lib.GetAdminDivFullNameCode(*c.CurrentPlaceCode, s.Repo)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	c.Hometown, err = lib.GetAdminDivFullNameCode(*c.HometownCode, s.Repo)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	c.ResidencePlace, err = lib.GetAdminDivFullNameCode(*c.ResidencePlaceCode, s.Repo)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	camp, err := s.GetValidCampaign(c.AdminDivCode)
	if err != nil {
		return nil, e.ErrNotCampaignYet
	}

	ctz, err := s.Repo.InsertCitizen(c)
	if err != nil {
		return nil, err
	}
	err = s.Repo.UpdateCampaign(&campaign.Search{
		Campaign: campaign.Campaign{BaseModel: database.BaseModel{
			ID: camp.ID,
		}},
	}, &campaign.Campaign{
		RecordNumber: utils.IntPtr(utils.IntVal(camp.RecordNumber) + 1),
	})
	if err != nil {
		return nil, err
	}
	return ctz, err
}

func (s *ServerModel) UpdateCitizen(k *citizen.Search, c *citizen.Citizen) error {
	var err error
	c.CurrentPlace, err = lib.GetAdminDivFullNameCode(*c.CurrentPlaceCode, s.Repo)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	c.Hometown, err = lib.GetAdminDivFullNameCode(*c.HometownCode, s.Repo)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	c.ResidencePlace, err = lib.GetAdminDivFullNameCode(*c.ResidencePlaceCode, s.Repo)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	err = s.Repo.UpdateCitizen(k, c)
	return err
}
