package router

import (
  "github.com/gin-gonic/gin"
  "github.com/i159800/web3_homework/go_study_task/task4/controller"
  "github.com/i159800/web3_homework/go_study_task/task4/middleware"
)

func InitRouter(r *gin.Engine) {
  r.POST("/register", controller.Register)
  r.POST("/login", controller.Login)
  r.GET("/posts/:name", controller.GetPosts)
  auth := r.Group("/")
  auth.Use(middleware.JwtAuth())
  auth.POST("/posts", controller.CreatePost)
}
