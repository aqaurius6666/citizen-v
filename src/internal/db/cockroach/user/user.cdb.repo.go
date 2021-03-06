package user

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db/user"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/aqaurius6666/go-utils/database/cockroach"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	_ user.UserRepo = (*UserCDBRepo)(nil)
)

func applySearch(db *gorm.DB, search *user.Search) *gorm.DB {
	if search.ID != uuid.Nil {
		db = db.Where(&user.User{
			BaseModel: database.BaseModel{
				ID: search.ID,
			},
		})
	}
	if search.Username != nil {
		db = db.Where(&user.User{
			Username: search.Username,
		})
	}
	if search.AdminDivID != uuid.Nil {
		db = db.Where(&user.User{
			AdminDivID: search.AdminDivID,
		})

	}
	if search.HashPassword != nil {
		db = db.Where(&user.User{
			HashPassword: search.HashPassword,
		})

	}
	if search.SuperiorCode != nil {
		db = db.Where(clause.And(clause.Like{
			Column: "admin_div_code",
			Value:  *search.SuperiorCode + "__",
		}))
	}

	return db
}

type UserCDBRepo struct {
	cockroach.CDBRepository
}

func (u *UserCDBRepo) SelectUser(search *user.Search) (*user.User, error) {
	r := user.User{}
	if err := applySearch(u.Db, search).Joins("Role").Joins("AdminDiv").First(&r).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, user.ErrNotFound
		}
		return nil, err
	}
	return &r, nil
}

func (u *UserCDBRepo) UpdateUser(search *user.Search, value *user.User) error {
	if err := applySearch(u.Db, search).Model(&user.User{}).Updates(value).Error; err != nil {
		return user.ErrUpdateFail
	}
	return nil
}

func (u *UserCDBRepo) InsertUser(value *user.User) (*user.User, error) {
	if err := u.Db.Debug().Create(value).Error; err != nil {
		return nil, user.ErrInsertFail
	}
	return value, nil
}

func (u *UserCDBRepo) CountUser(search *user.Search) (*int64, error) {
	var r int64
	if err := applySearch(u.Db, search).Model(&user.User{}).Count(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

func (u *UserCDBRepo) ListUser(search *user.Search) ([]*user.User, error) {
	r := make([]*user.User, 0)
	if err := applySearch(u.Db, search).Offset(search.Skip).Limit(search.Limit).Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}
