package model

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db/role"
	"golang.org/x/xerrors"
)

type Role interface {
	GetChildRole(string) (string, error)
	GetRoleByName(string) (*role.Role, error)
}

func (s *ServerModel) GetChildRole(n string) (string, error) {

	switch n {
	case role.ROLE_A1:
		return role.ROLE_A2, nil
	case role.ROLE_A2:
		return role.ROLE_A3, nil
	case role.ROLE_A3:
		return role.ROLE_B1, nil
	case role.ROLE_B1:
		return role.ROLE_B2, nil
	default:
		return "", xerrors.New("invalid role name")
	}

}

func (s *ServerModel) GetRoleByName(n string) (*role.Role, error) {
	if n == "" {
		return nil, xerrors.New("empty")
	}
	return s.Repo.SelectRole(&role.Search{
		Role: role.Role{Name: &n},
	})
}
