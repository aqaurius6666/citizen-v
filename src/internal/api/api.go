package api

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	G                *gin.Engine
	logger           *logrus.Logger
	serverRepo       db.ServerRepo
	Index            *IndexController
	LoggerMiddleware LoggerMiddleware
	Auth             *AuthController
	AuthMiddleware   *AuthMiddleware
	AdminDiv         *AdminDivController
	Citizen          *CitizenController
}

func (s *ApiServer) RegisterEndpoint() {
	gin.SetMode("debug")
	s.G.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type", "User-Agent", "Refer"},
	}))
	s.G.Use(gin.Recovery())
	s.G.Use(s.LoggerMiddleware.Logger())

	s.G.GET("/", s.Index.HandleIndexGet)
	s.G.GET("/random", s.Index.HandleRandomGet)

	api := s.G.Group("/api")

	// Auth group
	auth := api.Group("/auth")
	auth.POST("/register", s.Auth.HandlePostRegister)
	auth.POST("/login", s.Auth.HandlePostLogin)
	auth.POST("/ping", s.AuthMiddleware.CheckAuth, s.Index.HandleIndexGet)

	// Administrative division group
	admindiv := api.Group("/administrative-divisions")
	admindiv.GET("", s.AdminDiv.HandleGet)
	admindiv.GET("/:id", s.AdminDiv.HandleGetOne)
	admindiv.POST("", s.AdminDiv.HandlePost)
	admindiv.PUT("/:id", s.AdminDiv.HandlePutOne)

	// Citizen group
	citizen := api.Group("/citizens")
	citizen.GET("", s.Citizen.HandleGet)
	citizen.GET("/:id", s.Citizen.HandleGetById)
	citizen.POST("", s.Citizen.HandlePost)
	citizen.PUT("/:id", s.Citizen.HandlePutOne)

}
