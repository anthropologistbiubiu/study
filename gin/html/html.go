package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("templates/**/*")
	r.LoadHTMLFiles("gin/templates/posts/index.html", "gin/templates/users/index.html", "gin/templates/earth/index.html")
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

func main2() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	router.LoadHTMLFiles("gin/html/temp/index.tmpl")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>")
	})
	router.Run(":8080")
}
