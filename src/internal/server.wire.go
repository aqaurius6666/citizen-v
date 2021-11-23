//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/aquarius6666/citizen-v/src/internal/api"
	"github.com/aquarius6666/citizen-v/src/internal/db"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

type Server struct {
	ApiServer *api.ApiServer
	MainRepo  db.ServerRepo
}

type ServerOptions struct {
	DBDsn db.DBDsn
}

func InitMainServer(ctx context.Context, logger *logrus.Logger, opts ServerOptions) (*Server, error) {

	wire.Build(
		wire.FieldsOf(&opts, "DBDsn"),
		db.InitServerRepo,
		wire.Struct(new(api.ApiServerOptions), "*"),
		api.InitApiServer,
		wire.Struct(new(Server), "*"),
	)
	return &Server{}, nil
}
