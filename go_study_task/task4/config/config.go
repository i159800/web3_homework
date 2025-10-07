package config

import (
  "github.com/i159800/web3_homework/go_study_task/task4/model"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "log"
)

var DB *gorm.DB

func InitDB() {
  dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatalf("failed to connect database: %v", err)
  }
  DB = db
  db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
}
