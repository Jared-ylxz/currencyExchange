package middlewares

import (
	"exchangeapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		inputToken := ctx.GetHeader("Authorization")
		if inputToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is missing",
			})
			ctx.Abort() // 只执行这一个中间件，不再执行后续的中间件（如有多个中间件）
			return
		}

		username, err := utils.ParseJWT(inputToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Set("username", username) // 将用户名存入上下文
		ctx.Next()                    // 执行后续的中间件
	}
}
