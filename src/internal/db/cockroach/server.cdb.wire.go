//go:build wireinject
// +build wireinject

package cockroach

import (
	"context"

	"github.com/aquarius6666/citizen-v/src/internal/db/admindiv"
	admindivcdb "github.com/aquarius6666/citizen-v/src/internal/db/cockroach/admindiv"
	rolecdb "github.com/aquarius6666/citizen-v/src/internal/db/cockroach/role"
	usercdb "github.com/aquarius6666/citizen-v/src/internal/db/cockroach/user"
	"github.com/aquarius6666/citizen-v/src/internal/db/role"
	"github.com/aquarius6666/citizen-v/src/internal/db/user"
	"github.com/aquarius6666/go-utils/database/cockroach"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ServerCDBOptions struct {
	Cfg *gorm.Config
	Dsn string
}

func initServerCDBRepo(ctx context.Context, logger *logrus.Logger, opts ServerCDBOptions) (*ServerCDBRepo, error) {
	wire.Build(
		wire.FieldsOf(&opts, "Cfg", "Dsn"),
		cockroach.NewCDBConnection,
		cockroach.InitCDBRepository,
		usercdb.InitUserCDBRepo,
		rolecdb.InitRoleCDBRepo,
		admindivcdb.InitAdminDivCDBRepo,
		wire.Struct(new(ServerCDBRepo), "*"),
	)
	return &ServerCDBRepo{}, nil
}

func InitServerCDBRepo(ctx context.Context, logger *logrus.Logger, opts ServerCDBOptions) (*ServerCDBRepo, error) {
	s, err := initServerCDBRepo(ctx, logger, opts)
	if err != nil {
		return nil, err
	}
	s.SetInterfaces(&user.User{}, &role.Role{}, &admindiv.AdminDiv{})
	return s, nil
}
