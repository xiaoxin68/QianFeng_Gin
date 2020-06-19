package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type student struct {
	gorm.Model
	Name string
	Age sql.NullInt64
	Birthday *time.Time
	Email string `gorm:"type:varchar(100);unique_index"`
	Role  string  `gorm:"size:255"` // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"` // 忽略本字段
}

//主键（Primary Key） : GORM 默认会使用名为ID的字段作为表的主键
//type User struct {// // 默认表名是 `users`
//	Id string // 名为`ID`的字段会默认作为表的主键
//	Name string
//}

// 使用`AnimalID`作为主键
type Animal struct {
	AnimalId int64 `gorm:"primary_key"`
	Name string
	Age int
}

//表名默认就是结构体名称的复数!!!!
//将user的表名设置为`profiles`
func (User) TableName() string  {
	return "profiles"
}

//func (u User)TableName()string  {
//	if u.Name == "admin" {
//		return "admin_users"
//	} else {
//		return "users"
//	}
//}
//// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
//db.SingularTable(true)