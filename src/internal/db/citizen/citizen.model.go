package citizen

import (
	"github.com/aqaurius6666/go-utils/database"
	"github.com/google/uuid"
)

type Citizen struct {
	database.BaseModel
	Name               *string   `gorm:"column:name;not null" validate:"omitempty,vietnamese"`
	Birthday           *int64    `gorm:"column:birthday;not null" `
	PID                *string   `gorm:"column:pid;index:idx_citizen_pid,unique;not null" validate:"omitempty,pid"`
	Gender             *string   `gorm:"column:gender;not null" validate:"eq=Male|eq=Female"`
	Nationality        *string   `gorm:"column:nationality;not null" validate:"omitempty"`
	CurrentPlace       *string   `gorm:"column:current_place" validate:"omitempty,vietnamese"`
	ResidencePlace     *string   `gorm:"column:residence_place" validate:"omitempty,vietnamese"`
	Hometown           *string   `gorm:"column:hometown" validate:"omitempty,vietnamese"`
	JobName            *string   `gorm:"column:job_name" validate:"omitempty,vietnamese"`
	Religion           *string   `gorm:"column:religion" validate:"omitempty,vietnamese"`
	EducationalLevel   *string   `gorm:"column:educational_level"`
	FatherName         *string   `gorm:"column:father_name" validate:"omitempty,vietnamese"`
	FatherPID          *string   `gorm:"column:father_pid" validate:"omitempty,pid"`
	MotherName         *string   `gorm:"column:mother_name" validate:"omitempty,vietnamese"`
	MotherPID          *string   `gorm:"column:mother_pid" validate:"omitempty,pid"`
	AdminDivCode       *string   `gorm:"column:admin_div_code"`
	AdminDivID         uuid.UUID `gorm:"column:admin_div_id;type:uuid"`
	CurrentPlaceCode   *string   `gorm:"column:current_place_code" validate:"omitempty,code"`
	ResidencePlaceCode *string   `gorm:"column:residence_place_code" validate:"omitempty,code"`
	HometownCode       *string   `gorm:"column:hometown_code" validate:"omitempty,code"`
}

type Search struct {
	database.DefaultSearchModel
	Citizen
	ArrayCode []string
}
