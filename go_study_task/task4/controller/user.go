package controller

import (
  "github.com/gin-gonic/gin"
  "github.com/i159800/web3_homework/go_study_task/task4/config"
  "github.com/i159800/web3_homework/go_study_task/task4/model"
  "github.com/i159800/web3_homework/go_study_task/task4/utils"
  "log"
  "net/http"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or Password required"})
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Register success"})
}

func Login(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user model.User
	if err := config.DB.Where("username = ? and password = ?", req.Username, req.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	token, _ := utils.GenerateToken(uint(user.ID), []string{"admin"})
	log.Println(user.ID, token)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
