package handler

import (
	"GinMall/dao"
	"GinMall/service"
	"sync"
	"xorm.io/xorm"
)

var once sync.Once
var userInstance *UserHandlerFunc

func NewUserHandler(db *xorm.Engine) *UserHandlerFunc {
	once.Do(func() {
		userInstance = &UserHandlerFunc{
			Srv: &service.UserService{
				DAO: &dao.UserDAO{
					DB: db,
				},
			},
		}
	})
	return userInstance
}
