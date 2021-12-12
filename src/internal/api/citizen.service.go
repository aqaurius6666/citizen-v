package api

import (
	"strconv"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/lib/validate"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/aqaurius6666/go-utils/utils"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type CitizenService struct {
	Repo db.ServerRepo
}

func (s *CitizenService) UpdateOne(req *pb.PutOneCitizenRequest) (*pb.PutOneCitizenResponse_Data, error) {
	var err error
	var sid uuid.UUID
	var search citizen.Search
	if ok := validate.RequiredFields(req,
		"Id", "Name",
		"Birthday", "Gender",
		"Nationality", "FatherName", "FatherPid",
		"MotherName", "MotherPid", "CurrentPlace",
		"JobName", "Pid",
	); !ok {
		return nil, e.ErrMissingBody
	}
	if sid, err = uuid.Parse(req.Id); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	search.ID = sid
	tmpBirthday := uint64(int(req.Birthday))
	tmp := citizen.Citizen{
		Name:         &req.Name,
		Birthday:     &tmpBirthday,
		PID:          &req.Pid,
		Gender:       &req.Gender,
		Nationality:  &req.Nationality,
		FatherName:   &req.FatherName,
		FatherPID:    &req.FatherPid,
		MotherName:   &req.MotherName,
		MotherPID:    &req.MotherPid,
		CurrentPlace: &req.CurrentPlace,
		JobName:      &req.JobName,
	}
	if err := validate.Validate(tmp); err != nil {
		return nil, admindiv.ErrInvalid
	}

	err = s.Repo.UpdateCitizen(&search, &tmp)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &pb.PutOneCitizenResponse_Data{}, nil
}
func (s *CitizenService) GetCitizenById(req *pb.GetOneCitizenRequest) (*pb.GetOneCitizenResponse_Data, error) {
	var err error
	var search citizen.Search
	if req.Id == "" {
		return nil, e.ErrMissingBody
	}
	if search.ID, err = uuid.Parse(req.Id); err != nil {
		return nil, e.ErrIdInvalid
	}
	ctz, err := s.Repo.SelectCitizen(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &pb.GetOneCitizenResponse_Data{
		Citizen: &pb.Citizen{
			Id:           ctz.ID.String(),
			Name:         *ctz.Name,
			Birthday:     int32(*ctz.Birthday),
			Gender:       *ctz.Gender,
			Nationality:  *ctz.Nationality,
			FatherName:   *ctz.Nationality,
			FatherPid:    *ctz.FatherPID,
			MotherName:   *ctz.MotherName,
			MotherPid:    *ctz.MotherPID,
			CurrentPlace: *ctz.CurrentPlace,
			JobName:      *ctz.JobName,
			Pid:          *ctz.PID,
		},
	}, nil
}
func (s *CitizenService) CreateCitizen(req *pb.PostCitizenRequest) (*pb.PostCitizenResponse_Data, error) {
	var err error
	if req.CurrentPlace == "" ||
		req.Name == "" ||
		req.FatherName == "" ||
		req.FatherPid == "" ||
		req.Gender == "" ||
		req.JobName == "" ||
		req.MotherPid == "" ||
		req.MotherName == "" ||
		req.Nationality == "" ||
		req.Pid == "" {
		return nil, e.ErrMissingBody
	}

	tmp := uint64(int(req.Birthday))
	ctz, err := s.Repo.InsertCitizen(&citizen.Citizen{
		Name:         &req.Name,
		Birthday:     &tmp,
		PID:          &req.Pid,
		Gender:       &req.Gender,
		Nationality:  &req.Nationality,
		FatherName:   &req.FatherName,
		FatherPID:    &req.Pid,
		MotherName:   &req.MotherName,
		MotherPID:    &req.MotherPid,
		CurrentPlace: &req.CurrentPlace,
		JobName:      &req.JobName,
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.PostCitizenResponse_Data{
		Citizen: &pb.Citizen{
			Id:           ctz.ID.String(),
			Name:         *ctz.Name,
			Birthday:     int32(*ctz.Birthday),
			Gender:       *ctz.Gender,
			Nationality:  *ctz.Nationality,
			FatherName:   *ctz.Nationality,
			FatherPid:    *ctz.FatherPID,
			MotherName:   *ctz.MotherName,
			MotherPid:    *ctz.MotherPID,
			CurrentPlace: *ctz.CurrentPlace,
			JobName:      *ctz.JobName,
			Pid:          *ctz.PID,
		},
	}, nil
}
func (s *CitizenService) ListCitizen(req *pb.GetCitizenRequest) (*pb.GetCitizenResponse_Data, error) {
	var err error
	var skip, limit int
	var search citizen.Search
	if req.Name != "" {
		search.Name = &req.Name
	}
	if req.Id != "" {
		if sid, err := uuid.Parse(req.Id); err == nil {
			search.ID = sid
		}
	}
	if req.Birthday != "" {
		if birthday, err := strconv.ParseInt(req.Birthday, 10, 64); err != nil {
			temp := uint64(birthday)
			search.Birthday = &temp
		}
	}
	if req.Pid != "" {
		search.PID = &req.Pid
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
	total, err := s.Repo.CountCitizen(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	list, err := s.Repo.ListCitizen(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	result := lib.ConvertCitizens(list)
	return &pb.GetCitizenResponse_Data{
		Results: result,
		Pagination: &pb.Pagination{
			Limit:  int32(limit),
			Offset: int32(skip),
			Total:  int32(int(utils.Int64Val(total))),
		},
	}, err
}
