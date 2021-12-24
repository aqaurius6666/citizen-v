package api

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/lib/validate"
	"github.com/aqaurius6666/citizen-v/src/internal/model"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/aqaurius6666/go-utils/utils"
	"github.com/gin-gonic/gin"
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
		Results:   lib.ConvertRecords(ctz),
		ScopeCode: utils.StrVal(search.AdminDivCode),
	}, nil
}

func (s *StatisticService) HandleData(res *pb.GetStatisticsCitizensResponse_Data) interface{} {
	genderMap := make(map[string]int)
	ageMap := make(map[string]int)
	educationalLevelMap := make(map[string]int)
	jobNameMap := make(map[string]int)
	currentPlaceMap := make(map[string]int)
	residencePlaceMap := make(map[string]int)
	hometownMap := make(map[string]int)
	religionMap := make(map[string]int)
	for _, r := range res.Results {
		handleMap(genderMap, r.Gender)
		handleAge(ageMap, int(r.Age))
		handleMap(educationalLevelMap, r.EducationalLevel)
		handleMap(jobNameMap, r.JobName)
		handleMap(religionMap, r.Religion)
		handlePlace(currentPlaceMap, r.CurrentPlaceCode, res.ScopeCode)
		handlePlace(hometownMap, r.HometownCode, res.ScopeCode)
		handlePlace(residencePlaceMap, r.ResidencePlaceCode, res.ScopeCode)
	}
	return gin.H{
		"gender":           genderMap,
		"age":              ageMap,
		"educationalLevel": educationalLevelMap,
		"jobName":          jobNameMap,
		"currentPlace":     currentPlaceMap,
		"residencePlace":   residencePlaceMap,
		"hometown":         hometownMap,
		"religion":         religionMap,
	}
}

func handleMap(mp map[string]int, value string) {
	mp["total"]++
	mp[value]++
}

func handlePlace(mp map[string]int, value string, scope string) {
	code := value[0 : len(scope)+2]
	mp[code]++
	mp["total"]++
}

func handleAge(mp map[string]int, value int) {
	mp["total"]++
	if value <= 6 {
		mp["0-6"]++
		return
	}
	if value <= 17 {
		mp["6-17"]++
		return
	}
	if value <= 40 {
		mp["18-40"]++
		return
	}
	if value <= 65 {
		mp["41-65"]++
		return
	}
	mp["65+"]++

}
