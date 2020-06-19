package main

import (
	"QianFeng_Gin/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	connStr := "root:123456@tcp(localhost:3306)/iris"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//创建表
	//_, err = db.Exec("create table person(" +
	//	"id int auto_increment primary key," +
	//	"name varchar(12) not null," +
	//	"age int default 1" +
	//	");")
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return
	//}

	//插入数据
	//_, err = db.Exec("insert into person(name,age) "+
	//	"values(?,?);", "Lily", 15)
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return
	//} else {
	//	fmt.Println("数据插入成功")
	//}

	//查询数据库记录
	rows, err := db.Query("select id,name,age from person")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	columns, err := rows.Columns()
	fmt.Println(columns)

	scan:
	if rows.Next() {
		person := new(model.Person)
		err = rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Id, person.Name, person.Age)
		goto scan
	}
}
