package model

import (
	"fmt"

	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type AdminDiv interface {
	GetNewCode(uuid.UUID) (string, error)
	IsChild(code string, codes []string) bool
	GetChildType(uuid.UUID) (string, error)
	GetAdminDivByCode(string) (*admindiv.AdminDiv, error)
	GetAdminDivById(addid uuid.UUID) (*admindiv.AdminDiv, error)
}

func (s *ServerModel) GetAdminDivByCode(code string) (*admindiv.AdminDiv, error) {

	if code == "" {
		return nil, xerrors.New("empty")
	}
	return s.Repo.SelectAdminDiv(&admindiv.Search{
		AdminDiv: admindiv.AdminDiv{Code: &code},
	})
}

func (s *ServerModel) GetAdminDivById(addid uuid.UUID) (*admindiv.AdminDiv, error) {

	if addid == uuid.Nil {
		return nil, xerrors.New("empty")
	}
	return s.Repo.SelectAdminDiv(&admindiv.Search{
		AdminDiv: admindiv.AdminDiv{
			BaseModel: database.BaseModel{ID: addid},
		},
	})
}

func (s *ServerModel) GetChildType(superiorId uuid.UUID) (string, error) {
	if superiorId == uuid.Nil {
		return admindiv.CITY, nil
	}
	add, err := s.Repo.SelectAdminDiv(&admindiv.Search{
		AdminDiv: admindiv.AdminDiv{
			BaseModel: database.BaseModel{ID: superiorId},
		},
	})
	if err != nil {
		return "", xerrors.Errorf("%w", err)
	}
	switch *add.Type {
	case admindiv.CITY:
		return admindiv.DISTRICT, nil
	case admindiv.DISTRICT:
		return admindiv.TOWN, nil
	case admindiv.TOWN:
		return admindiv.BLOCK, nil
	default:
		return "", xerrors.New("invalid type")
	}
}

func (s *ServerModel) GetNewCode(superid uuid.UUID) (string, error) {
	superCode := ""
	var count *int64
	var err error
	if superid != uuid.Nil {
		add, err := s.Repo.SelectAdminDiv(&admindiv.Search{
			AdminDiv: admindiv.AdminDiv{
				BaseModel: database.BaseModel{ID: superid},
			},
		})
		if err != nil {
			return "", err
		}
		superCode = *add.Code
		count, err = s.Repo.CountAdminDiv(&admindiv.Search{
			AdminDiv: admindiv.AdminDiv{
				SuperiorID: superid,
			},
		})
		if err != nil {
			return "", err
		}
	} else {
		count, err = s.Repo.CountAdminDiv(&admindiv.Search{
			AdminDiv: admindiv.AdminDiv{
				Type: &admindiv.CITY,
			},
		})
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("%s%02d", superCode, *count+1), nil
}

func (s *ServerModel) IsChild(code string, codes []string) bool {
	if len(codes) == 0 {
		return false
	}
	if len(code) == 0 {
		return true
	}
	l := len(code)
	for _, c := range codes {
		if len(c) < l {
			return false
		}
		if len(c)-2 != l {
			return false
		}
		if c[0:l] != code {
			return false
		}
	}
	return true
}
