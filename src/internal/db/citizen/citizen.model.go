package citizen

import (
	"github.com/aquarius6666/go-utils/database"
)

type Citizen struct {
	database.BaseModel
	Name     *string `gorm:"column:name;index:idx_name,unique;not null"`
	Birthday uint64
}

type Search struct {
	database.DefaultSearchModel
	Citizen
}
