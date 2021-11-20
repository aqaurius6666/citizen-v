//go:build wireinject
// +build wireinject

package user

import (
	"context"

	"github.com/aquarius6666/go-utils/database/cockroach"
	"github.com/aquarius6666/go-utils/logger"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitUserCDBRepo(ctx context.Context, logger *logrus.Logger, db *gorm.DB) (*UserCDBRepo, error) {
	wire.Build(
		cockroach.InitCDBRepository,
		wire.Struct(new(UserCDBRepo), "CDBRepository"),
	)
	return &UserCDBRepo{}, nil
}

func InitUserCDBMockRepo(ctx context.Context, db *gorm.DB) (*UserCDBRepo, error) {
	wire.Build(
		logger.InitLoggerWithoutCLIContext,
		cockroach.InitCDBRepository,
		wire.Struct(new(UserCDBRepo), "CDBRepository"),
	)
	return &UserCDBRepo{}, nil
}
