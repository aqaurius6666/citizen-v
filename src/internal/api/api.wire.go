//go:build wireinject
// +build wireinject

package api

import (
	"context"
	"github.com/aquarius6666/citizen-v/src/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

type ApiServerOptions struct {
	MainRepo db.ServerRepo
}

func InitApiServer(ctx context.Context, logger *logrus.Logger, opts ApiServerOptions) (*ApiServer, error) {
	wire.Build(
		wire.FieldsOf(&opts, "MainRepo"),
		gin.New,
		wire.Struct(new(IndexController), "*"),
		wire.Struct(new(LoggerMiddleware), "*"),
		wire.Struct(new(ApiServer), "*"),
	)
	return &ApiServer{}, nil
}
