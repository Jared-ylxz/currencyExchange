package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // create a gin router instance
	r.Use(favicon.New("./favicon.ico"))

	public := r.Group("/api/public")
	{
		public.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	private := r.Group("/api/auth")
	{
		private.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(
				http.StatusOK,
				gin.H{
					"message": "login success",
				},
			)
		})
		private.POST("/register", func(ctx *gin.Context) {
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
