package model

import (
	"context"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/sirupsen/logrus"
)

type Server interface {
	Role
	User
	AdminDiv
	Citizen
	Campaign
}

var (
	_ Server = (*ServerModel)(nil)
)

func NewServerModel(ctx context.Context, logger *logrus.Logger, repo db.ServerRepo) (Server, error) {
	model, err := InitModel(ctx, logger, repo)
	return model, err
}

type ServerModel struct {
	Repo db.ServerRepo
}
