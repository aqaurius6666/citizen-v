//go:build wireinject
// +build wireinject

package admindiv

import (
	"context"

	"github.com/aqaurius6666/go-utils/database/cockroach"
	"github.com/aqaurius6666/go-utils/logger"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitAdminDivCDBRepo(ctx context.Context, logger *logrus.Logger, db *gorm.DB) (*AdminDivCDBRepo, error) {
	wire.Build(
		cockroach.InitCDBRepository,
		wire.Struct(new(AdminDivCDBRepo), "CDBRepository"),
	)
	return &AdminDivCDBRepo{}, nil
}

func InitAdminDivCDBMockRepo(ctx context.Context, db *gorm.DB) (*AdminDivCDBRepo, error) {
	wire.Build(
		logger.InitLoggerWithoutCLIContext,
		cockroach.InitCDBRepository,
		wire.Struct(new(AdminDivCDBRepo), "CDBRepository"),
	)
	return &AdminDivCDBRepo{}, nil
}
