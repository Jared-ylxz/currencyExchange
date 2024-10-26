package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // create a gin router instance

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(
				http.StatusOK,
				gin.H{
					"message": "login success",
				},
			)
		})
		auth.POST("/register", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(
				http.StatusOK,
				gin.H{
					"message": "register success",
				},
			)
		})
	}
	return r
}
