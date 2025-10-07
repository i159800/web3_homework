package gorm1

import (
  "fmt"
  "gorm.io/gorm"
)

type Product struct {
  gorm.Model
  Code string
  Price uint
}

func GormProductInsert() {
  //Migrate the schema
  GormDb.AutoMigrate(&Product{})

  //Create
  //GormDb.Create(&Product{Code: "D42", Price: 100})

  //Read
  var product Product
  GormDb.First(&product, 1) //根据主键查找
  fmt.Println("products select: ", product)

  //Update 将product的price更新为200
  GormDb.Model(&product).Update("Price", 200)
  fmt.Println("products select: ", product)

  //Read2
  GormDb.First(&product, "code = ?", "D42")
  fmt.Println("products select: ", product)

  //Delete
  GormDb.Delete(&product, 1)
}
