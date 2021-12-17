package model

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/role"
	"github.com/aqaurius6666/citizen-v/src/internal/db/user"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/aqaurius6666/go-utils/utils"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type User interface {
	HasPermission(user, add uuid.UUID) (bool, error)
	IsRoleActive(userId uuid.UUID) (bool, error)
	GetRoleId(addid uuid.UUID) (uuid.UUID, error)
}

var (
	_ User = (*UserModel)(nil)
)

type UserModel struct {
	Repo db.ServerRepo
}

func (u *UserModel) GetRoleId(addid uuid.UUID) (uuid.UUID, error) {
	if addid == uuid.Nil {
		return uuid.Nil, nil
	}
	add, err := u.Repo.SelectAdminDiv(&admindiv.Search{
		AdminDiv: admindiv.AdminDiv{
			BaseModel: database.BaseModel{ID: addid},
		},
	})
	if err != nil {
		return uuid.Nil, err
	}
	var search string
	switch *add.Type {
	case admindiv.CITY:
		search = role.ROLE_A1
	case admindiv.DISTRICT:
		search = role.ROLE_A2
	case admindiv.TOWN:
		search = role.ROLE_A3
	case admindiv.BLOCK:
		search = role.ROLE_B1
	}
	ro, err := u.Repo.SelectRole(&role.Search{
		Role: role.Role{
			Name: &search,
		},
	})
	if err != nil {
		return uuid.Nil, err
	}
	return ro.ID, nil
}

func (u *UserModel) HasPermission(uid uuid.UUID, addid uuid.UUID) (bool, error) {
	if uid == uuid.Nil || addid == uuid.Nil {
		return false, nil
	}
	usr, err := u.Repo.SelectUser(&user.Search{
		User: user.User{
			BaseModel: database.BaseModel{ID: uid},
		},
	})
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}
	if *usr.Role.Name == role.ROLE_ADMIN {
		return true, nil
	}
	add, err := u.Repo.SelectAdminDiv(&admindiv.Search{
		AdminDiv: admindiv.AdminDiv{
			BaseModel: database.BaseModel{ID: addid},
		},
	})
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}
	valid := false
	tmpAdd := add
	for {
		if usr.AdminDivID == tmpAdd.ID {
			valid = true
			break
		}
		if tmpAdd.SuperiorID == uuid.Nil {
			break
		}
		if tmpAdd, err = u.Repo.SelectAdminDiv(&admindiv.Search{
			AdminDiv: admindiv.AdminDiv{
				BaseModel: database.BaseModel{
					ID: tmpAdd.SuperiorID,
				},
			},
		}); err != nil {
			return false, xerrors.Errorf("%w", err)
		}
	}

	return valid, nil
}

func (u *UserModel) IsRoleActive(userId uuid.UUID) (bool, error) {
	if userId == uuid.Nil {
		return false, nil
	}
	usr, err := u.Repo.SelectUser(&user.Search{
		User: user.User{
			BaseModel: database.BaseModel{ID: userId},
		},
	})
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}
	if *usr.Role.Name == role.ROLE_ADMIN {
		return true, nil
	}
	add, err := u.Repo.SelectAdminDiv(&admindiv.Search{
		AdminDiv: admindiv.AdminDiv{
			BaseModel: database.BaseModel{ID: usr.AdminDivID},
		},
	})
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}
	valid := utils.BoolVal(usr.IsActive)
	tmpAdd := add
	for valid {
		usr, err := u.Repo.SelectUser(&user.Search{
			User: user.User{
				AdminDivID: tmpAdd.ID,
			},
		})
		if err != nil {
			return false, xerrors.Errorf("%w", err)
		}
		if !utils.BoolVal(usr.IsActive) {
			valid = false
			break
		}
		if tmpAdd.SuperiorID == uuid.Nil {
			break
		}
		tmpAdd, err = u.Repo.SelectAdminDiv(&admindiv.Search{
			AdminDiv: admindiv.AdminDiv{
				BaseModel: database.BaseModel{ID: tmpAdd.SuperiorID},
			},
		})
		if err != nil {
			return false, xerrors.Errorf("%w", err)
		}

	}
	return valid, nil
}
