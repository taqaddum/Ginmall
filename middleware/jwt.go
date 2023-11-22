package middleware

import (
	"GinMall/serializer/status"
	"GinMall/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		claims, err := utils.ParseToken(token)
		if err != nil {
			ctx.JSON(401, gin.H{
				"code": status.ErrorUnLogin,
				"msg":  status.Msg[status.ErrorUnLogin],
			})
			ctx.Abort()
			return
		}

		if time.Now().After(claims.ExpiresAt.Time) {
			ctx.JSON(401, gin.H{
				"code": status.ErrorTokenExpired,
				"msg":  status.Msg[status.ErrorTokenExpired],
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
