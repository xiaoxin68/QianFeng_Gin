package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:9000") // 监听并在 0.0.0.0:8080 上启动服务
}
