package seed

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/campaign"
	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
	"github.com/aqaurius6666/citizen-v/src/internal/db/cockroach"
	"github.com/aqaurius6666/citizen-v/src/internal/db/user"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/aqaurius6666/go-utils/utils"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSeedAdminDiv(t *testing.T) {
	logger := logrus.New()
	repo, err := cockroach.InitServerCDBRepo(context.Background(), logger, cockroach.ServerCDBOptions{
		Cfg: &gorm.Config{},
		Dsn: "postgresql://root:root@cdb:26257/defaultdb?sslmode=disable",
	})
	if err != nil {
		logger.Error(err)
	}
	data := lib.ReadCSV("admindivs.csv")
	for _, e := range data {
		id := e[0]
		name := e[1]
		code := e[2]
		supid := e[3]
		t := e[4]

		add := admindiv.AdminDiv{
			Name: &name,
			Code: &code,
			BaseModel: database.BaseModel{
				ID: uuid.MustParse(id),
			},
			SuperiorID: uuid.MustParse(supid),
			Type:       &t,
		}
		repo.InsertAdminDiv(&add)
	}
}

func TestSeedCampaign(t *testing.T) {
	logger := logrus.New()
	repo, err := cockroach.InitServerCDBRepo(context.Background(), logger, cockroach.ServerCDBOptions{
		Cfg: &gorm.Config{},
		Dsn: "postgresql://root:root@cdb:26257/defaultdb?sslmode=disable",
	})
	if err != nil {
		logger.Error(err)
	}
	data := lib.ReadCSV("campaigns.csv")
	for _, e := range data {
		id := e[0]
		name := e[1]
		code := e[2]
		records, _ := strconv.Atoi(e[3])
		done := e[4] == "true"

		add := campaign.Campaign{
			Name: &name,
			Code: &code,
			BaseModel: database.BaseModel{
				ID: uuid.MustParse(id),
			},
			IsDone:       &done,
			RecordNumber: &records,
		}
		repo.InsertCampaign(&add)
	}
}

func GetRandomPID() string {
	return fmt.Sprintf("001200%06d", rand.Int()%999999)
}

func TestSeedUser(t *testing.T) {
	logger := logrus.New()
	repo, err := cockroach.InitServerCDBRepo(context.Background(), logger, cockroach.ServerCDBOptions{
		Cfg: &gorm.Config{},
		Dsn: "postgresql://root:root@cdb:26257/defaultdb?sslmode=disable",
	})
	if err != nil {
		logger.Error(err)
	}

	data := lib.ReadCSV("admindivs.csv")
	for _, e := range data {
		id := e[0]
		code := e[2]
		username := fmt.Sprintf("citizen%s", code)
		repo.InsertUser(&user.User{
			Username:           &username,
			UseDefaultPassword: utils.BoolPtr(false),
			RoleID:             uuid.Nil,
			AdminDivID:         uuid.MustParse(id),
			IsActive:           utils.BoolPtr(true),
			AdminDivCode:       &code,
			HashPassword:       utils.StrPtr("Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU="),
		})
	}
}

func TestSeedCitizens(t *testing.T) {
	logger := logrus.New()
	repo, err := cockroach.InitServerCDBRepo(context.Background(), logger, cockroach.ServerCDBOptions{
		Cfg: &gorm.Config{},
		Dsn: "postgresql://root:root@cdb:26257/defaultdb?sslmode=disable",
	})
	if err != nil {
		logger.Error(err)
	}

	NAME := []string{
		"Nguy???n V??n A",
		"Nguy???n Th??? B",
		"Nguy???n Tr???n C",
		"Nguy???n Thanh D",
		"T??? Minh E",
		"Tr???n ????nh F",
		"Nguy???n Quang H",
		"Nguy???n Th??nh G",
		"Nguy???n Minh K",
		"Nguy???n Nh?? J",
	}

	JOBS := []string{
		"T??i x???",
		"Gi??o vi??n",
		"Sinh vi??n",
		"H???c sinh",
		"C??ng nh??n",
		"Kinh doanh",
		"Lao c??ng",
	}
	EDUCATIONAL_LEVELS := []string{
		"12/12",
		"9/12",
		"?????i h???c",
		"Cao ?????ng",
	}

	RELIGION := []string{
		"Kh??ng",
		"Ph???t Gi??o",
		"C??ng Gi??o",
	}
	GENDER := []string{
		"male",
		"female",
	}

	data := lib.ReadCSV("campaigns.csv")
	for _, e := range data {
		code := e[2]
		records, _ := strconv.Atoi(e[3])

		for i := 0; i < records; i++ {
			name := NAME[rand.Int()%len(NAME)]
			fname := NAME[rand.Int()%len(NAME)]
			mname := NAME[rand.Int()%len(NAME)]

			religion := RELIGION[rand.Int()%len(RELIGION)]

			job := JOBS[rand.Int()%len(JOBS)]
			edu := EDUCATIONAL_LEVELS[rand.Int()%len(EDUCATIONAL_LEVELS)]

			age := rand.Int() % 60
			tp := time.Date(time.Now().Year()-age, time.Month(rand.Int()%12), rand.Int()%25, 0, 0, 0, 0, time.Local)
			tmp := tp.UnixMilli()

			pid := GetRandomPID()
			fid := GetRandomPID()
			mid := GetRandomPID()

			add, err := repo.SelectAdminDiv(&admindiv.Search{
				AdminDiv: admindiv.AdminDiv{
					Code: &code,
				},
			})
			if err != nil {
				continue
			}

			gender := GENDER[rand.Int()%len(GENDER)]
			place, err := lib.GetAdminDivFullNameCode(code, repo)
			assert.Nil(t, err)
			repo.InsertCitizen(&citizen.Citizen{
				Name:               &name,
				Birthday:           &tmp,
				PID:                &pid,
				Gender:             &gender,
				Nationality:        utils.StrPtr("Vi???t Nam"),
				CurrentPlace:       place,
				ResidencePlace:     place,
				Hometown:           place,
				JobName:            &job,
				EducationalLevel:   &edu,
				FatherName:         &fname,
				FatherPID:          &fid,
				MotherName:         &mname,
				MotherPID:          &mid,
				AdminDivCode:       &code,
				AdminDivID:         add.ID,
				Religion:           &religion,
				CurrentPlaceCode:   &code,
				ResidencePlaceCode: &code,
				HometownCode:       &code,
			})
		}

	}
}

func TestGenPID(t *testing.T) {

	a := GetRandomPID()
	if len(a) != 12 {
		t.Error(a)
	}
}
