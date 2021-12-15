package main

import (
	"context"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/cockroach/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/cockroach/citizen"
	"github.com/aqaurius6666/citizen-v/src/internal/db/cockroach/role"
	"github.com/aqaurius6666/citizen-v/src/internal/db/cockroach/campaign"
	"github.com/aqaurius6666/citizen-v/src/internal/db/cockroach/user"
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
	err = seedAdminDiv(service)
	if err != nil {
		logger.Error("seed admin div fail, err:", err)
	}
	err = seedCitizen(service)
	if err != nil {
		logger.Error("seed citizen fail, err:", err)
	}
	err = seedUser(service)
	if err != nil {
		logger.Error("seed user fail, err:", err)
	}
	err = seedCampaign(service)
	if err != nil {
		logger.Error("seed fail, err:", err)
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

func seedAdminDiv(dbase db.ServerRepo) error {
	err := dbase.RawSQL(admindiv.SQL)
	if err != nil {
		return err
	}
	logger.Info("seed admin div successfully")
	return nil
}

func seedCitizen(dbase db.ServerRepo) error {
	err := dbase.RawSQL(citizen.SQL)
	if err != nil {
		return err
	}
	logger.Info("seed citizen successfully")
	return nil
}

func seedCampaign(dbase db.ServerRepo) error {
	err := dbase.RawSQL(campaign.SQL)
	if err != nil {
		return err
	}
	logger.Info("seed campaign successfully")
	return nil
}
func seedUser(dbase db.ServerRepo) error {
	err := dbase.RawSQL(user.SQL)
	if err != nil {
		return err
	}
	logger.Info("seed user successfully")
	return nil
}
