package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RequestInfos2() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("请求Path：", path)
		fmt.Println("请求Method：", method)
		context.Next()
		fmt.Println(context.Writer.Status())
	}
}

func main() {
	engine := gin.Default()
	engine.Use(RequestInfos2()) //自定义中间件并使用

	engine.GET("/query", func(context *gin.Context) {
		fmt.Println("中间件的使用方法")
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  context.FullPath(),
		})
	})
	engine.Run("localhost:9000")
}

