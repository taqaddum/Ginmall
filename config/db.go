package config

import (
	"fmt"
	"net/url"
)

type dbConf struct {
	Host   string `validate:"required,ipv4"`
	Port   uint16 `validate:"required_with=Host,gte=0"`
	User   string `validate:"required"`
	Passwd string `validate:"required"`
	Name   string `validate:"required"`
	Params map[string]string
}

func (db dbConf) DSN() string {
	var params = make(url.Values)
	for k, v := range db.Params {
		params.Add(k, v)
	}

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(db.User, db.Passwd),
		Host:     fmt.Sprintf("%s:%d", db.Host, db.Port),
		Path:     db.Name,
		RawQuery: params.Encode(),
	}
	return u.String()
}
