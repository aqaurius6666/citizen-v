package api

import (
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/gin-gonic/gin"
)

type CampaignController struct {
	Service *CampaignService
}

func (s *CampaignController) HandlePost(g *gin.Context) {
	var req pb.PostCampaignRequest
	var err error
	err = lib.GetBody(g, &req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	req.XCallerId = g.GetString("uid")
	res, err := s.Service.New(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *CampaignController) HandleGet(g *gin.Context) {
	req := &pb.GetCampaignsRequest{
		XCallerId:    g.GetString("uid"),
		StartTime:    g.Query("startTime"),
		EndTime:      g.Query("endTime"),
		Limit:        g.Query("limit"),
		Offset:       g.Query("offset"),
		AdminDivCode: g.Query("adminDivCode"),
	}
	res, err := s.Service.List(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *CampaignController) HandlePostClose(g *gin.Context) {
	req := &pb.PostCampaignDoneRequest{
		XCallerId: g.GetString("uid"),
		Id:        g.Param("id"),
	}
	res, err := s.Service.Close(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
