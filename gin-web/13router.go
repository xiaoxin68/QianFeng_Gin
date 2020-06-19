package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
	 fmt.Println("test")
	})

	//为/demo的所有请求匹配
	r.Any("/demo", func(c *gin.Context) {
		fmt.Println("demo")
	})

	//为没有匹配到的方法
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")
	})
	r.Run("localhost:8080")
}
