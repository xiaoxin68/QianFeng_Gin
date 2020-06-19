package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	userGroup := engine.Group("user")
	userGroup.GET("/register", registerHandle)
	userGroup.GET("/login", loginHandle)
	userGroup.GET("/info", infoHandle)
	engine.Run("localhost:9000")
}
func registerHandle(context *gin.Context) {
	fullPath := " 用户注册功能 " + context.FullPath()
	fmt.Println(fullPath)
	context.Writer.WriteString(fullPath)
}

func loginHandle(context *gin.Context) {
	fullPath := " 用户登录功能 " + context.FullPath()
	fmt.Println(fullPath)
	context.Writer.WriteString(fullPath)
}

func infoHandle(context *gin.Context) {
	fullPath := " 信息查看功能 " + context.FullPath()
	fmt.Println(fullPath)
	context.Writer.WriteString(fullPath)
}
