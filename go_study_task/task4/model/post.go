package model

import "gorm.io/gorm"

type Post struct {
  gorm.Model
  Title    string `gorm:"not null"`
  Content  string `gorm:"not null"`
  UserId   uint
  Comments []Comment
  //CommentStatus uint `gorm:"not null"`
}
