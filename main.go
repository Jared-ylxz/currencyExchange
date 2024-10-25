package main

import (
	"exchangeapp/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	fmt.Println("Hello, World!")

	r := gin.Default() // create a gin router instance
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf("%s:%s", config.AppConfig.App.Host, config.AppConfig.App.Port))
}
