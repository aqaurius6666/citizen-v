package role

import "github.com/aqaurius6666/go-utils/database"

type Role struct {
	database.BaseModel
	Name *string `gorm:"column:name;index:idx_role_name,unique;not null"`
}

type Search struct {
	database.DefaultSearchModel
	Role
}
