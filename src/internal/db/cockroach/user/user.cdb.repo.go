package user

import (
	"github.com/aqaurius6666/boilerplate-server-go/src/internal/db/user"
	"github.com/aquarius6666/go-utils/database"
	"github.com/aquarius6666/go-utils/database/cockroach"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

	return db
}

type UserCDBRepo struct {
	cockroach.CDBRepository
}

func (u *UserCDBRepo) SelectUser(search *user.Search) (*user.User, error) {
	r := user.User{}
	if err := applySearch(u.Db, search).First(&r).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, user.ErrNotFound
		}
		return nil, err
	}
	return &r, nil
}

func (u *UserCDBRepo) InsertUser(value *user.User) (*user.User, error) {
	if err := u.Db.Debug().Create(value).Error; err != nil {
		return nil, user.ErrInsertFail
	}
	return value, nil
}
