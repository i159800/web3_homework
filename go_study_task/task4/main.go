package main

import (
  "github.com/gin-gonic/gin"
  "github.com/i159800/web3_homework/go_study_task/task4/config"
  "github.com/i159800/web3_homework/go_study_task/task4/router"
)

func main() {
  config.InitDB()
  r := gin.Default()
  router.InitRouter(r)
  r.Run(":8080")
}
