package middlewares

import (
	"exchangeapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleWare 验证token中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		//如果不存在相应token，则拦截请求
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization."})
			ctx.Abort()
			return
		}
		username, err := utils.ParseJWT(token)
		//如果不存在相应用户，则拦截请求
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token."})
			ctx.Abort()
			return
		}
		//将用户名存入gin的上下文，方便其他函数Get使用
		ctx.Set("username", username)
		//继续处理请求
		ctx.Next()

	}
}
