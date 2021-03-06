//go:build wireinject
// +build wireinject

package api

import (
	"context"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/model"
	"github.com/aqaurius6666/citizen-v/src/internal/services/excelexporter"
	"github.com/aqaurius6666/citizen-v/src/internal/services/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

type ApiServerOptions struct {
	MainRepo db.ServerRepo
	Sec      jwt.SecretKey
	Model    model.Server
	Exporter excelexporter.Exporter
}

func InitApiServer(ctx context.Context, logger *logrus.Logger, opts ApiServerOptions) (*ApiServer, error) {
	wire.Build(
		wire.FieldsOf(&opts, "MainRepo", "Sec", "Model", "Exporter"),
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
		wire.Struct(new(CampaignService), "*"),
		wire.Struct(new(CampaignController), "*"),
		wire.Struct(new(StatisticService), "*"),
		wire.Struct(new(StatisticController), "*"),

		wire.Struct(new(ApiServer), "*"),
	)
	return &ApiServer{}, nil
}
