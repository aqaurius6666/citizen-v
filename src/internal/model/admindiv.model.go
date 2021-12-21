package model

import (
	"fmt"

	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/google/uuid"
)

type AdminDiv interface {
	GetNewCode(uuid.UUID) (string, error)
	IsChild(code string, codes []string) bool
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
	if len(codes) == 0 || code == "" {
		return false
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
