package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engin := gin.Default() //gin.Default()和gin.New()的区别在于gin.Default也使用gin.New()创建engine实例，但是会默认使用Logger和Recovery中间件。

	//http://localhost:9000/hello?name=张三
	engin.Handle("GET", "/hello", func(context *gin.Context) {
		//获取请求接口
		fmt.Println(context.FullPath())
		//获取字符串参数
		name := context.DefaultQuery("name", "") //可以通过context.Query和context.DefaultQuery获取GET请求携带的参数
		fmt.Println(name)

		pwd := context.Query("pwd")
		fmt.Println(pwd)

		//输出
		context.Writer.Write([]byte("hello ," + name)) //可以通过context.Writer.Write向请求发起端返回数据。
	})

	engin.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		username := context.PostForm("username")
		fmt.Println(username)

		/*password := context.PostForm("password")
		fmt.Println(password)*/
		password,exists:= context.GetPostForm("password")
		if exists {
			fmt.Println(password)
		}

		context.Writer.Write([]byte("User login"))
	})

	//通过路由的:id来定义一个要删除用户的id变量值，同时使用context.Param进行获取。
	engin.DELETE("/user/:id",DeleteHandle)

	engin.Run("localhost:9000")
}

func DeleteHandle(context *gin.Context) {
	fmt.Println(context.FullPath())
	userID := context.Param("id")
	fmt.Println(userID)

	context.Writer.Write([]byte("Delete user's id : " + userID))
}
