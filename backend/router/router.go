package router

import (
	"exchangeapp/controllers"
	"exchangeapp/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func SetupRouter() *gin.Engine {
	router := gin.Default() // create a gin router instance
	router.Use(favicon.New("./favicon.ico"))

	// middleware to handle CORS
	router.Use(middlewares.CORSMiddleware())

	public := router.Group("/api/v1/public")
	{
		public.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	user := router.Group("/api/v1/users")
	{
		user.POST("/login", controllers.Login)
		user.POST("/register", controllers.Register)
	}

	private := router.Group("/api/v1")
	private.GET("/exchange-rates", controllers.GetExchangeRate)
	private.GET("/articles", controllers.GetArticles)
	private.GET("/articles/:id", controllers.GetArticle)
	private.GET("/article-likes/:articleId", controllers.GetLikes)
	private.Use(middlewares.AuthMiddleware())
	{
		private.POST("/exchange-rates", controllers.CreateExchangeRate)
		private.POST("/articles", controllers.CreateArticle)
		private.DELETE("/articles/:id", controllers.DeleteArticle)
		private.POST("/article-likes/:articleId", controllers.LikeArticle)
	}

	return router
}
