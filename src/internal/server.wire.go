//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/aqaurius6666/citizen-v/src/internal/api"
	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/model"
	"github.com/aqaurius6666/citizen-v/src/internal/services/excelexporter"
	"github.com/aqaurius6666/citizen-v/src/internal/services/jwt"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

type Server struct {
	ApiServer *api.ApiServer
	MainRepo  db.ServerRepo
}

type ServerOptions struct {
	DBDsn db.DBDsn
	Sec   jwt.SecretKey
}

func InitMainServer(ctx context.Context, logger *logrus.Logger, opts ServerOptions) (*Server, error) {

	wire.Build(
		wire.FieldsOf(&opts, "DBDsn", "Sec"),
		db.InitServerRepo,
		model.NewServerModel,
		excelexporter.NewExporter,
		wire.Struct(new(api.ApiServerOptions), "*"),
		api.InitApiServer,
		wire.Struct(new(Server), "*"),
	)
	return &Server{}, nil
}
