package admindiv

import (
	"github.com/aquarius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aquarius6666/go-utils/database"
	"github.com/aquarius6666/go-utils/database/cockroach"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

	return db
}

type AdminDivCDBRepo struct {
	cockroach.CDBRepository
}

func (u *AdminDivCDBRepo) SelectAdminDiv(search *admindiv.Search) (*admindiv.AdminDiv, error) {
	r := admindiv.AdminDiv{}
	if err := applySearch(u.Db, search).First(&r).Error; err != nil {
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
