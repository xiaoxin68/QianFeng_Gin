package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string `gorm:"varchar(20);not null"`
	TelePhone string `gorm:"varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`
}

func main() {
	db,err := gorm.Open("mysql","root:123456@(localhost:3306)/ginessential?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	user := User{Name: "q1mi", TelePhone: "231",Password: "12312"}


	fmt.Println(db)
	fmt.Println(user)

	db.Create(&user)   // 创建user
}

