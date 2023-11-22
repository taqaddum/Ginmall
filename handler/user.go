package handler

import (
	"GinMall/serializer/request"
	"GinMall/serializer/response"
	"GinMall/serializer/status"
	"GinMall/service"
	"github.com/gin-gonic/gin"
)

type UserHandlerFunc struct {
	Srv service.UserSrvApi
}

func (handle UserHandlerFunc) Login(ctx *gin.Context) {
	var req = request.LoginRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, response.FailedCode(status.InvalidParams))
		return
	}
	ctx.JSON(200, handle.Srv.LoginUser(&req))
}

func (handle UserHandlerFunc) Register(ctx *gin.Context) {
	var req = request.RegisterRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, response.FailedCode(status.InvalidParams))
		return
	}
	ctx.JSON(200, handle.Srv.RegisterUser(&req))
}

func (handle UserHandlerFunc) DetailInfo(ctx *gin.Context) {
	var uname = ctx.GetString("username")
	ctx.JSON(200, handle.Srv.GetUserDetail(uname))
}

func (handle UserHandlerFunc) UpdateInfo(ctx *gin.Context) {
	var req = request.UpdateRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, response.FailedCode(status.InvalidParams))
		return
	}
	ctx.JSON(200, handle.Srv.UpdateUserInfo(&req))
}
