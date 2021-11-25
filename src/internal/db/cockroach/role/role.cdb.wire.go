//go:build wireinject
// +build wireinject

package role

import (
	"context"

	"github.com/aquarius6666/go-utils/database/cockroach"
	"github.com/aquarius6666/go-utils/logger"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitRoleCDBRepo(ctx context.Context, logger *logrus.Logger, db *gorm.DB) (*RoleCDBRepo, error) {
	wire.Build(
		cockroach.InitCDBRepository,
		wire.Struct(new(RoleCDBRepo), "CDBRepository"),
	)
	return &RoleCDBRepo{}, nil
}

func InitRoleCDBMockRepo(ctx context.Context, db *gorm.DB) (*RoleCDBRepo, error) {
	wire.Build(
		logger.InitLoggerWithoutCLIContext,
		cockroach.InitCDBRepository,
		wire.Struct(new(RoleCDBRepo), "CDBRepository"),
	)
	return &RoleCDBRepo{}, nil
}
