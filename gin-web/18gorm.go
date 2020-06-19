package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
type UserInfo struct {
	ID uint `gorm:"primary_key 'id'" json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Hobby string `json:"hobby"`
}

func main()  {
	db,err := gorm.Open("mysql","root:123456@tcp(localhost:3306)/iris?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//自动迁移
	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{1, "七米", "男", "篮球"}
	u2 := UserInfo{2, "沙河娜扎", "女", "足球"}

	//创建记录
	db.Create(&u1)
	db.Create(&u2)

	//查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu,"hobby=?","足球")
	fmt.Printf("%#v\n", uu)

	//更新
	db.Model(&u).Update("hobby","双色球")

	//删除
	db.Delete(&u)
}
