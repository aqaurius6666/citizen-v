package api

import (
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/gin-gonic/gin"
)

type StatisticController struct {
	Service *StatisticService
}

func (s *StatisticController) HandleGetCitizen(g *gin.Context) {
	var err error
	adminDivCodes, err := lib.StrToStrArray(g.Query("adminDivCodes"))
	if err != nil {
		lib.BadRequest(g, err)
		return
	}

	req := &pb.GetStatisticsCitizensRequest{
		AdminDivCode:  g.Query("adminDivCode"),
		XCallerId:     g.GetString("uid"),
		AdminDivCodes: adminDivCodes,
	}
	res, err := s.Service.GetCitizens(req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
