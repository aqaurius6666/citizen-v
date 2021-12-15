package campaign

type CampaignRepo interface {
	SelectCampaign(*Search) (*Campaign, error)
	InsertCampaign(*Campaign) (*Campaign, error)
	ListCampaign(*Search) ([]*Campaign, error)
	CountCampaign(*Search) (*int64, error)
	UpdateCampaign(*Search, *Campaign) error
	TotalCampaignRecord(*Search) (*Campaign, error)
}

// func (u *Citizen) AfterFind(db *gorm.DB) error {
// 	r := make([]*Citizen, 0)
// 	db.Model(&Citizen{}).Where(&Citizen{
// 		SuperiorID: u.ID,
// 	}).Find(&r)
// 	u.SubDiv = r
// 	return nil
// }

// func (u *Citizen) AfterCreate(db *gorm.DB) error {
// 	newDb := db.Begin()
// 	if err := newDb.Model(&DivSubDiv{}).Create(&DivSubDiv{
// 		SuperiorID: u.SuperiorID,
// 		AdminDivID: u.ID,
// 	}).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
