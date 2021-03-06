package lib

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/campaign"
	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
	"github.com/aqaurius6666/citizen-v/src/internal/db/role"
	"github.com/aqaurius6666/citizen-v/src/internal/db/user"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/aqaurius6666/go-utils/utils"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
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

func GetAdminDivFullName(id uuid.UUID, repo db.ServerRepo) (*string, error) {
	var name string
	count := 0
	tmp := id
	for tmp != uuid.Nil {
		add, err := repo.SelectAdminDiv(&admindiv.Search{
			DefaultSearchModel: database.DefaultSearchModel{
				Fields: []string{"name", "superior_id"},
			},
			AdminDiv: admindiv.AdminDiv{
				BaseModel: database.BaseModel{
					ID: tmp,
				},
			},
		})
		if err != nil {
			return nil, err
		}
		name = fmt.Sprintf("%s$$%s", *add.Name, name)
		tmp = add.SuperiorID
		count++
	}
	name = strings.Replace(name, "$$", ", ", count-1)
	name = strings.ReplaceAll(name, "$$", "")

	return &name, nil
}

func GetAdminDivFullNameCode(code string, repo db.ServerRepo) (*string, error) {
	var name string
	count := 0
	if len(code) == 0 {
		return &name, nil
	}
	add, err := repo.SelectAdminDiv(&admindiv.Search{
		DefaultSearchModel: database.DefaultSearchModel{
			Fields: []string{"id"},
		},
		AdminDiv: admindiv.AdminDiv{
			Code: &code,
		},
	})
	if err != nil {
		return nil, err
	}
	tmp := add.ID
	for tmp != uuid.Nil {
		add, err := repo.SelectAdminDiv(&admindiv.Search{
			DefaultSearchModel: database.DefaultSearchModel{
				Fields: []string{"name", "superior_id"},
			},
			AdminDiv: admindiv.AdminDiv{
				BaseModel: database.BaseModel{
					ID: tmp,
				},
			},
		})
		if err != nil {
			return nil, err
		}
		name = fmt.Sprintf("%s$$%s", *add.Name, name)
		tmp = add.SuperiorID
		count++
	}
	name = strings.Replace(name, "$$", ", ", count-1)
	name = strings.ReplaceAll(name, "$$", "")

	return &name, nil
}

