package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	//gorm.Model是一个包含了ID, CreatedAt, UpdatedAt, DeletedAt四个字段的Golang结构体
	Id   int // 成为默认的主键
	Name string
	Age  int
}

func main() {
	dsn := "root:1.*&dhauisd01.@tcp(121.199.72.50:3306)/Gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DATABASE conn error: ", err)
		return
	}

	println(db.AutoMigrate(new(Student)))

}
