package api

import (
	"github.com/aquarius6666/citizen-v/src/internal/lib"
	"github.com/aquarius6666/citizen-v/src/pb"
	"github.com/gin-gonic/gin"
)

type AdminDivController struct {
	Service *AdminDivService
}

func (s *AdminDivController) HandleGet(g *gin.Context) {
	var err error
	req := &pb.GetAdminDivRequest{
		Code:       g.Query("code"),
		Name:       g.Query("name"),
		SuperiorId: g.Query("superiorId"),
		Type:       g.Query("type"),
		Id:         g.Query("id"),
		Limit:      g.Query("limit"),
		Offset:     g.Query("offset"),
	}
	res, err := s.Service.ListAdminDiv(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *AdminDivController) HandlePost(g *gin.Context) {
	var req pb.PostAdminDivRequest
	var err error
	err = lib.GetBody(g, &req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	res, err := s.Service.CreateAdminDiv(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
