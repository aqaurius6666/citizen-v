package api

import (
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *UserService
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