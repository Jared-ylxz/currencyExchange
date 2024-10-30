package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if inputToken := ctx.GetHeader("Authorization"); inputToken != "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is missing",
			})
			ctx.Abort() // 只执行这一个中间件，不再执行后续的中间件（如有多个中间件）
			return
		}
		// username, err := utils.ParseJWT(inputToken)

		// if err!= nil {}
	}
}
