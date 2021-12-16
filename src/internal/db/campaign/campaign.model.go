package campaign

import (
	"github.com/aqaurius6666/go-utils/database"
)

type Campaign struct {
	database.BaseModel
	Name         *string  `gorm:"column:name;not null" validate:"omitempty,vietnamese"`
	IsDone       *bool    `gorm:"column:is_done;default:false;not null"`
	Code         *string  `gorm:"column:code"`
	RecordNumber *int     `gorm:"record_number;default:0"`
	EndTime      *int64   `gorm:"end_time"`
	Percent      *float32 `gorm:"column:percent"`
}

type Search struct {
	database.DefaultSearchModel
	Campaign
}
