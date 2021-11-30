package citizen

import (
	"github.com/aquarius6666/go-utils/database"
)

type Citizen struct {
	database.BaseModel
	Name         *string `gorm:"column:name;not null"`
	Birthday     *uint64 `gorm:"column:birthday;not null"`
	PID          *string `gorm:"column:pid;index:idx_citizen_pid,unique;not null"`
	Gender       *string `gorm:"column:gender;not null"`
	Nationality  *string `gorm:"column:nationality;not null"`
	FatherName   *string `gorm:"column:father_name"`
	FatherPID    *string `gorm:"column:father_pid"`
	MotherName   *string `gorm:"column:mother_name"`
	MotherPID    *string `gorm:"column:mother_pid"`
	CurrentPlace *string `gorm:"column:current_place"`
	JobName      *string `gorm:"column:job_name"`
}

type Search struct {
	database.DefaultSearchModel
	Citizen
}
