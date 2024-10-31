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
	if exists {
		username := any_username.(string)
		user := models.User{Username: username}
		fmt.Println(2222222, user)
		a := global.Db.First(&user)
		fmt.Println(111111111111111111, &a, a, a.Error, a.RowsAffected)
		// var author models.Article = global.Db.First(&user)
		// fmt.Println(11111111111, author)
		// article.AuthorID = author.ID
	}

	err := global.Db.Create(&article).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": article})
}
