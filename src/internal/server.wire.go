package main

import (
	"context"

	"github.com/aqaurius6666/boilerplate-server-go/src/internal/api"
	"github.com/sirupsen/logrus"
)

type Server struct {
	ApiServer *api.ApiServer
}

type ServerOptions struct {
	apiOptions api.ApiServerOptions
}

func InitMainServer(ctx context.Context, logger *logrus.Logger, opts ServerOptions) {
}
