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
	return fmt.Sprintf("001200%06d", rand.Int())
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
		"Nguyễn Văn A",
		"Nguyễn Thị B",
		"Nguyễn Trần C",
		"Nguyễn Thanh D",
		"Tạ Minh E",
		"Trần Đình F",
		"Nguyễn Quang H",
		"Nguyễn Thành G",
		"Nguyễn Minh K",
		"Nguyễn Như J",
	}

	JOBS := []string{
		"Tài xế",
		"Giáo viên",
		"Sinh viên",
		"Học sinh",
		"Công nhân",
		"Kinh doanh",
		"Lao công",
	}
	EDUCATIONAL_LEVELS := []string{
		"12/12",
		"9/12",
		"Đại học",
		"Cao đẳng",
	}

	RELIGION := []string{
		"Không",
		"Phật Giáo",
		"Công Giáo",
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
			tmp := uint64(tp.UnixMilli())

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
				Name:             &name,
				Birthday:         &tmp,
				PID:              &pid,
				Gender:           &gender,
				Nationality:      utils.StrPtr("Việt Nam"),
				CurrentPlace:     place,
				ResidencePlace:   place,
				Hometown:         place,
				JobName:          &job,
				EducationalLevel: &edu,
				FatherName:       &fname,
				FatherPID:        &fid,
				MotherName:       &mname,
				MotherPID:        &mid,
				AdminDivCode:     &code,
				AdminDivID:       add.ID,
				Religion:         &religion,
			})
		}

	}
}
