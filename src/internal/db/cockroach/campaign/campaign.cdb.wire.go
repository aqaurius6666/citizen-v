//go:build wireinject
// +build wireinject

package campaign

import (
	"context"

	"github.com/aqaurius6666/go-utils/database/cockroach"
	"github.com/aqaurius6666/go-utils/logger"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitCampaignCDBRepo(ctx context.Context, logger *logrus.Logger, db *gorm.DB) (*CampaignCDBRepo, error) {
	wire.Build(
		cockroach.InitCDBRepository,
		wire.Struct(new(CampaignCDBRepo), "CDBRepository"),
	)
	return &CampaignCDBRepo{}, nil
}

func InitRoleCDBMockRepo(ctx context.Context, db *gorm.DB) (*CampaignCDBRepo, error) {
	wire.Build(
		logger.InitLoggerWithoutCLIContext,
		cockroach.InitCDBRepository,
		wire.Struct(new(CampaignCDBRepo), "CDBRepository"),
	)
	return &CampaignCDBRepo{}, nil
}
