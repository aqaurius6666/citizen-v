//go:build wireinject
// +build wireinject

package api

import (
	"context"

	"github.com/aquarius6666/citizen-v/src/internal/db"
	"github.com/aquarius6666/citizen-v/src/internal/services/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

type ApiServerOptions struct {
	MainRepo db.ServerRepo
	Sec      jwt.SecretKey
}

func InitApiServer(ctx context.Context, logger *logrus.Logger, opts ApiServerOptions) (*ApiServer, error) {
	wire.Build(
		wire.FieldsOf(&opts, "MainRepo", "Sec"),
		gin.New,
		jwt.NewJWT,
		wire.Struct(new(IndexController), "*"),
		wire.Struct(new(AuthMiddleware), "*"),
		wire.Struct(new(AuthController), "*"),
		wire.Struct(new(AdminDivController), "*"),
		wire.Struct(new(AdminDivService), "*"),
		wire.Struct(new(AuthService), "*"),
		wire.Struct(new(LoggerMiddleware), "*"),
		wire.Struct(new(ApiServer), "*"),
	)
	return &ApiServer{}, nil
}
