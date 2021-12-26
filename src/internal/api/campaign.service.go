package api

import (
	"strconv"
	"time"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/campaign"
	"github.com/aqaurius6666/citizen-v/src/internal/db/user"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/lib/validate"
	"github.com/aqaurius6666/citizen-v/src/internal/model"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/aqaurius6666/go-utils/utils"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type CampaignService struct {
	Repo  db.ServerRepo
	Model model.Server
}

func (s *CampaignService) Close(req *pb.PostCampaignDoneRequest) (*pb.PostCampaignDoneResponse_Data, error) {

	if f, ok := validate.RequiredFields(req, "Id", "XCallerId"); !ok {
		return nil, e.ErrMissingField(f)
	}
	var err error
	var uid uuid.UUID
	if uid, err = uuid.Parse(req.Id); err != nil {
		return nil, e.ErrIdInvalid
	}
	var search campaign.Search
	search.ID = uid
	camp, err := s.Repo.SelectCampaign(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	usr, err := s.Model.GetUserById(uuid.MustParse(req.XCallerId))
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	if usr.AdminDivCode != camp.Code {
		return nil, e.ErrAuthNoPermission
	}
	err = s.Repo.UpdateCampaign(&search, &campaign.Campaign{
		IsDone: utils.BoolPtr(true),
	})
	if err != nil {
		return nil, err
	}
	return &pb.PostCampaignDoneResponse_Data{}, nil
}

func (s *CampaignService) List(req *pb.GetCampaignsRequest) (*pb.GetCampaignsResponse_Data, error) {
	var skip, limit int
	if f, ok := validate.RequiredFields(req, "XCallerId"); !ok {
		return nil, e.ErrMissingField(f)
	}
	var search campaign.Search
	usr, err := s.Model.GetUserById(uuid.MustParse(req.XCallerId))
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	if req.StartTime != "" {
		tmp, err := strconv.ParseInt(req.StartTime, 10, 64)
		if err != nil {
			return nil, xerrors.Errorf("%w", err)
		}

		search.StartTime = &tmp
	}
	if req.EndTime != "" {
		tmp, err := strconv.ParseInt(req.EndTime, 10, 64)
		if err != nil {
			return nil, xerrors.Errorf("%w", err)
		}

		search.EndTime = &tmp
	}
	if req.AdminDivCode != "" {
		if ok, err := s.Model.HasPermissionByCode(usr.ID, req.AdminDivCode); err != nil || !ok {
			return nil, e.ErrAuthNoPermission
		}
		search.SuperiorCode = &req.AdminDivCode
	} else {
		search.SuperiorCode = usr.AdminDivCode
	}

	limit = 10
	if req.Limit != "" {
		if o, err := strconv.Atoi(req.Limit); err == nil {
			limit = o
		}
	}
	search.Limit = limit

	skip = 0
	if req.Offset != "" {
		if o, err := strconv.Atoi(req.Offset); err == nil {
			skip = o
		}
	}
	search.Skip = skip
	total, err := s.Repo.CountCampaign(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	camps, err := s.Repo.ListCampaign(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.GetCampaignsResponse_Data{
		Results:    lib.ConvertCampigns(camps, s.Repo),
		Pagination: lib.ConvertPagination(skip, limit, *total),
	}, nil
}

func (s *CampaignService) New(req *pb.PostCampaignRequest) (*pb.PostCampaignResponse_Data, error) {
	if f, ok := validate.RequiredFields(req, "XCallerId", "Codes", "EndTime", "StartTime"); !ok {
		return nil, e.ErrMissingField(f)
	}
	if req.StartTime >= req.EndTime || req.EndTime <= time.Now().UnixMilli() {
		return nil, e.ErrInvalidTime
	}
	usr, err := s.Repo.SelectUser(&user.Search{
		User: user.User{BaseModel: database.BaseModel{
			ID: uuid.MustParse(req.XCallerId),
		}},
	})
	if err != nil {
		return nil, err
	}
	if ok := s.Model.IsChild(*usr.AdminDivCode, req.Codes); !ok {
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
