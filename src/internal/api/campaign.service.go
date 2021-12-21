package api

import (
	"time"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/campaign"
	"github.com/aqaurius6666/citizen-v/src/internal/db/user"
	"github.com/aqaurius6666/citizen-v/src/internal/lib/validate"
	"github.com/aqaurius6666/citizen-v/src/internal/model"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/google/uuid"
)

type CampaignService struct {
	Repo  db.ServerRepo
	Model model.Server
}

func (s *CampaignService) New(req *pb.PostCampaignRequest) (*pb.PostCampaignResponse_Data, error) {
	if f, ok := validate.RequiredFields(req, "CallerId", "Codes", "EndTime", "StartTime"); !ok {
		return nil, e.ErrMissingField(f)
	}
	if req.StartTime >= req.EndTime || req.EndTime <= time.Now().UnixMilli() {
		return nil, e.ErrInvalidTime
	}
	usr, err := s.Repo.SelectUser(&user.Search{
		User: user.User{BaseModel: database.BaseModel{
			ID: uuid.MustParse(req.CallerId),
		}},
	})
	if err != nil {
		return nil, err
	}
	if ok := s.Model.IsChild(*usr.AdminDiv.Code, req.Codes); !ok {
		return nil, e.ErrInvalidCodes
	}
	results := make([]*pb.Campaign, 0)
	for _, c := range req.Codes {
		camp, err := s.Model.NewCampaign(&campaign.Campaign{
			Code:      &c,
			EndTime:   &req.EndTime,
			StartTime: &req.StartTime,
		})
		if err != nil {
			return nil, err
		}
		results = append(results, &pb.Campaign{
			Id:        camp.ID.String(),
			Name:      *camp.Name,
			StartTime: *camp.StartTime,
			EndTime:   *camp.EndTime,
		})
	}
	return &pb.PostCampaignResponse_Data{
		Campaign: results,
	}, nil
}
