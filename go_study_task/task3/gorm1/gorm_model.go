package gorm1

import (
  "fmt"
  "gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Posts    []Post
	PostNum  uint   `gorm:"not null"`
}

type Post struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	UserId   uint
	Comments []Comment
	CommentStatus uint `gorm:"not null"`
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
  tx.Model(&User{}).Where("id = ?", p.UserId).Update("post_num", gorm.Expr("post_num + ?", 1))
  return
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	PostId  uint
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
  if c.ID == 0 {
    return
  }
  var postRow Post
  tx.First(&postRow, c.PostId)
  tx.Model(&Post{}).Preload("Comments").Find(&postRow)
  commentNum := len(postRow.Comments)
  fmt.Printf("postInfo postId:%d postCommentCount:%d\n", postRow.ID, commentNum)
  if commentNum <= 0 {
    postRow.CommentStatus = 0
    tx.Save(&postRow)
    fmt.Printf("postInfo postId:%d has no comment\n", postRow.ID)
  }
  return
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
  var post Post
  tx.Find(&post, c.PostId)
  if post.CommentStatus == 0 {
    post.CommentStatus = 1
    tx.Save(&post)
  }
  return
}
