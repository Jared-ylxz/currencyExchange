package controllers

import (
	"exchangeapp/global"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func LikeArticle(ctx *gin.Context) {
	articleId := ctx.Param("articleId")
	redisKey := fmt.Sprintf("article:%s:likes", articleId)

	if err := global.RedisClient.Incr(ctx, redisKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to like article",
		})
		return
	}
	likes := global.RedisClient.Get(ctx, redisKey).Val()

	ctx.JSON(http.StatusOK, gin.H{
		"likes": likes,
	})
}

func GetLikes(ctx *gin.Context) {
	articleId := ctx.Param("articleId")
	redisKey := fmt.Sprintf("article:%s:likes", articleId)

	likes, err := global.RedisClient.Get(ctx, redisKey).Result()
	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get likes",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"likes": likes,
	})
}
