package citizen

type CitizenRepo interface {
	SelectCitizen(*Search) (*Citizen, error)
	InsertCitizen(*Citizen) (*Citizen, error)
	// 	ListCitizen(*Search) ([]*Citizen, error)
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
