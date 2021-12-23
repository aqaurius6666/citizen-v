package api

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/lib/validate"
	"github.com/aqaurius6666/citizen-v/src/internal/model"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type StatisticService struct {
	Repo  db.ServerRepo
	Model model.Server
}

func (s *StatisticService) GetCitizens(req *pb.GetStatisticsCitizensRequest) (*pb.GetStatisticsCitizensResponse_Data, error) {
	var err error
	if f, ok := validate.RequiredFields(req, "XCallerId"); !ok {
		return nil, e.ErrMissingField(f)
	}
	usr, err := s.Model.GetUserById(uuid.MustParse(req.XCallerId))
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	var search citizen.Search
	if req.AdminDivCode != "" {
		if ok, err := s.Model.HasPermissionByCode(usr.ID, req.AdminDivCode); err != nil || !ok {
			return nil, e.ErrAuthNoPermission
		}
		search.AdminDivCode = &req.AdminDivCode
	}

	if req.AdminDivCodes != nil {
		for _, c := range req.AdminDivCodes {
			if ok, err := s.Model.HasPermissionByCode(uuid.MustParse(req.XCallerId), c); err != nil || !ok {
				return nil, e.ErrAuthNoPermission
			}
		}
		search.ArrayCode = req.AdminDivCodes
	}
	if req.AdminDivCode == "" && req.AdminDivCodes == nil {
		code := ""
		if usr.AdminDivID != uuid.Nil {
			add, err := s.Model.GetAdminDivById(usr.AdminDivID)
			if err != nil {
				return nil, xerrors.Errorf("%w", err)
			}
			code = *add.Code
		}
		search.AdminDivCode = &code
	}
	ctz, err := s.Repo.ListCitizen(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.GetStatisticsCitizensResponse_Data{
		Results: lib.ConvertRecords(ctz),
	}, nil
}
