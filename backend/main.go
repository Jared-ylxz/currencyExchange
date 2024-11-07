package main

import (
	"exchangeapp/config"
	"exchangeapp/router"
	"fmt"
)

func main() {
	config.InitConfig()
	r := router.SetupRouter()
	port := config.AppConfig.App.Port
	host := config.AppConfig.App.Host
	r.Run(fmt.Sprintf("%s:%s", host, port))
	// TODO 优雅地关闭服务器
}
