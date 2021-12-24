package admindiv

import (
	"github.com/aqaurius6666/go-utils/database"
	"github.com/google/uuid"
)

type AdminDiv struct {
	database.BaseModel
	Name             *string     `gorm:"column:name;index:idx_admin_div_name,unique;not null" validate:"omitempty,vietnamese"`
	Code             *string     `gorm:"column:code;index:idx_admin_div_code,unique;not null" validate:"omitempty,regexp=^[0-9]+$"`
	Type             *string     `gorm:"column:type;not null" validate:"eq=CITY|eq=DISTRICT|eq=TOWN|eq=BLOCK"`
	SuperiorID       uuid.UUID   `gorm:"column:superior_id;type:uuid;not null"`
	SuperiorAdminDiv *AdminDiv   `gorm:"foreignKey:SuperiorID"`
	SubDiv           []*AdminDiv `gorm:"many2many:div_sub_divs;joinForeignKey:SuperiorID"`
}

type Search struct {
	database.DefaultSearchModel
	AdminDiv
	SuperiorCode *string
}
