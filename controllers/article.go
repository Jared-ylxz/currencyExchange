package controllers

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	any_username, exists := c.Get("username")
	fmt.Println(111, any_username, any_username.(string))
	if exists {
		username := any_username.(string)
		fmt.Println(222, username)
		var user models.User
		result := global.Db.First(&user, "username = ?", username)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		article.AuthorID = user.ID
	}

	err := global.Db.Create(&article).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, article)
}

func GetArticles(c *gin.Context) {
	var articles []models.Article
	result := global.Db.Find(&articles, "deleted_at IS NULL")
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Articles not found"})
		return
	}
	c.JSON(http.StatusOK, articles)
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
