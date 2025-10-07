package sqlx1

import "fmt"

type Book struct {
  Id  int `db:"id"`
  Title string `db:"title"`
  Author string `db:"author"`
  Price float64 `db:"price"`
}

func SqlxBooksInsert() {
  r, err := Db.Exec("insert into books (title, author, price) values (?, ?, ?)", "book4", "author4", 90.0)
  if err != nil {
    fmt.Println("books insert failed: ", err)
    return
  }
  id, err := r.LastInsertId()
  if err != nil {
    fmt.Println("books insert failed: ", err)
    return
  }
  fmt.Println("books insert success: ", id)
}

func SqlxBooksRows() {
  var books []Book
  err := Db.Select(&books, "select id,title,author,price from books where price > ?", 50)
  if err != nil {
    fmt.Println("books select failed: ", err)
    return
  }
  fmt.Println("books select success: ", books)
}

