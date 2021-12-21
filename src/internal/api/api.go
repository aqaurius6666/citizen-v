package api

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/role"
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
	RoleMiddleware   RoleMiddleware
	Auth             *AuthController
	AuthMiddleware   *AuthMiddleware
	AdminDiv         *AdminDivController
	Citizen          *CitizenController
	User             *UserController
	Campaign         *CampaignController
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

	auth := api.Group("/auth")
	auth.POST("/register", s.Auth.HandlePostRegister)
	auth.POST("/login", s.Auth.HandlePostLogin)
	auth.POST("/ping", s.AuthMiddleware.CheckAuth, s.Index.HandleIndexGet)
	auth.POST("/password", s.AuthMiddleware.CheckAuth, s.Auth.HandlePostPassword)
	auth.GET("", s.AuthMiddleware.CheckAuth, s.Auth.HandleGet)

	admindiv := api.Group("/administrative-divisions")
	admindiv.GET("", s.AuthMiddleware.CheckAuth, s.AdminDiv.HandleGet)
	admindiv.GET("/:id", s.AuthMiddleware.CheckAuth, s.AdminDiv.HandleGetOne)
	admindiv.POST("", s.AuthMiddleware.CheckAuth, s.RoleMiddleware.OnlyRole(role.ROLE_ADMIN), s.AdminDiv.HandlePost)
	admindiv.PUT("/:id", s.AuthMiddleware.CheckAuth, s.RoleMiddleware.OnlyRole(role.ROLE_ADMIN), s.AdminDiv.HandlePutOne)
	admindiv.GET("/options", s.AuthMiddleware.CheckAuth, s.AdminDiv.HandleGetOptions)

	citizen := api.Group("/citizens")
	citizen.GET("", s.AuthMiddleware.CheckAuth, s.Citizen.HandleGet)
	citizen.GET("/:id", s.AuthMiddleware.CheckAuth, s.Citizen.HandleGetById)
	citizen.POST("", s.AuthMiddleware.CheckAuth, s.RoleMiddleware.OnlyActive(), s.RoleMiddleware.OnlyRole(role.ROLE_B1, role.ROLE_B2), s.Citizen.HandlePost)
	citizen.PUT("/:id", s.AuthMiddleware.CheckAuth, s.RoleMiddleware.OnlyActive(), s.RoleMiddleware.OnlyRole(role.ROLE_B1), s.Citizen.HandlePutOne)
	citizen.DELETE("/:id", s.AuthMiddleware.CheckAuth, s.RoleMiddleware.OnlyActive(), s.RoleMiddleware.OnlyRole(role.ROLE_B1), s.Citizen.HandleDeleteById)

	user := api.Group("/users")
	user.GET("", s.AuthMiddleware.CheckAuth, s.User.HandleGet)
	user.GET("/:id", s.AuthMiddleware.CheckAuth, s.User.HandleGetOne)
	user.POST("/issue", s.AuthMiddleware.CheckAuth, s.RoleMiddleware.OnlyActive(), s.User.HandlePostIssue)
	user.POST("/:id/ban", s.AuthMiddleware.CheckAuth, s.RoleMiddleware.OnlyActive(), s.User.HandlePostBan)
	user.POST("/:id/unban", s.AuthMiddleware.CheckAuth, s.RoleMiddleware.OnlyActive(), s.User.HandlePostUnban)

	campaign := api.Group("/campaigns")
	campaign.POST("", s.AuthMiddleware.CheckAuth, s.RoleMiddleware.OnlyActive(), s.RoleMiddleware.OnlyRole(role.ROLE_A1, role.ROLE_A2, role.ROLE_A3), s.Campaign.HandlePost)

}
