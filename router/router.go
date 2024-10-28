package router

import (
	"exchangeapp/controllers"
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
		private.POST("/login", controllers.Login)
		private.POST("/register", controllers.Register)
	}

	return r
}
