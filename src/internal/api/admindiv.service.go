package api

import (
	"strconv"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/campaign"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/lib/validate"
	"github.com/aqaurius6666/citizen-v/src/internal/model"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/aqaurius6666/go-utils/utils"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type AdminDivService struct {
	Repo  db.ServerRepo
	Model model.Server
}

func (s *AdminDivService) UpdateOne(req *pb.PutOneAdminDivRequest) (*pb.PutOneAdminDivResponse_Data, error) {
	var err error
	var sid uuid.UUID
	var search admindiv.Search
	if f, ok := validate.RequiredFields(req, "Id", "Name"); !ok {
		return nil, e.ErrMissingField(f)
	}
	if sid, err = uuid.Parse(req.Id); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	search.ID = sid

	tempAdminDiv := admindiv.AdminDiv{
		Name: &req.Name,
	}
	if err := validate.Validate(tempAdminDiv); err != nil {
		return nil, admindiv.ErrInvalid
	}

	err = s.Repo.UpdateAdminDiv(&search, &tempAdminDiv)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.PutOneAdminDivResponse_Data{}, nil
}

func (s *AdminDivService) CreateAdminDiv(req *pb.PostAdminDivRequest) (*pb.PostAdminDivResponse_Data, error) {
	var err error
	if f, ok := validate.RequiredFields(req, "Name", "XCallerId"); !ok {
		return nil, e.ErrMissingField(f)
	}
	usr, err := s.Model.GetUserById(uuid.MustParse(req.XCallerId))
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	code, err := s.Model.GetNewCode(usr.AdminDivID)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	t, err := s.Model.GetChildType(usr.AdminDivID)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	tempAdminDiv := admindiv.AdminDiv{
		Name:       &req.Name,
		Code:       &code,
		Type:       &t,
		SuperiorID: usr.AdminDivID,
	}
	if err := validate.Validate(tempAdminDiv); err != nil {
		return nil, admindiv.ErrInvalid
	}

	addiv, err := s.Repo.InsertAdminDiv(&tempAdminDiv)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &pb.PostAdminDivResponse_Data{
		Admindiv: &pb.AdminDiv{
			Code:       *addiv.Code,
			Name:       *addiv.Name,
			SuperiorId: addiv.SuperiorID.String(),
			Id:         addiv.ID.String(),
			Type:       *addiv.Type,
		},
	}, nil
}

func (s *AdminDivService) GetOptions(req *pb.GetAdminDivOptionsRequest) (*pb.GetAdminDivOptionsResponse_Data, error) {
	var err error
	var search admindiv.Search
	uid := uuid.Nil
	search.Type = &admindiv.CITY
	if req.SuperiorId != "" {
		uid, err = uuid.Parse(req.SuperiorId)
		if err != nil {
			return nil, e.ErrIdInvalid
		}
		if uid != uuid.Nil {
			search.SuperiorID = uid
			search.Type = nil
		}
	}
	add, err := s.Repo.ListAdminDiv(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &pb.GetAdminDivOptionsResponse_Data{
		AdminDiv: lib.ConvertAdminDivs(add),
	}, nil
}

func (s *AdminDivService) GetAdminDivById(req *pb.GetOneAdminDivRequest) (*pb.GetOneAdminDivResponse_Data, error) {
	var err error
	var search admindiv.Search
	if f, ok := validate.RequiredFields(req, "Id"); !ok {
		return nil, e.ErrMissingField(f)
	}
	if search.ID, err = uuid.Parse(req.Id); err != nil {
		return nil, e.ErrIdInvalid
	}
	add, err := s.Repo.SelectAdminDiv(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &pb.GetOneAdminDivResponse_Data{
		AdminDiv: lib.ConvertOneAdminDiv(add),
	}, nil
}

func (s *AdminDivService) ListAdminDiv(req *pb.GetAdminDivRequest) (*pb.GetAdminDivResponse_Data, error) {
	var err error
	var limit, skip int
	var search admindiv.Search
	if req.Code != "" {
		search.Code = &req.Code
	}
	if req.Name != "" {
		search.Name = &req.Name
	}
	if req.SuperiorId != "" {
		sid, err := uuid.Parse(req.SuperiorId)
		if err != nil {
			return nil, xerrors.Errorf("%w", err)
		}
		if ok, err := s.Model.HasPermission(uuid.MustParse(req.XCallerId), sid); err != nil || !ok {
			return nil, e.ErrAuthNoPermission
		}
		search.SuperiorID = sid
	}
	if req.Type != "" {
		search.Type = &req.Type
	}
	if req.Id != "" {
		sid, err := uuid.Parse(req.Id)
		if err != nil {
			return nil, xerrors.Errorf("%w", err)
		}
		if ok, err := s.Model.HasPermission(uuid.MustParse(req.XCallerId), sid); err != nil || !ok {
			return nil, e.ErrAuthNoPermission
		}
		search.ID = sid
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

	usr, err := s.Model.GetUserById(uuid.MustParse(req.XCallerId))
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	if search.SuperiorID == uuid.Nil && search.ID == uuid.Nil {
		if usr.AdminDivID != uuid.Nil {
			add, err := s.Model.GetAdminDivById(usr.AdminDivID)
			if err != nil {
				return nil, xerrors.Errorf("%w", err)
			}
			search.SuperiorID = add.ID
		} else {
			search.Type = &admindiv.CITY
			search.SuperiorID = uuid.Nil
		}
	}

	search.Fields = []string{"id", "name", "superior_id", "code", "type"}

	total, err := s.Repo.CountAdminDiv(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	list, err := s.Repo.ListAdminDiv(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	results := make([]*pb.GetAdminDivResponse_Data_Results, 0)
	for _, l := range list {
		add := lib.ConvertOneAdminDiv(l)
		camp, _ := s.Repo.TotalCampaignRecord(&campaign.Search{
			Campaign: campaign.Campaign{
				Code: &add.Code,
			},
		})
		percent := utils.Float32Val(camp.Percent)
		records := utils.IntVal(camp.RecordNumber)
		results = append(results, &pb.GetAdminDivResponse_Data_Results{
			AdminDiv: add,
			Campaign: &pb.Campaign{
				Percent: percent,
				Record:  int32(records),
			},
		})
	}
	return &pb.GetAdminDivResponse_Data{
		Results: results,
		Pagination: &pb.Pagination{
			Limit:  int32(limit),
			Offset: int32(skip),
			Total:  int32(int(utils.Int64Val(total))),
		},
	}, err
}
