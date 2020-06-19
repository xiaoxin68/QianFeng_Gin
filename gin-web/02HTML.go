package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../html/*")
	r.Static("/img", "../img")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{
			"title": "Main website",
			"fullpath" : c.FullPath(),
		})
	})
	r.Run("localhost:8080")
}
