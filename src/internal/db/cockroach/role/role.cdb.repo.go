package role

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db/role"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/aqaurius6666/go-utils/database/cockroach"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	_ role.RoleRepo = (*RoleCDBRepo)(nil)
)

func applySearch(db *gorm.DB, search *role.Search) *gorm.DB {
	if search.ID != uuid.Nil {
		db = db.Where(&role.Role{
			BaseModel: database.BaseModel{
				ID: search.ID,
			},
		})
	}
	if search.Name != nil {
		db = db.Where(&role.Role{
			Name: search.Name,
		})
	}

	return db
}

type RoleCDBRepo struct {
	cockroach.CDBRepository
}

func (u *RoleCDBRepo) SelectRole(search *role.Search) (*role.Role, error) {
	r := role.Role{}
	if err := applySearch(u.Db, search).Select(search.Fields).First(&r).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, role.ErrNotFound
		}
		return nil, err
	}
	return &r, nil
}

func (u *RoleCDBRepo) InsertRole(value *role.Role) (*role.Role, error) {
	if err := u.Db.Debug().Create(value).Error; err != nil {
		return nil, role.ErrInsertFail
	}
	return value, nil
}
