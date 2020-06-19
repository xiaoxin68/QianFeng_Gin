package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func StatCost()gin.HandlerFunc  {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	r := gin.New()
	r.Use(StatCost())

	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string)
		log.Println(name)
		c.JSON(http.StatusOK,gin.H{
			"message":"hello world",
		})
	})

	r.Run("localhost:8080")
}
