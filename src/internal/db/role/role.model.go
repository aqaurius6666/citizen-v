package role

import "github.com/aquarius6666/go-utils/database"

type Role struct {
	database.BaseModel
	Name *string `gorm:"column:name;index:idx_name,unique;not null"`
}

type Search struct {
	database.DefaultSearchModel
	Role
}
