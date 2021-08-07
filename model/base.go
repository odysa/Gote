package model

import (
	"time"
)

type BaseModel struct {
	ID       int       `json:"id" gorm:"AUTO_INCREMENT"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	IsDelete int       `json:"is_delete"`
}
