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
	req.CallerId = g.GetString("uid")
	res, err := s.Service.New(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
