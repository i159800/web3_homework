package sqlx1

import (
  "fmt"
  "github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
  database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog")
  if err != nil {
    fmt.Println("open mysql failed,", err)
    return
  }
  Db = database
}
