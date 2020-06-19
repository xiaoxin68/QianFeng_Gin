package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	ID   int64
	Name string `gorm:"default:'小王子'"`
	Age  int64
}

func (Person) TableName() string {
	return "person"
}

func main() {
	db, _ := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/iris?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()

	//插入
	//person := Person{Name: "张三",Age: 15}
	//db.NewRecord(person)
	//db.Create(&person)
	//db.NewRecord(person)

	//查询
	var person Person

	//db.First(&person)

	//db.Take(&person)

	//db.Last(&person)

	//db.Find(&person, 2)

	//db.Where("name=?", "张三").First(&person)

	db.Where(&Person{Name: "张三", Age: 15}).First(&person)

	fmt.Println(person)

	var persons []Person

	//db.Find(&persons)

	//db.Where("name = ?","张三").Find(&persons)

	//db.Where("name <> ?","张三").Find(&persons)

	//db.Where("name in (?) ",[]string{"张三", "李四","tom"}).Find(&persons)

	//db.Where("name like ? ","%三").Find(&persons)

	//db.Where("name = ? AND age >= ?", "张三", 16).Find(&persons)

	//db.Where("updated_at > ?", lastWeek).Find(&users)

	//db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)

	// Map
	//db.Where(map[string]interface{}{"name": "张三", "age": 15}).Find(&persons)

	// 主键的切片
	db.Where([]int64{1, 3, 4}).Find(&persons)
	//// SELECT * FROM users WHERE id IN (20, 21, 22);

	fmt.Println(persons)
}
