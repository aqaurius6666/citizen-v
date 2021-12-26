package api

import (
	"fmt"
	"time"

	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/gin-gonic/gin"
)

type CitizenController struct {
	Service *CitizenService
}

func (s *CitizenController) HandlePostExport(g *gin.Context) {
	var err error
	req := &pb.PostCitizensExportRequest{
		XCallerId: g.GetString("uid"),
	}
	reader, contentLength, err := s.Service.Export(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	now := time.Now()
	g.DataFromReader(200, contentLength, "application/octet-stream", reader, map[string]string{
		"Content-Disposition": "attachment; filename= " + fmt.Sprintf("%d_%d_%d.xlsx", now.Day(), now.Month(), now.Year()),
	})

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
	adminDivCodes, err := lib.StrToStrArray(g.Query("adminDivCodes"))
	if err != nil {
		lib.BadRequest(g, err)
		return
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
