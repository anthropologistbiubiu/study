package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main1() {
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

func main3() {
	router := gin.Default()
	// 设置静态资源文件目录，并且绑定一个Url前缀
	// 静态资源文件目录：/var/www/tizi365/assets
	// /assets是访问静态资源的url前缀
	// 例如：
	//   /assets/images/1.jpg 这个url文件，存储在/var/www/tizi365/assets/images/1.jpg
	router.Static("/assets", "/var/www/tizi365/assets")

	// 为单个静态资源文件，绑定url
	// 这里的意思就是将/favicon.ico这个url，绑定到./resources/favicon.ico这个文件
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

func main() {
	r := gin.Default()

	// ①
	r.Static("/assets", "assets")
	// ②
	//r.StaticFS("/assets", http.Dir("assets"))
	// ③
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	r.Run(":8080")
}
