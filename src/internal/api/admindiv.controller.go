package api

import (
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/gin-gonic/gin"
)

type AdminDivController struct {
	Service *AdminDivService
}

func (s *AdminDivController) HandleGetOne(g *gin.Context) {
	var err error
	req := &pb.GetOneAdminDivRequest{
		Id: g.Param("id"),
	}
	res, err := s.Service.GetAdminDivById(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
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
		XCallerId:  g.GetString("uid"),
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
	req.XCallerId = g.GetString("uid")
	res, err := s.Service.CreateAdminDiv(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *AdminDivController) HandleGetOptions(g *gin.Context) {
	var err error
	req := &pb.GetAdminDivOptionsRequest{
		SuperiorId:   g.Query("superiorId"),
		XCallerId:    g.GetString("uid"),
		SuperiorCode: g.Query("superiorCode"),
	}
	res, err := s.Service.GetOptions(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *AdminDivController) HandlePutOne(g *gin.Context) {
	var req pb.PutOneAdminDivRequest
	var err error
	err = lib.GetBody(g, &req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	req.Id = g.Param("id")
	res, err := s.Service.UpdateOne(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *AdminDivController) HandleGetName(g *gin.Context) {
	req := pb.GetAdminDivNameRequest{
		XCallerId:    g.GetString("uid"),
		AdminDivId:   g.Query("adminDivId"),
		AdminDivCode: g.Query("adminDivCode"),
	}
	res, err := s.Service.GetName(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
