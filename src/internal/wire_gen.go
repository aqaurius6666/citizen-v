// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/aqaurius6666/citizen-v/src/internal/api"
	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/model"
	"github.com/aqaurius6666/citizen-v/src/internal/services/excelexporter"
	"github.com/aqaurius6666/citizen-v/src/internal/services/jwt"
	"github.com/sirupsen/logrus"
)

// Injectors from server.wire.go:

func InitMainServer(ctx context.Context, logger2 *logrus.Logger, opts ServerOptions) (*Server, error) {
	dbDsn := opts.DBDsn
	serverRepo, err := db.InitServerRepo(ctx, logger2, dbDsn)
	if err != nil {
		return nil, err
	}
	secretKey := opts.Sec
	server, err := model.NewServerModel(ctx, logger2, serverRepo)
	if err != nil {
		return nil, err
	}
	exporter := excelexporter.NewExporter()
	apiServerOptions := api.ApiServerOptions{
		MainRepo: serverRepo,
		Sec:      secretKey,
		Model:    server,
		Exporter: exporter,
	}
	apiServer, err := api.InitApiServer(ctx, logger2, apiServerOptions)
	if err != nil {
		return nil, err
	}
	mainServer := &Server{
		ApiServer: apiServer,
		MainRepo:  serverRepo,
	}
	return mainServer, nil
}

// server.wire.go:

type Server struct {
	ApiServer *api.ApiServer
	MainRepo  db.ServerRepo
}

type ServerOptions struct {
	DBDsn db.DBDsn
	Sec   jwt.SecretKey
}
