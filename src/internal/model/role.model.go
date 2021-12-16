package model

import "github.com/aqaurius6666/citizen-v/src/internal/db"

type Role interface {
}

var (
	_ Role = (*RoleModel)(nil)
)

type RoleModel struct {
	Repo db.ServerRepo
}
