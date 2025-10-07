package sqlx1

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

type Employee struct {
  Id int `db:"id"`
  Name string `db:"name"`
  Department string `db:"department"`
  Salary float64 `db:"salary"`
}

func SqlxEmployeeInsert() {
  r, err := Db.Exec("insert into employees(name, department, salary) values (?, ?, ?)", "zhangsan", "depart03", 50000.0)
  if err != nil {
    fmt.Println("exec failed, ", err)
    return
  }
  id, err := r.LastInsertId()
  if err != nil {
    fmt.Println("exec failed, ", err)
    return
  }

  fmt.Println("insert succ:", id)
}

func SqlxEmployeeSelectRows() {
  var employees []Employee
  err := Db.Select(&employees, "select id, name, department, salary from employees where id = ?", 1)
  if err != nil {
    fmt.Println("exec failed, ", err)
    return
  }

  fmt.Println("select succ:", employees)
}

func SqlxEmployeeSelectRow() {
  var employee Employee
  err := Db.Get(&employee, "select id, name, department, salary from employees order by salary desc limit 1")
  if err != nil {
    fmt.Println("exec failed, ", err)
    return
  }

  fmt.Println("select succ:", employee)
}
