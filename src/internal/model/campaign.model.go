package model

import (
	"fmt"
	"time"

	"github.com/aqaurius6666/citizen-v/src/internal/db/campaign"
	"github.com/aqaurius6666/citizen-v/src/internal/lib/validate"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/aqaurius6666/go-utils/utils"
)

type Campaign interface {
	GetValidCampaign(*string) (*campaign.Campaign, error)
	NewCampaign(*campaign.Campaign) (*campaign.Campaign, error)
}

func (s *ServerModel) GetValidCampaign(code *string) (*campaign.Campaign, error) {
	var err error
	camp := &campaign.Campaign{}
	camp, err = s.Repo.SelectCampaign(&campaign.Search{
		Campaign: campaign.Campaign{
			Code:      code,
			EndTime:   utils.Int64Ptr(time.Now().UnixMilli()),
			StartTime: utils.Int64Ptr(time.Now().UnixMilli()),
		},
	})
	if err != nil {
		return nil, err
	}
	return camp, nil
}

func (s *ServerModel) NewCampaign(camp *campaign.Campaign) (*campaign.Campaign, error) {
	var err error
	if f, ok := validate.RequiredFields(camp, "Code", "EndTime", "StartTime"); !ok {
		return nil, e.ErrMissingField(f)
	}
	existed, err := s.GetValidCampaign(camp.Code)
	if err == nil || existed != nil {
		return nil, e.ErrAdminDivCampaignExisted
	}
	name := fmt.Sprintf("Campaign %s", time.UnixMilli(*camp.StartTime).Format(time.RFC1123))
	camp.Name = &name
	camp, err = s.Repo.InsertCampaign(camp)
	if err != nil {
		return nil, err
	}
	return camp, nil
}
