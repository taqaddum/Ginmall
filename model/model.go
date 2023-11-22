package model

import "time"

type Preset struct {
	ID        int        `json:"id"`
	CreatedAt time.Time  `xorm:"created" json:"created_at"`
	UpdatedAt time.Time  `xorm:"updated" json:"updated_at"`
	DeletedAt *time.Time `xorm:"deleted" json:"-"`
}
