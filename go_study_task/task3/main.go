package main

import (
  "github.com/i159800/web3_homework/go_study_task/task3/gorm1"
  "github.com/i159800/web3_homework/go_study_task/task3/sqlx1"
)

func main() {
  //使用sqlx查询employees表中工资最高的员工信息，并将结果映射到一个Employee结构体中。
  sqlx1.SqlxEmployeeInsert()
  sqlx1.SqlxEmployeeSelectRow()
  //查询价格大于50元的书籍，并将结果映射到Book结构体切片中，确保类型安全。
  sqlx1.SqlxBooksInsert()
  sqlx1.SqlxBooksRows()
  //gorm1.GormProductInsert()
  //编写go代码，使用gorm创建这些模型对应的数据库表。
  gorm1.GormUserInsert()
  //使用gorm查询某个用户发布的所有文章及其对应的评论信息
  gorm1.GormSelectAllByUserName("wangwu")
  //使用gorm查询评论数量最多的文章信息
  gorm1.GormSelectMostCommentsPost()
  //为Post模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
  gorm1.GormPostsAdd(2, "title_22", "content_22")
  //为Comment模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为0，则更新文章的评论状态为无评论。
  //gorm1.GormAddComment(9, "comment91")
  //gorm1.GormDeleteComment(23)
}
