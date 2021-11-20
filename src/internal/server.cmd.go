package main

import (
	"os"
	"time"

	cli2 "github.com/aquarius6666/go-utils/cli"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	serviceName = "main-service"
)

var logger *logrus.Logger

func main() {
	logger = logrus.New()
	if err := makeApp().Run(os.Args); err != nil {
		logger.WithField("err", err).Error("shutting down due to error")
		_ = os.Stderr.Sync()
		os.Exit(1)
	}
}

func makeApp() *cli.App {
	app := &cli.App{
		Name:                 serviceName,
		Version:              "v1.0.1",
		EnableBashCompletion: true,
		Compiled:             time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Vu Nguyen",
				Email: "aqaurius6666@gmail.com",
			},
		},
		Action: runMain,
		Commands: []*cli.Command{
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Usage:   "run server",
				Action:  runMain,
			},
			// {
			// 	Name:    "seed-data",
			// 	Aliases: []string{"sd"},
			// 	Usage:   "seed data",
			// 	Action:  seedData,
			// 	Flags: []cli.Flag{
			// 		&cli.BoolFlag{
			// 			Name:    "clean",
			// 			EnvVars: []string{"CLEAN_DB"},
			// 			Usage:   "Clean DB before seeding",
			// 		},
			// 	},
			// },
			// {
			// 	Name:    "clean",
			// 	Aliases: []string{"c"},
			// 	Usage:   "clean DB",
			// 	Action:  clean,
			// },
		},
		Flags: append([]cli.Flag{
			&cli.StringFlag{
				Name:     "db-uri",
				Required: true,
				EnvVars:  []string{"DB_URI", "DATABASE_URL"},
				Usage:    "The URI for connecting to database (supported URIs: in-memory://, postgresql://auth@host:26257/linkgraph?sslmode=disable)",
			},
		},
			append(cli2.CommonServerFlag,
				cli2.LoggerFlag...)...),
	}
	return app

}
