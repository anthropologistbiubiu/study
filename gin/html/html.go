package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("templates/**/*")
	r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html", "templates/earth/index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	r.GET("earth/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "earth/index.html", gin.H{
			"title": "earth/index",
		})
	})
	r.Run(":8080")
}
