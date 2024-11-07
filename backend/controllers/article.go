package controllers

import (
	"encoding/json"
	"exchangeapp/global"
	"exchangeapp/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var allCacheKey = "articles"
var oneCacheKey = "articles:%d"

func CreateArticle(ctx *gin.Context) {
	var article models.Article
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	any_username, exists := ctx.Get("username")
	if exists {
		username := any_username.(string)
		var user models.User
		result := global.Db.First(&user, "username = ?", username)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		article.AuthorID = user.ID
	}

	err := global.Db.Create(&article).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, article)

	if err := global.RedisClient.Del(ctx, allCacheKey).Err(); err != nil {
		fmt.Println("Redis delete error:", err)
	}
}

func GetArticles(ctx *gin.Context) {
	redisData, err := global.RedisClient.Get(ctx, allCacheKey).Result()
	if err == nil && redisData != "" {
		// redis 缓存命中
		fmt.Println("Redis get data!")
		var articles []models.Article
		// 将 JSON 字符串反序列化为文章列表
		err := json.Unmarshal([]byte(redisData), &articles)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal articles from cache"})
			return
		}
		ctx.JSON(http.StatusOK, articles)
		return
	} else if err != nil {
		// redis 缓存未命中, 从数据库获取数据并缓存
		fmt.Println("Redis not found:", err)
		var articles []models.Article
		result := global.Db.Find(&articles, "deleted_at IS NULL")
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Articles not found"})
			return
		}
		// 将文章列表序列化为 JSON 字符串
		jsonData, err := json.Marshal(articles)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal articles"})
			return
		}
		// 将 JSON 字符串存储到 Redis 中
		statusCmd := global.RedisClient.Set(ctx, allCacheKey, jsonData, time.Hour*24)
		if statusCmd.Err() != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache in Redis"})
			return
		}
		ctx.JSON(http.StatusOK, articles)
		return
	} else {
		// redis 报错
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
}

func GetArticle(ctx *gin.Context) {
	var article models.Article
	id := ctx.Param("id")
	result := global.Db.First(&article, "id = ?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	ctx.JSON(http.StatusOK, article)
}

func DeleteArticle(ctx *gin.Context) {
	var article models.Article
	id := ctx.Param("id")
	result := global.Db.First(&article, "id = ?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	global.Db.Delete(&article) // 如果一个 model 有 DeletedAt 字段，则软删除。硬删除需要 db.Unscoped().Delete(&article)
	ctx.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})

	if err := global.RedisClient.Del(ctx, allCacheKey).Err(); err != nil {
		fmt.Println("Redis delete error:", err)
	}
}
