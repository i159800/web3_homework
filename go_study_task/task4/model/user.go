package model

import "gorm.io/gorm"

type User struct {
  gorm.Model
  Username string `gorm:"unique;not null" uri:"name"`
  Password string `gorm:"not null"`
  Email    string `gorm:"unique;not null"`
  Posts    []Post
  //PostNum  uint   `gorm:"not null"`
}
