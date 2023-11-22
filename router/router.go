package router

import (
	"GinMall/handler"
	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func InitRouter(db *xorm.Engine) *gin.Engine {
	router := gin.Default()
	InitUserRouter(router, handler.NewUserHandler(db))
	return router
}
