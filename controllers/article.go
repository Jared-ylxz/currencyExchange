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

	c.JSON(http.StatusCreated, gin.H{"data": article})
}
