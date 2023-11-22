package router

import (
	"GinMall/handler"
	"GinMall/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.Engine, userHandler *handler.UserHandlerFunc) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/login", userHandler.Login)
		userGroup.GET("/info", middleware.JWTAuth(), userHandler.DetailInfo)
		userGroup.POST("/update", middleware.JWTAuth(), userHandler.UpdateInfo)
	}
}
