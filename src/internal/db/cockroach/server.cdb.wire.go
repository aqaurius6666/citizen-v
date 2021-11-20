//go:build wireinject
// +build wireinject

package cockroach

import (
	"context"

	usercdb "github.com/aqaurius6666/boilerplate-server-go/src/internal/db/cockroach/user"
	"github.com/aqaurius6666/boilerplate-server-go/src/internal/db/user"
	"github.com/aquarius6666/go-utils/database/cockroach"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ServerCDBOptions struct {
	cfg *gorm.Config
	dsn string
}

func initServerCDBRepo(ctx context.Context, logger *logrus.Logger, opts ServerCDBOptions) (*ServerCDBRepo, error) {
	wire.Build(
		wire.FieldsOf(&opts, "cfg", "dsn"),
		cockroach.NewCDBConnection,
		cockroach.InitCDBRepository,
		usercdb.InitUserCDBRepo,
		wire.Struct(new(ServerCDBRepo), "*"),
	)
	return &ServerCDBRepo{}, nil
}

func InitServerCDBRepo(ctx context.Context, logger *logrus.Logger, opts ServerCDBOptions) (*ServerCDBRepo, error) {
	s, err := initServerCDBRepo(ctx, logger, opts)
	if err != nil {
		return nil, err
	}
	s.SetInterfaces(&user.User{})
	return s, nil
}
