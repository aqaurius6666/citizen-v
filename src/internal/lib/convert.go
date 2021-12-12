package lib

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/aqaurius6666/go-utils/utils"
)

func ConvertAdminDivs(db []*admindiv.AdminDiv) []*pb.AdminDiv {
	r := make([]*pb.AdminDiv, 0)
	for _, s := range db {
		r = append(r, ConvertOneAdminDiv(s))
	}
	return r
}

func ConvertOneAdminDiv(s *admindiv.AdminDiv) *pb.AdminDiv {
	return &pb.AdminDiv{
		Code:       utils.StrVal(s.Code),
		Name:       utils.StrVal(s.Name),
		SuperiorId: s.SuperiorID.String(),
		Type:       utils.StrVal(s.Type),
		Id:         s.ID.String(),
		Subdiv:     ConvertAdminDivs(s.SubDiv),
	}
}

func ConvertOneCitizen(s *citizen.Citizen) *pb.Citizen {
	return &pb.Citizen{
		Name:         utils.StrVal(s.Name),
		Id:           s.ID.String(),
		Birthday:     int32(*s.Birthday),
		Gender:       utils.StrVal(s.Gender),
		Nationality:  utils.StrVal(s.Nationality),
		FatherName:   utils.StrVal(s.FatherName),
		FatherPid:    utils.StrVal(s.FatherPID),
		MotherName:   utils.StrVal(s.MotherName),
		MotherPid:    utils.StrVal(s.MotherPID),
		CurrentPlace: utils.StrVal(s.CurrentPlace),
		JobName:      utils.StrVal(s.JobName),
		Pid:          utils.StrVal(s.PID),
	}
}

func ConvertCitizens(db []*citizen.Citizen) []*pb.Citizen {
	r := make([]*pb.Citizen, 0)
	for _, s := range db {
		r = append(r, ConvertOneCitizen(s))
	}
	return r
}
