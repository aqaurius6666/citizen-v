package citizen

import (
	"github.com/aqaurius6666/go-utils/database"
)

type Citizen struct {
	database.BaseModel
	Name         *string `gorm:"column:name;not null" validate:"omitempty,vietnamese"`
	Birthday     *uint64 `gorm:"column:birthday;not null" `
	PID          *string `gorm:"column:pid;index:idx_citizen_pid,unique;not null" validate:"omitempty,pid"`
	Gender       *string `gorm:"column:gender;not null" validate:"eq=male|eq=female"`
	Nationality  *string `gorm:"column:nationality;not null" validate:"omitempty"`
	FatherName   *string `gorm:"column:father_name" validate:"omitempty,vietnamese"`
	FatherPID    *string `gorm:"column:father_pid" validate:"omitempty,pid"`
	MotherName   *string `gorm:"column:mother_name" validate:"omitempty,vietnamese"`
	MotherPID    *string `gorm:"column:mother_pid" validate:"omitempty,pid"`
	CurrentPlace *string `gorm:"column:current_place" validate:"omitempty,vietnamese"`
	JobName      *string `gorm:"column:job_name" validate:"omitempty,vietnamese"`
}

type Search struct {
	database.DefaultSearchModel
	Citizen
}
