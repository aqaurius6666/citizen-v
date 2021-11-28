package admindiv

import (
	"gorm.io/gorm"
)

type AdminDivRepo interface {
	SelectAdminDiv(*Search) (*AdminDiv, error)
	InsertAdminDiv(*AdminDiv) (*AdminDiv, error)
	ListAdminDiv(*Search) ([]*AdminDiv, error)
}

func (u *AdminDiv) AfterFind(db *gorm.DB) error {
	r := make([]*AdminDiv, 0)
	db.Model(&AdminDiv{}).Where(&AdminDiv{
		SuperiorID: u.ID,
	}).Find(&r)
	u.SubDiv = r
	return nil
}
