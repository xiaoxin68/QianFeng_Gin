package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"QianFeng_Gin/model"
	"net/http"
)

func main() {
	engin := gin.Default()

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

		//[]byte格式
		context.Writer.Write([]byte("hello," + user.Username))

		//字符串格式
		context.Writer.WriteString(user.Username + " " + user.Phone + " " + user.Password)

		//json格式
		context.JSON(200, user)

		context.JSON(201, map[string]interface{}{
			"code":    1,
			"message": "OK",
			"data":    user,
		})

		//html格式
		engin.LoadHTMLGlob("./html/*") //设置html的目录
		//加载静态资源文件
		engin.Static("/img", "./img") //*************
		engin.GET("/hellohtml", func(context *gin.Context) {
			fullPath := "请求路径:" + context.FullPath()

			context.HTML(http.StatusOK, "index.html", gin.H{
				"title":    "Gin教程",
				"fullpath": fullPath,
			})
		})
	})
	engin.Run("localhost:9000")
}
