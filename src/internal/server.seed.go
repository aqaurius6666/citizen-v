package main

import (
	"context"

	"github.com/aquarius6666/citizen-v/src/internal/db"
	"github.com/aquarius6666/citizen-v/src/internal/db/cockroach/role"
	"github.com/urfave/cli/v2"
)

func clean(appCtx *cli.Context) error {
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()
	service, err := db.InitServerRepo(ctx, logger, db.DBDsn(appCtx.String("db-uri")))
	if err != nil {
		return err
	}
	defer func(db db.ServerRepo) {
		e := db.Close()
		if e != nil {
			panic("cannot close DB")
		}
	}(service)
	if err != nil {
		return err
	}
	return service.Drop()
}

func seedData(appCtx *cli.Context) error {
	logger.Info("Seed starting!")
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	service, err := db.InitServerRepo(ctx, logger, db.DBDsn(appCtx.String("db-uri")))
	if err != nil {
		return err
	}
	defer func(db db.ServerRepo) {
		e := db.Close()
		if e != nil {
			panic("cannot close DB")
		}
	}(service)
	if appCtx.Bool("clean") {
		logger.Info("start cleaning DB")
		err := clean(appCtx)
		if err != nil {
			logger.Error("failed cleaning DB")
			return err
		}
		logger.Info("sucessed cleaning DB")
	}
	err = service.Migrate()
	if err != nil {
		return err
	}
	err = seedRole(service)
	if err != nil {
		logger.Error("seed seedRole fail, err:", err)
	}

	return nil
}

func seedRole(dbase db.ServerRepo) error {
	err := dbase.RawSQL(role.SQL)
	if err != nil {
		return err
	}
	logger.Info("seed role successfully")
	return nil
}
