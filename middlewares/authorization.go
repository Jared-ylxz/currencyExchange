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
			ctx.Abort() // stop the chain of handlers
			return
		}
		// username, err := utils.ParseJWT(inputToken)
	}
}
