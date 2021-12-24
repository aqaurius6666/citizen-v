package admindiv

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/aqaurius6666/go-utils/database/cockroach"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	_ admindiv.AdminDivRepo = (*AdminDivCDBRepo)(nil)
)

func applySearch(db *gorm.DB, search *admindiv.Search) *gorm.DB {
	if search.ID != uuid.Nil {
		db = db.Where(&admindiv.AdminDiv{
			BaseModel: database.BaseModel{
				ID: search.ID,
			},
		})
	}
	if search.Name != nil {
		db = db.Where(&admindiv.AdminDiv{
			Name: search.Name,
		})
	}
	if search.Code != nil {
		db = db.Where(&admindiv.AdminDiv{
			Code: search.Code,
		})
	}
	if search.Type != nil {
		db = db.Where(&admindiv.AdminDiv{
			Type: search.Type,
		})
	}
	if search.SuperiorID != uuid.Nil {
		db = db.Where(&admindiv.AdminDiv{
			SuperiorID: search.SuperiorID,
		})
	}
	if search.SuperiorCode != nil {
		db = db.Where(clause.Like{
			Column: "code",
			Value:  *search.SuperiorCode + "__",
		})
	}

	orderBy := "name"
	isDesc := true
	if a := search.OrderBy; a != "" {
		orderBy = a

	}

	if search.Fields != nil {
		db = db.Select(search.Fields)
	}

	if orderType := search.OrderType; orderType != "DESC" {
		isDesc = false
	}
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: orderBy}, Desc: isDesc})
	db = db.Offset(search.Skip)
	db = db.Limit(search.Limit)
	return db
}

type AdminDivCDBRepo struct {
	cockroach.CDBRepository
}

func (u *AdminDivCDBRepo) SelectAdminDiv(search *admindiv.Search) (*admindiv.AdminDiv, error) {
	r := admindiv.AdminDiv{}
	if err := applySearch(u.Db, search).Debug().First(&r).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, admindiv.ErrNotFound
		}
		return nil, err
	}
	return &r, nil
}

func (u *AdminDivCDBRepo) InsertAdminDiv(value *admindiv.AdminDiv) (*admindiv.AdminDiv, error) {
	if err := u.Db.Create(value).Error; err != nil {
		return nil, admindiv.ErrInsertFail
	}
	return value, nil
}

func (u *AdminDivCDBRepo) ListAdminDiv(search *admindiv.Search) ([]*admindiv.AdminDiv, error) {
	r := make([]*admindiv.AdminDiv, 0)
	if err := applySearch(u.Db, search).Debug().Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (u *AdminDivCDBRepo) CountAdminDiv(search *admindiv.Search) (*int64, error) {
	var r int64
	if err := applySearch(u.Db, search).Model(&admindiv.AdminDiv{}).Count(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

func (u *AdminDivCDBRepo) UpdateAdminDiv(search *admindiv.Search, value *admindiv.AdminDiv) error {
	if err := applySearch(u.Db, search).Model(&admindiv.AdminDiv{}).Updates(value).Error; err != nil {
		return admindiv.ErrUpdateFail
	}
	return nil
}
