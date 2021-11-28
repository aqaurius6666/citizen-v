//go:build wireinject
// +build wireinject

package citizen

import (
	"context"

	"github.com/aquarius6666/go-utils/database/cockroach"
	"github.com/aquarius6666/go-utils/logger"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitCitizenCDBRepo(ctx context.Context, logger *logrus.Logger, db *gorm.DB) (*CitizenCDBRepo, error) {
	wire.Build(
		cockroach.InitCDBRepository,
		wire.Struct(new(CitizenCDBRepo), "CDBRepository"),
	)
	return &CitizenCDBRepo{}, nil
}

func InitRoleCDBMockRepo(ctx context.Context, db *gorm.DB) (*CitizenCDBRepo, error) {
	wire.Build(
		logger.InitLoggerWithoutCLIContext,
		cockroach.InitCDBRepository,
		wire.Struct(new(CitizenCDBRepo), "CDBRepository"),
	)
	return &CitizenCDBRepo{}, nil
}
