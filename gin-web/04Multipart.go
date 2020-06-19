package main

import (
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User string  `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.User == "tom" && form.Password == "123"{
				c.JSON(200,gin.H{"status":"you are logged in"})
			}else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	r.Run("localhost:8080")
}