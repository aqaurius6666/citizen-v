//go:build wireinject
// +build wireinject

package model

import (
	"context"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

func InitModel(ctx context.Context, logger *logrus.Logger, repo db.ServerRepo) (*ServerModel, error) {

	wire.Build(
		wire.Struct(new(ServerModel), "*"),
	)
	return &ServerModel{}, nil
}