func GetRoleName(id uuid.UUID, repo db.ServerRepo) (*string, error) {
	r, err := repo.SelectRole(&role.Search{
		DefaultSearchModel: database.DefaultSearchModel{
			Fields: []string{"name"},
		},
		Role: role.Role{
			BaseModel: database.BaseModel{
				ID: id,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return r.Name, nil
}

func ConvertUsers(d []*user.User, repo db.ServerRepo) []*pb.User {
	usrs := make([]*pb.User, 0)
	for _, s := range d {
		usrs = append(usrs, ConvertOneUser(s, repo))
	}
	return usrs
}

func ConvertPagination(skip, limit int, total int64) *pb.Pagination {
	return &pb.Pagination{
		Total:  int32(total),
		Limit:  int32(limit),
		Offset: int32(skip),
	}
}

func GetAdminDivCode(id uuid.UUID, repo db.ServerRepo) (string, error) {
	if id == uuid.Nil {
		return "", nil
	}
	add, err := repo.SelectAdminDiv(&admindiv.Search{
		AdminDiv: admindiv.AdminDiv{BaseModel: database.BaseModel{
			ID: id,
		}},
	})
	if err != nil {
		return "", err
	}
	return *add.Code, nil

}

func StrToStrArray(a string) ([]string, error) {
	var array []string
	var err error
	if a != "" {
		err = json.Unmarshal([]byte(a), &array)
		if err != nil {
			return nil, xerrors.Errorf("%w", err)
		}
	}
	return array, nil
}
func ConvertOneRecord(s *citizen.Citizen) *pb.Record {

	k := time.Since(time.UnixMilli(utils.Int64Val(s.Birthday))).Hours()
	age := math.Floor(k / 365.25 / 24)
	return &pb.Record{
		AdminDivCode:       utils.StrVal(s.AdminDivCode),
		Gender:             utils.StrVal(s.Gender),
		Age:                int32(age),
		EducationalLevel:   utils.StrVal(s.EducationalLevel),
		CurrentPlaceCode:   utils.StrVal(s.CurrentPlaceCode),
		ResidencePlaceCode: utils.StrVal(s.ResidencePlaceCode),
		HometownCode:       utils.StrVal(s.HometownCode),
		JobName:            utils.StrVal(s.JobName),
		Religion:           utils.StrVal(s.Religion),
	}
}
func ConvertRecords(s []*citizen.Citizen) []*pb.Record {
	ret := make([]*pb.Record, 0)
	for _, k := range s {
		ret = append(ret, ConvertOneRecord(k))
	}
	return ret
}

func ConvertCampaign(s *campaign.Campaign, repo db.ServerRepo) *pb.Campaign {
	if s == nil {
		return nil
	}
	add, err := repo.SelectAdminDiv(&admindiv.Search{
		AdminDiv: admindiv.AdminDiv{
			Code: s.Code,
		},
	})
	if err != nil {
		fmt.Println("err ", err)
	}
	return &pb.Campaign{
		IsDone:       utils.BoolVal(s.IsDone),
		Name:         utils.StrVal(s.Name),
		StartTime:    utils.Int64Val(s.StartTime),
		EndTime:      utils.Int64Val(s.EndTime),
		Id:           s.ID.String(),
		AdminDivCode: utils.StrVal(s.Code),
		AdminDivName: utils.StrVal(add.Name),
		Record:       int32(utils.IntVal(s.RecordNumber)),
	}
}

func ConvertCampigns(camps []*campaign.Campaign, repo db.ServerRepo) []*pb.Campaign {
	if camps == nil {
		return nil
	}
	ret := make([]*pb.Campaign, 0)
	for _, c := range camps {
		ret = append(ret, ConvertCampaign(c, repo))
	}
	return ret
}

func ConvertOneUser(s *user.User, repo db.ServerRepo) *pb.User {
	var adminDivName, roleName *string
	var err error
	if s.AdminDivID != uuid.Nil {
		adminDivName, err = GetAdminDivFullName(s.AdminDivID, repo)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

	}
	if s.RoleID != uuid.Nil {
		roleName, err = GetRoleName(s.RoleID, repo)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}
	adminDivCode, err := GetAdminDivCode(s.AdminDivID, repo)
	if err != nil {
		fmt.Printf("err %w", err)
	}
	return &pb.User{
		Id:                 s.ID.String(),
		Username:           utils.StrVal(s.Username),
		AdminDivId:         s.AdminDivID.String(),
		RoleId:             s.RoleID.String(),
		IsActive:           utils.BoolVal(s.IsActive),
		AdminDivName:       utils.StrVal(adminDivName),
		RoleName:           utils.StrVal(roleName),
		AdminDivCode:       adminDivCode,
		UseDefaultPassword: utils.BoolVal(s.UseDefaultPassword),
	}
}

func ConvertOneCitizen(s *citizen.Citizen) *pb.Citizen {
	return &pb.Citizen{
		Name:               utils.StrVal(s.Name),
		Id:                 s.ID.String(),
		Birthday:           utils.Int64Val(s.Birthday),
		Gender:             utils.StrVal(s.Gender),
		Nationality:        utils.StrVal(s.Nationality),
		FatherName:         utils.StrVal(s.FatherName),
		FatherPid:          utils.StrVal(s.FatherPID),
		MotherName:         utils.StrVal(s.MotherName),
		MotherPid:          utils.StrVal(s.MotherPID),
		CurrentPlace:       utils.StrVal(s.CurrentPlace),
		JobName:            utils.StrVal(s.JobName),
		Pid:                utils.StrVal(s.PID),
		Hometown:           utils.StrVal(s.Hometown),
		Religion:           utils.StrVal(s.Religion),
		EducationalLevel:   utils.StrVal(s.EducationalLevel),
		AdminDivCode:       utils.StrVal(s.AdminDivCode),
		ResidencePlace:     utils.StrVal(s.ResidencePlace),
		AdminDivId:         s.AdminDivID.String(),
		ResidencePlaceCode: utils.StrVal(s.ResidencePlaceCode),
		CurrentPlaceCode:   utils.StrVal(s.CurrentPlaceCode),
		HometownCode:       utils.StrVal(s.HometownCode),
	}
}

func ConvertCitizens(db []*citizen.Citizen) []*pb.Citizen {
	r := make([]*pb.Citizen, 0)
	for _, s := range db {
		r = append(r, ConvertOneCitizen(s))
	}
	return r
}
