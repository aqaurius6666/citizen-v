package api

import (
	"github.com/aqaurius6666/boilerplate-server-go/src/internal/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	G          *gin.Engine
	logger     *logrus.Logger
	serverRepo db.ServerRepo
	Index      *IndexController
}

func (s *ApiServer) RegisterEndpoint() {
	s.G.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type", "User-Agent"},
	}))
	s.G.Use(gin.Recovery())
	s.G.Use(gin.Logger())

	s.G.GET("/", s.Index.HandleIndexGet)
}
