package api

import (
	"encoding/json"

	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/gin-gonic/gin"
)

type CitizenController struct {
	Service *CitizenService
}

func (s *CitizenController) HandleDeleteById(g *gin.Context) {
	var err error
	req := &pb.DeleteCitizenRequest{
		Id:        g.Param("id"),
		XCallerId: g.GetString("uid"),
	}
	res, err := s.Service.Delete(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *CitizenController) HandlePutOne(g *gin.Context) {
	var err error
	var req pb.PutOneCitizenRequest
	err = lib.GetBody(g, &req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	req.Id = g.Param("id")
	req.XCallerId = g.GetString("uid")
	res, err := s.Service.UpdateOne(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *CitizenController) HandleGetById(g *gin.Context) {
	var err error
	req := &pb.GetOneCitizenRequest{
		Id: g.Param("id"),
	}
	res, err := s.Service.GetCitizenById(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *CitizenController) HandleGet(g *gin.Context) {
	var err error
	adminDivCodesStr := g.Query("adminDivCodes")
	var adminDivCodes []string
	if adminDivCodesStr != "" {
		err = json.Unmarshal([]byte(adminDivCodesStr), &adminDivCodes)
		if err != nil {
			lib.BadRequest(g, err)
			return
		}
	}
	req := &pb.GetCitizenRequest{
		Name:          g.Query("name"),
		Id:            g.Query("id"),
		Pid:           g.Query("pid"),
		Birthday:      g.Query("birthday"),
		Limit:         g.Query("limit"),
		Offset:        g.Query("offset"),
		AdminDivId:    g.Query("adminDivId"),
		AdminDivCode:  g.Query("adminDivCode"),
		XCallerId:     g.GetString("uid"),
		AdminDivCodes: adminDivCodes,
	}

	res, err := s.Service.ListCitizen(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *CitizenController) HandlePost(g *gin.Context) {
	var req pb.PostCitizenRequest
	var err error
	err = lib.GetBody(g, &req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	req.XCallerId = g.GetString("uid")
	res, err := s.Service.CreateCitizen(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
