package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page","0")
		name := c.PostForm("name")
		message := c.DefaultPostForm("message","你好")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
	})
	router.Run("localhost:8080")
}
