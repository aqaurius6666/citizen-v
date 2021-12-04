package user

import (
	"github.com/aqaurius6666/citizen-v/src/internal/db/role"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/google/uuid"
)

type User struct {
	database.BaseModel
	Username     *string    `gorm:"column:username;index:idx_user_username;unique;not null"`
	HashPassword *string    `gorm:"column:hash_password;not null"`
	Role         *role.Role `gorm:"foreignKey:RoleID"`
	RoleID       uuid.UUID  `gorm:"column:role_id;type:uuid;not null"`
}

type Search struct {
	database.DefaultSearchModel
	User
}
