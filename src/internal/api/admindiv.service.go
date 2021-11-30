package api

import (
	"strconv"

	"github.com/aquarius6666/citizen-v/src/internal/db"
	"github.com/aquarius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aquarius6666/citizen-v/src/internal/lib"
	"github.com/aquarius6666/citizen-v/src/internal/var/e"
	"github.com/aquarius6666/citizen-v/src/pb"
	"github.com/aquarius6666/go-utils/utils"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type AdminDivService struct {
	Repo db.ServerRepo
}

func (s *AdminDivService) CreateAdminDiv(req *pb.PostAdminDivRequest) (*pb.PostAdminDivResponse_Data, error) {
	var err error
	if req.Code == "" || req.Name == "" || req.SuperiorId == "" || req.Type == "" {
		return nil, e.ErrMissingBody
	}
	var sid uuid.UUID
	if sid, err = uuid.Parse(req.SuperiorId); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	addiv, err := s.Repo.InsertAdminDiv(&admindiv.AdminDiv{
		Name:       &req.Name,
		Code:       &req.Code,
		Type:       &req.Type,
		SuperiorID: sid,
	})
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
		if sid, err := uuid.Parse(req.SuperiorId); err == nil {
			search.SuperiorID = sid
		}
	}
	if req.Type != "" {
		search.Type = &req.Type
	}
	if req.Id != "" {
		if sid, err := uuid.Parse(req.Id); err == nil {
			search.ID = sid
		}
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
	total, err := s.Repo.CountAdminDiv(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	list, err := s.Repo.ListAdminDiv(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	result := lib.ConvertAdminDiv(list)
	return &pb.GetAdminDivResponse_Data{
		Results: result,
		Pagination: &pb.Pagination{
			Limit:  int32(limit),
			Offset: int32(skip),
			Total:  int32(int(utils.Int64Val(total))),
		},
	}, err
}
