package gorm1

import "fmt"

func GormUserInsert() {
  user := User{
    Username: "wangwu",
    Email: "wangwu@163.com",
    Posts: []Post{
      {Title: "post_title7", Content: "post_content7", Comments: []Comment{
        {Content: "comment71"},
        {Content: "comment72"},
        {Content: "comment73"},
      }},
      {Title: "post_title8", Content: "post_content8", Comments: []Comment{
        {Content: "comment81"},
        {Content: "comment82"},
        {Content: "comment83"},
        {Content: "comment84"},
      }},
      {Title: "post_title9", Content: "post_content9", Comments: []Comment{
        {Content: "comment91"},
        {Content: "comment92"},
        {Content: "comment93"},
        {Content: "comment94"},
        {Content: "comment95"},
      }},
    },
  }

  GormDb.AutoMigrate(&User{}, &Post{}, &Comment{})
  GormDb.Create(&user)
}

//根据某个用户名字查询发布的所有文章和对应的评论信息
func GormSelectAllByUserName(userName string) {
  var users []User
  GormDb.Model(&User{}).Preload("Posts").Where("username = ?", userName).Find(&users)
  for _, userRow := range users {
    fmt.Printf("userInfo: userId:%d username:%s email:%s\n", userRow.ID, userRow.Username, userRow.Email)
    for _, postRow := range userRow.Posts {
      fmt.Printf("\tpostInfo: postId:%d userId:%d postTitle:%s postConent:%s\n", postRow.ID, postRow.UserId, postRow.Title, postRow.Content)
      GormDb.Model(&Post{}).Preload("Comments").Find(&postRow)
      for _, commentRow := range postRow.Comments {
        fmt.Printf("\t\tcommentInfo: commentId:%d commentPostId:%d commentContent:%s\n", commentRow.ID, commentRow.PostId, commentRow.Content)
      }
    }
    fmt.Println()
  }
}

//获取评论最多的文章信息
func GormSelectMostCommentsPost() {
  var posts []Post
  var maxCommentCount int64
  var maxCommentPostId uint
  GormDb.Model(&Post{}).Find(&posts)
  for _, postRow := range posts {
    commentCount := GormDb.Model(&postRow).Association("Comments").Count()
    if (commentCount > maxCommentCount) {
      maxCommentCount = commentCount
      maxCommentPostId = postRow.ID
    }
    fmt.Printf("postId:%d commentCount:%d\n", postRow.ID, commentCount)
  }

  fmt.Println()
  fmt.Printf("most comments of post info: postId:%d commentCount:%d \n", maxCommentPostId, maxCommentCount)
}

//钩子函数
func GormPostsAdd(uid uint, postTitle, postContent string) {
  post := Post{
    Title: postTitle,
    Content: postContent,
    UserId: uid,
  }
  GormDb.AutoMigrate(&User{}, &Post{})
  GormDb.Create(&post)
  var users []User
  GormDb.Model(&User{}).Where("id = ?", uid).Find(&users)
  for _, userRow := range users {
    fmt.Printf("userInfo uid:%d post_num:%d\n", userRow.ID, userRow.PostNum)
  }
}

func GormDeleteComment(commentId uint) {
  GormDb.AutoMigrate(&Post{}, &Comment{})
  var comment Comment
  GormDb.Where("id = ?", commentId).Find(&comment).Delete(&comment)
}

func GormAddComment(postId uint, commentContent string) {
  comment := Comment{
    Content: commentContent,
    PostId: postId,
  }
  GormDb.Create(&comment)
  var postRow Post
  GormDb.First(&postRow, postId)
  GormDb.Model(&Post{}).Preload("Comments").Find(&postRow)
  commentIds := []uint{}
  for _, commentRow := range postRow.Comments {
    commentIds = append(commentIds, commentRow.ID)
  }
  fmt.Printf("postInfo postId:%d commentIds:%v\n", postRow.ID, commentIds)
}
