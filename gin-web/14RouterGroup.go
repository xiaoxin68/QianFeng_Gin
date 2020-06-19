package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//路由组
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			fmt.Println("index")
		})

		userGroup.GET("/login", func(c *gin.Context) {
			fmt.Println("login")
		})

		userGroup.GET("/post", func(c *gin.Context) {
			fmt.Println("post")
		})

		//嵌套路由组
		v1 := userGroup.Group("/v1")
		v1.GET("/oo", func(c *gin.Context) {
			fmt.Println("oo")
		})
	}

	r.Run("localhost:8080")
}
