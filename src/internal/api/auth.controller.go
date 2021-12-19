package api

import (
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *AuthService
}

func (s *AuthController) HandleGet(g *gin.Context) {

	var err error
	req := &pb.GetAuthRequest{
		CallerId: g.GetString("uid"),
	}

	res, err := s.Service.Auth(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *AuthController) HandlePostRegister(g *gin.Context) {
	var req pb.PostRegisterRequest
	var err error
	err = lib.GetBody(g, &req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	res, err := s.Service.Register(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *AuthController) HandlePostLogin(g *gin.Context) {
	var req pb.PostLoginRequest
	var err error
	err = lib.GetBody(g, &req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	res, err := s.Service.Login(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *AuthController) HandlePostPassword(g *gin.Context) {
	var req pb.PostAuthPasswordRequest
	var err error
	err = lib.GetBody(g, &req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	req.Id = g.GetString("uid")
	res, err := s.Service.ChangePassword(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
