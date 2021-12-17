//go:build wireinject
// +build wireinject

package api

import (
	"context"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/model"
	"github.com/aqaurius6666/citizen-v/src/internal/services/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

type ApiServerOptions struct {
	MainRepo db.ServerRepo
	Sec      jwt.SecretKey
	Model    model.Server
}

func InitApiServer(ctx context.Context, logger *logrus.Logger, opts ApiServerOptions) (*ApiServer, error) {
	wire.Build(
		wire.FieldsOf(&opts, "MainRepo", "Sec", "Model"),
		gin.New,
		jwt.NewJWT,
		wire.Struct(new(IndexController), "*"),
		wire.Struct(new(AuthMiddleware), "*"),
		wire.Struct(new(AuthController), "*"),
		wire.Struct(new(AdminDivController), "*"),
		wire.Struct(new(CitizenController), "*"),
		wire.Struct(new(AdminDivService), "*"),
		wire.Struct(new(UserService), "*"),
		wire.Struct(new(UserController), "*"),
		wire.Struct(new(CitizenService), "*"),
		wire.Struct(new(RoleMiddleware), "*"),
		wire.Struct(new(AuthService), "*"),
		wire.Struct(new(LoggerMiddleware), "*"),
		wire.Struct(new(ApiServer), "*"),
	)
	return &ApiServer{}, nil
}
