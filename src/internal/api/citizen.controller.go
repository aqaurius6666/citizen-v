package api

import (
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/gin-gonic/gin"
)

type CitizenController struct {
	Service *CitizenService
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
	req := &pb.GetCitizenRequest{
		Name:     g.Query("name"),
		Id:       g.Query("id"),
		Pid:      g.Query("pid"),
		Birthday: g.Query("birthday"),
		Limit:    g.Query("limit"),
		Offset:   g.Query("offset"),
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
	res, err := s.Service.CreateCitizen(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
