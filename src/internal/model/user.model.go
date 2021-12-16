package model

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/user"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type User interface {
	HasPermission(user, add uuid.UUID) (bool, error)
}

var (
	_ User = (*UserModel)(nil)
)

type UserModel struct {
	Repo db.ServerRepo
}

func (u *UserModel) HasPermission(uid uuid.UUID, addid uuid.UUID) (bool, error) {
	usr, err := u.Repo.SelectUser(&user.Search{
		User: user.User{
			BaseModel: database.BaseModel{ID: uid},
		},
	})
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}
	add, err := u.Repo.SelectAdminDiv(&admindiv.Search{
		AdminDiv: admindiv.AdminDiv{
			BaseModel: database.BaseModel{ID: addid},
		},
	})
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}
	tmpUser := usr
	tmpAdd := add
	valid := *tmpUser.IsActive
	for *tmpUser.IsActive && tmpAdd.ID != uuid.Nil {

	}

	return valid, nil
}
