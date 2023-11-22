package model

import (
	"GinMall/model/enum"
)

type User struct {
	Preset
	Authority enum.Authority `xorm:"smallint(2) default(2)" json:"role,omitempty"`
	Username  string         `xorm:"varchar(64) unique" json:"username,omitempty"`
	Password  string         `xorm:"varchar(128)" json:"-"`
	Gender    enum.Gender    `xorm:"char(1) default('M')" json:"gender,omitempty"`
	Phone     string         `xorm:"varchar(11)" json:"phone,omitempty"`
	Email     string         `xorm:"varchar(128)" json:"email,omitempty"`
	Avatar    string         `xorm:"varchar(255)" json:"avatar,omitempty"`
}
