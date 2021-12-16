package model

import (
	"context"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Server interface {
	Role
	User
}

var (
	_ Server = (*ServerModel)(nil)
)

func NewServerModel(ctx context.Context, logger *logrus.Logger, repo db.ServerRepo) (Server, error) {
	model, err := InitModel(ctx, logger, repo)
	return model, err
}

type ServerModel struct {
	Role *RoleModel
	User *UserModel
}

func (s *ServerModel) HasPermission(user uuid.UUID, add uuid.UUID) (bool, error) {
	return s.User.HasPermission(user, add)
}
