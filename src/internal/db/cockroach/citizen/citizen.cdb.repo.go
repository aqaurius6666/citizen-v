package citizen

import (
	"github.com/aquarius6666/citizen-v/src/internal/db/citizen"
	"github.com/aquarius6666/go-utils/database"
	"github.com/aquarius6666/go-utils/database/cockroach"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	_ citizen.CitizenRepo = (*CitizenCDBRepo)(nil)
)

func applySearch(db *gorm.DB, search *citizen.Search) *gorm.DB {
	if search.ID != uuid.Nil {
		db = db.Where(&citizen.Citizen{
			BaseModel: database.BaseModel{
				ID: search.ID,
			},
		})
	}
	if search.Name != nil {
		db = db.Where(&citizen.Citizen{
			Name: search.Name,
		})
	}

	return db
}

type CitizenCDBRepo struct {
	cockroach.CDBRepository
}

func (u *CitizenCDBRepo) SelectCitizen(search *citizen.Search) (*citizen.Citizen, error) {
	r := citizen.Citizen{}
	if err := applySearch(u.Db, search).First(&r).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, citizen.ErrNotFound
		}
		return nil, err
	}
	return &r, nil
}

func (u *CitizenCDBRepo) InsertCitizen(value *citizen.Citizen) (*citizen.Citizen, error) {
	if err := u.Db.Debug().Create(value).Error; err != nil {
		return nil, citizen.ErrInsertFail
	}
	return value, nil
}
