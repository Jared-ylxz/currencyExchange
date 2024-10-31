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

	// var user_id int
	// user_id, exists := c.Get("user_id")
	// if user_id.(string) {
	// 	if exists {
	// 		nuser_id, err := strconv.ParseUint(user_id, 10, 64)
	// 		if err == nil {
	// 			article.AuthorID = nuser_id
	// 		}
	// 	}
	username, _ := c.Get("username")
	fmt.Println(11111111111, username)

	if err := global.Db.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": article})
}
