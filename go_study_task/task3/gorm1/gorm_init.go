package gorm1

import (
  "database/sql"
  "fmt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

var GormDb *gorm.DB

func init() {
  Db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog2?charset=utf8&parseTime=True&loc=Local")
  gormDB, err := gorm.Open(mysql.New(mysql.Config{
    Conn: Db,
  }), &gorm.Config{})
  if err != nil {
    fmt.Println("mysql conn failed: ", err)
    return
  }
  GormDb = gormDB
}
