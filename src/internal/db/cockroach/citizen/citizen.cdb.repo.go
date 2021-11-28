package citizen

import (
	"github.com/aquarius6666/citizen-v/src/internal/db/citizen"
	"github.com/aquarius6666/go-utils/database"
	"github.com/aquarius6666/go-utils/database/cockroach"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if search.Birthday != nil {
		db = db.Where(&citizen.Citizen{
			Birthday: search.Birthday,
		})
	}
	if search.PID != nil {
		db = db.Where(&citizen.Citizen{
			PID: search.PID,
		})
	}
	if skip := search.DefaultSearchModel.Skip; skip != 0 {
		db = db.Offset(skip)
	}
	if limit := search.DefaultSearchModel.Skip; limit != 0 {
		db = db.Limit(limit)
	}

	orderBy := "name"
	isDesc := true
	if a := search.OrderBy; a != "" {
		orderBy = a

	}
	if orderType := search.OrderType; orderType != "DESC" {
		isDesc = false
	}
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: orderBy}, Desc: isDesc})
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

func (u *CitizenCDBRepo) ListCitizen(search *citizen.Search) ([]*citizen.Citizen, error) {
	r := make([]*citizen.Citizen, 0)
	if err := applySearch(u.Db, search).Debug().Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}
