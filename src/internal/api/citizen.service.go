package api

import (
	"strconv"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/lib/validate"
	"github.com/aqaurius6666/citizen-v/src/internal/model"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/aqaurius6666/go-utils/utils"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type CitizenService struct {
	Repo  db.ServerRepo
	Model model.Server
}

func (s *CitizenService) Delete(req *pb.DeleteCitizenRequest) (*pb.DeleteCitizenResponse_Data, error) {
	var err error
	var sid uuid.UUID
	var search citizen.Search
	if f, ok := validate.RequiredFields(req, "Id", "CallerId"); !ok {
		return nil, e.ErrMissingField(f)
	}
	if sid, err = uuid.Parse(req.Id); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	search.ID = sid
	err = s.Repo.DeleteCitizen(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &pb.DeleteCitizenResponse_Data{}, nil
}

func (s *CitizenService) UpdateOne(req *pb.PutOneCitizenRequest) (*pb.PutOneCitizenResponse_Data, error) {
	var err error
	var sid uuid.UUID
	var search citizen.Search
	if f, ok := validate.RequiredFields(req,
		"Id", "Name",
		"Birthday", "Gender",
		"Nationality", "FatherName", "FatherPid",
		"MotherName", "MotherPid", "CurrentPlace",
		"JobName", "Pid", "CallerId", "ResidencePlace",
		"Hometown", "Religion", "EducationalLevel",
		"AdminDivCode",
	); !ok {
		return nil, e.ErrMissingField(f)
	}
	if sid, err = uuid.Parse(req.Id); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	add, err := s.Repo.SelectAdminDiv(&admindiv.Search{
		AdminDiv: admindiv.AdminDiv{
			Code: utils.StrPtr(req.AdminDivCode),
		},
	})
	if err != nil || add == nil {
		return nil, e.ErrBodyInvalid
	}
	if ok, err := s.Model.HasPermission(uuid.MustParse(req.CallerId), add.ID); err != nil || !ok {
		return nil, e.ErrAuthNoPermission
	}
	search.ID = sid
	tmpBirthday := uint64(int(req.Birthday))
	tmp := citizen.Citizen{
		Name:             &req.Name,
		Birthday:         &tmpBirthday,
		PID:              &req.Pid,
		Gender:           &req.Gender,
		Nationality:      &req.Nationality,
		FatherName:       &req.FatherName,
		FatherPID:        &req.FatherPid,
		MotherName:       &req.MotherName,
		MotherPID:        &req.MotherPid,
		CurrentPlace:     &req.CurrentPlace,
		JobName:          &req.JobName,
		ResidencePlace:   &req.ResidencePlace,
		Hometown:         &req.Hometown,
		Religion:         &req.Religion,
		EducationalLevel: &req.EducationalLevel,
		AdminDivCode:     &req.AdminDivCode,
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
		Citizen: lib.ConvertOneCitizen(ctz),
	}, nil
}
func (s *CitizenService) CreateCitizen(req *pb.PostCitizenRequest) (*pb.PostCitizenResponse_Data, error) {
	var err error
	if f, ok := validate.RequiredFields(req,
		"Pid", "Name", "Birthday", "Gender",
		"Nationality", "MotherName", "MotherPid",
		"FatherName", "FatherPid", "CurrentPlace",
		"JobName", "ResidencePlace", "Hometown",
		"Religion", "EducationalLevel", "CallerId"); !ok {
		return nil, e.ErrMissingField(f)
	}
	add, err := s.Repo.SelectAdminDiv(&admindiv.Search{
		AdminDiv: admindiv.AdminDiv{
			Code: utils.StrPtr(req.AdminDivCode),
		},
	})
	if err != nil || add == nil {
		return nil, e.ErrBodyInvalid
	}
	if ok, err := s.Model.HasPermission(uuid.MustParse(req.CallerId), add.ID); err != nil || !ok {
		return nil, e.ErrAuthNoPermission
	}
	tmp := uint64(int(req.Birthday))
	tmpCitizen := citizen.Citizen{
		Name:             &req.Name,
		Birthday:         &tmp,
		PID:              &req.Pid,
		Gender:           &req.Gender,
		Nationality:      &req.Nationality,
		FatherName:       &req.FatherName,
		FatherPID:        &req.Pid,
		MotherName:       &req.MotherName,
		MotherPID:        &req.MotherPid,
		CurrentPlace:     &req.CurrentPlace,
		JobName:          &req.JobName,
		ResidencePlace:   &req.ResidencePlace,
		Hometown:         &req.Hometown,
		Religion:         &req.Religion,
		EducationalLevel: &req.EducationalLevel,
		AdminDivCode:     &req.AdminDivCode,
	}
	if err := validate.Validate(tmpCitizen); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	ctz, err := s.Repo.InsertCitizen(&tmpCitizen)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.PostCitizenResponse_Data{
		Citizen: lib.ConvertOneCitizen(ctz),
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
