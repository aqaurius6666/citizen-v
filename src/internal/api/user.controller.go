package api

import (
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *UserService
}

func (s *UserController) HandlePostUnban(g *gin.Context) {
	var err error
	req := &pb.PostUserActiveRequest{
		Id:       g.Param("id"),
		Value:    true,
		CallerId: g.GetString("uid"),
	}
	res, err := s.Service.Active(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *UserController) HandlePostBan(g *gin.Context) {
	var err error
	req := &pb.PostUserActiveRequest{
		Id:       g.Param("id"),
		Value:    false,
		CallerId: g.GetString("uid"),
	}
	res, err := s.Service.Active(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *UserController) HandleGetOne(g *gin.Context) {
	var err error
	req := &pb.GetUserOneRequest{
		Id: g.Param("id"),
	}
	res, err := s.Service.Get(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *UserController) HandleGet(g *gin.Context) {
	var err error
	req := &pb.GetUsersRequest{
		RoleId:     g.Query("roleId"),
		AdminDivId: g.Query("adminDivId"),
		Username:   g.Query("username"),
		Id:         g.Query("id"),
		Limit:      g.Query("limit"),
		Offset:     g.Query("offset"),
	}
	res, err := s.Service.ListUsers(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *UserController) HandlePostIssue(g *gin.Context) {
	var req pb.PostUserIssueRequest
	var err error
	err = lib.GetBody(g, &req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	req.Id = g.GetString("uid")
	res, err := s.Service.Issue(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
