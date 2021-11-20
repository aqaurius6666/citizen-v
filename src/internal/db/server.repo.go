package db

import (
	"context"
	"net/url"

	"github.com/aqaurius6666/boilerplate-server-go/src/internal/db/cockroach"
	"github.com/aqaurius6666/boilerplate-server-go/src/internal/db/user"
	"github.com/aquarius6666/go-utils/database"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type ServerRepo interface {
	database.CommonRepository
	user.UserRepo
}
type DBDsn string

func InitServerRepo(ctx context.Context, logger *logrus.Logger, dsn DBDsn) (ServerRepo, error) {
	uri, err := url.Parse(string(dsn))
	if err != nil {
		return nil, xerrors.Errorf("could not parse DB URI: %w", err)
	}

	switch uri.Scheme {
	case "in-memory":
		logger.Info("using in-memory graph")
		return nil, xerrors.Errorf("Not implemented!", err)
	case "postgresql":
		return cockroach.InitServerCDBRepo(ctx, logger, cockroach.ServerCDBOptions{
			Cfg: &gorm.Config{},
			Dsn: string(dsn),
		})
	case "postgres":
		return cockroach.InitServerCDBRepo(ctx, logger, cockroach.ServerCDBOptions{
			Cfg: &gorm.Config{},
			Dsn: string(dsn),
		})
	default:
		return nil, xerrors.Errorf("unsupported DB URI scheme: %q", uri.Scheme)
	}
}
