package model

import (
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
	HasPermissionByCode(user uuid.UUID, addCode string) (bool, error)
	IsRoleActive(userId uuid.UUID) (bool, error)
	GetRoleId(addid uuid.UUID) (uuid.UUID, error)
	GetUserById(uid uuid.UUID) (*user.User, error)
	ListUsers(search *user.Search) ([]*user.User, error)
	CheckPermissionCode(c1 string, c2 string) bool
}

func (s *ServerModel) ListUsers(search *user.Search) ([]*user.User, error) {
	usrs, err := s.Repo.ListUser(search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	for _, u := range usrs {
		tmp, err := s.IsRoleActive(u.ID)
		if err != nil {
			return nil, xerrors.Errorf("%w", err)
		}
		u.IsActive = &tmp
	}
	return usrs, nil
}

func (s *ServerModel) GetUserById(uid uuid.UUID) (*user.User, error) {
	usr, err := s.Repo.SelectUser(&user.Search{
		User: user.User{BaseModel: database.BaseModel{ID: uid}},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	isActive, err := s.IsRoleActive(usr.ID)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	usr.IsActive = &isActive
	return usr, nil
}

func (u *ServerModel) GetRoleId(addid uuid.UUID) (uuid.UUID, error) {
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

func (u *ServerModel) HasPermission(uid uuid.UUID, addid uuid.UUID) (bool, error) {
	if uid == uuid.Nil || addid == uuid.Nil {
		return false, nil
	}
	usr, err := u.GetUserById(uid)
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}
	add, err := u.GetAdminDivById(addid)
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}

	return u.CheckPermissionCode(*usr.AdminDivCode, *add.Code), nil
}

func (s *ServerModel) CheckPermissionCode(c1, c2 string) bool {
	if c1 == "" {
		return true
	}
	l1 := len(c1)
	l2 := len(c2)
	if l2 < l1 {
		return false
	}
	if c2[0:l1] == c1 {
		return true
	}
	return false
}

func (u *ServerModel) HasPermissionByCode(uid uuid.UUID, addCode string) (bool, error) {
	if uid == uuid.Nil || addCode == "" {
		return false, nil
	}
	usr, err := u.GetUserById(uid)
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}

	_, err = u.GetAdminDivByCode(addCode)
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}

	return u.CheckPermissionCode(*usr.AdminDivCode, addCode), nil
}

func (u *ServerModel) IsRoleActive(userId uuid.UUID) (bool, error) {
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
	if *usr.Role.Name == role.ROLE_A1 {
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
