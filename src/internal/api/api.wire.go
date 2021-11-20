//go:build wireinject
// +build wireinject

package api

import (
	"context"
	"github.com/aqaurius6666/boilerplate-server-go/src/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

type ApiServerOptions struct {
	mainRepo db.ServerRepo
}

func InitApiServer(ctx context.Context, logger *logrus.Logger, opts ApiServerOptions) (*ApiServer, error) {
	wire.Build(
		wire.FieldsOf(&opts, "mainRepo"),
		gin.New,
		wire.Struct(new(ApiServer), "*"),
	)
	return &ApiServer{}, nil
}
