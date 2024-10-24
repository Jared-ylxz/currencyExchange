package main

import (
	"fmt"
	"exchangeapp/config"
)

func main() {
	config.InitConfig()
	fmt.Println("Hello, World!")
	fmt.Println(config.AppConfig.App.Port)
}