package seed

import (
	"context"
	"strconv"
	"testing"

	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/campaign"
	"github.com/aqaurius6666/citizen-v/src/internal/db/cockroach"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
