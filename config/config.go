package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name   string
		Port   string
		Host   string
		Secret string
	}
	Database struct {
		Host            string
		Port            string
		User            string
		Name            string
		Password        string
		MaxOpenConns    int
		MaxIdleConns    int
		ConnMaxLifetime int
	}
	Redis struct {
		Host     string
		Port     string
		Password string
		DB       int
		PoolSize int
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")   // name of config file (without extension)不需要加.yaml后缀
	viper.AddConfigPath("./config") // set the path of the config file to current directory
	viper.SetConfigType("yaml")     // set the config file type to yaml

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
		// panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	AppConfig = &Config{} // 实例化Config结构体
	err = viper.Unmarshal(AppConfig)
	if err != nil {
		log.Fatalf("Error unmarshaling config file, %s", err)
	}

	InitDB()
	InitRedis()
}
