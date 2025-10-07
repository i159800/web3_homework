package controller

import (
  "github.com/gin-gonic/gin"
  "github.com/i159800/web3_homework/go_study_task/task4/config"
  "github.com/i159800/web3_homework/go_study_task/task4/model"
  "net/http"
)

func CreatePost(c *gin.Context) {
	var post model.Post
	userIDVal, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := uint(0)
	if f, ok := userIDVal.(float64); ok {
		userID = uint(f)
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.UserId = userID
	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func GetPosts(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Preload("Posts").Where("username = ?", user.Username).Find(&user)
	c.JSON(http.StatusOK, gin.H{"userId":user.ID,"userName":user.Username,"userPosts": user.Posts})
}
