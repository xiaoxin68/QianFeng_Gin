package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//重定向
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")
	})

	//路由重定向
	r.GET("/demo", func(c *gin.Context) {
		c.Request.URL.Path = "/demo2"
		r.HandleContext(c)
	})

	r.GET("/demo2", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"hello":"world"})
	})
	r.Run("localhost:8080")
}