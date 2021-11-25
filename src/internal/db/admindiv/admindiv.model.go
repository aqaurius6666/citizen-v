package admindiv

import (
	"github.com/aquarius6666/go-utils/database"
	"github.com/google/uuid"
)

type AdminDiv struct {
	database.BaseModel
	Name             *string   `gorm:"column:name;index:idx_name,unique;not null"`
	Code             *string   `gorm:"column:code;index:idx_code,unique;not null"`
	Type             *string   `gorm:"column:type;not null"`
	SuperiorID       uuid.UUID `gorm:"column:superior_id;type:uuid;not null"`
	SuperiorAdminDiv *AdminDiv `gorm:"foreignKey:SuperiorID;reference:code"`
}

type Search struct {
	database.DefaultSearchModel
	AdminDiv
}
