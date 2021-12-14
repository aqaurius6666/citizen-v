package user

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/role"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/google/uuid"
)

type User struct {
	database.BaseModel
	Username           *string            `gorm:"column:username;index:idx_user_username;unique;not null"`
	HashPassword       *string            `gorm:"column:hash_password;not null"`
	UseDefaultPassword *bool              `gorm:"column:use_default_password;default:true;not null"`
	Role               *role.Role         `gorm:"foreignKey:RoleID"`
	RoleID             uuid.UUID          `gorm:"column:role_id;type:uuid;not null" `
	PermissionZoneID   uuid.UUID          `gorm:"column:permission_zone_id;uuid" `
	PermissionZone     *admindiv.AdminDiv `gorm:"foreignKey:PermissionZoneID"`
}

type Search struct {
	database.DefaultSearchModel
	User
}
