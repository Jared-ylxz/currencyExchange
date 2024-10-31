package router

import (
	"exchangeapp/controllers"
	"exchangeapp/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // create a gin router instance
	r.Use(favicon.New("./favicon.ico"))

	public := r.Group("/api/v1/public")
	{
		public.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	user := r.Group("/api/v1/users")
	{
		user.POST("/login", controllers.Login)
		user.POST("/register", controllers.Register)
	}

	private := r.Group("/api/v1")
	private.GET("/exchange-rates", controllers.GetExchangeRate)
	private.Use(middlewares.AuthMiddleware())
	{
		private.POST("/exchange-rates", controllers.CreateExchangeRate)
	}

	return r
}
