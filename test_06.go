package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("请求Path：", path)
		fmt.Println("请求Method：", method)
	}
}

func main() {
	engine := gin.Default()
	engine.Use(RequestInfos()) //自定义中间件并使用

	engine.GET("/query", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  context.FullPath(),
		})
	})
	engine.Run("localhost:9000")
}
