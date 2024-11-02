package controllers

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"fmt"
	"net/http"

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
		fmt.Println("Redis get data:", redisData)
		ctx.JSON(http.StatusOK, redisData)
		return
	} else if err != nil {
		fmt.Println("Redis get error:", err)
		var articles []models.Article
		result := global.Db.Find(&articles, "deleted_at IS NULL")
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Articles not found"})
			return
		}
		ctx.JSON(http.StatusOK, articles)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// var articles []models.Article
	// // var ArticleWithLikes struct {
	// // 	models.Article
	// // 	Likes int `json:"likes"`
	// // }

	// result := global.Db.Find(&articles, "deleted_at IS NULL")
	// if result.Error != nil {
	// 	ctx.JSON(http.StatusNotFound, gin.H{"error": "Articles not found"})
	// 	return
	// }

	// // for a, _ := range articles {
	// // 	var a.likes int
	// // 	a.likes := global.RedisClient.Get(fmt.Sprintf("article:%d:likes", a.ID)).Int()
	// // }
	// ctx.JSON(http.StatusOK, articles)
}

func GetArticle(c *gin.Context) {
	var article models.Article
	id := c.Param("id")
	result := global.Db.First(&article, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, article)
}

func DeleteArticle(c *gin.Context) {
	var article models.Article
	id := c.Param("id")
	result := global.Db.First(&article, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	global.Db.Delete(&article) // 如果一个 model 有 DeletedAt 字段，则软删除。硬删除需要 db.Unscoped().Delete(&article)
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
