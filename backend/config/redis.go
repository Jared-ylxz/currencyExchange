package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"

	"exchangeapp/global"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Host + ":" + AppConfig.Redis.Port,
		Password: AppConfig.Redis.Password,
		DB:       AppConfig.Redis.DB,
		PoolSize: AppConfig.Redis.PoolSize,
	})
	if RedisClient.Ping(context.Background()).Err() != nil {
		log.Fatalln("Redis connection failed")
	}
	global.RedisClient = RedisClient
}
