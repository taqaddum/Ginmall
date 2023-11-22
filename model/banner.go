package model

type Banner struct {
	Preset
	Url      string `json:"url"`
	Hash     string `json:"hash" xorm:"varchar(255)"`
	Username string `json:"username" xorm:"varchar(64)"`
}
