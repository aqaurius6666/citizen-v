package campaign

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db/campaign"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/aqaurius6666/go-utils/database/cockroach"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	_ campaign.CampaignRepo = (*CampaignCDBRepo)(nil)
)

func applySearch(db *gorm.DB, search *campaign.Search) *gorm.DB {
	if search.ID != uuid.Nil {
		db = db.Where(&campaign.Campaign{
			BaseModel: database.BaseModel{
				ID: search.ID,
			},
		})
	}
	if search.Name != nil {
		db = db.Where(&campaign.Campaign{
			Name: search.Name,
		})
	}
	// if search.Code != nil {
	// 	db = db.Where(clause.Like{
	// 		Column: "code",
	// 		Value:  *search.Code + "%",
	// 	})
	// }
	orderBy := "name"
	isDesc := true
	if a := search.OrderBy; a != "" {
		orderBy = a

	}
	if orderType := search.OrderType; orderType != "DESC" {
		isDesc = false
	}
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: orderBy}, Desc: isDesc})
	return db
}

type CampaignCDBRepo struct {
	cockroach.CDBRepository
}

func (u *CampaignCDBRepo) SelectCampaign(search *campaign.Search) (*campaign.Campaign, error) {
	r := campaign.Campaign{}
	if err := applySearch(u.Db, search).First(&r).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, campaign.ErrNotFound
		}
		return nil, err
	}
	return &r, nil
}

func (u *CampaignCDBRepo) InsertCampaign(value *campaign.Campaign) (*campaign.Campaign, error) {
	if err := u.Db.Create(value).Error; err != nil {
		return nil, campaign.ErrInsertFail
	}
	return value, nil
}

func (u *CampaignCDBRepo) ListCampaign(search *campaign.Search) ([]*campaign.Campaign, error) {
	r := make([]*campaign.Campaign, 0)
	if err := applySearch(u.Db, search).Offset(search.Skip).Limit(search.Limit).Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (u *CampaignCDBRepo) CountCampaign(search *campaign.Search) (*int64, error) {
	var r int64
	if err := applySearch(u.Db, search).Model(&campaign.Campaign{}).Count(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

func (u *CampaignCDBRepo) UpdateCampaign(search *campaign.Search, value *campaign.Campaign) error {
	if err := applySearch(u.Db, search).Model(&campaign.Campaign{}).Updates(&value).Error; err != nil {
		return campaign.ErrUpdateFail
	}
	return nil
}

func (u *CampaignCDBRepo) TotalCampaignRecord(search *campaign.Search) (*campaign.Campaign, error) {
	var v campaign.Campaign
	if err := u.Db.Model(&campaign.Campaign{}).Select(`
		sum(record_number) as record_number,
		bool_and(is_done) as is_done,
		sum(if(is_done, 1, 0)) / count(*) as percent
		`).
		Where("code like ?", *search.Code+"%").
		Scan(&v).Error; err != nil {
		return &v, campaign.ErrUpdateFail
	}
	return &v, nil
}
