package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"QianFeng_Gin/model"
)



func main() {
	engin := gin.Default()

	//http://localhost:9000/hello?username=张三&phone=1313&password=hfdsaf
	//ShouldBindQuery : 使用ShouldBindQuery可以实现Get方式的数据请求的绑定
	engin.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var user model.UserRegister

		err := context.ShouldBindQuery(&user)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(user.Username)
		fmt.Println(user.Phone)
		fmt.Println(user.Password)

		context.Writer.Write([]byte("hello," + user.Username))
	})

	//使用ShouldBind可以实现Post方式的提交数据的绑定工作。
	engin.Handle("POST", "/redister", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var register model.UserRegister
		if err := context.ShouldBind(&register); err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(register.Username)
		fmt.Println(register.Phone)
		fmt.Println(register.Password)

		context.Writer.Write([]byte("hello," + register.Username))
	})

	//ShouldBindJson : 当客户端使用Json格式进行数据提交时，可以采用ShouldBindJson对数据进行绑定并自动解析
	engin.Handle("POST", "/add", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var register model.UserRegister
		if err := context.BindJSON(&register); err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(register.Username)
		fmt.Println(register.Phone)
		fmt.Println(register.Password)

		context.Writer.Write([]byte("hello," + register.Username))
	})
	engin.Run("localhost:9000")
}
