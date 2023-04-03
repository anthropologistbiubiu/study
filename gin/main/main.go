package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

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
	router.LoadHTMLFiles("gin/temp/index.tmpl")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "")
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

func main4() {
	r := gin.Default()

	// ①
	r.Static("/assets", "assets")
	// ②
	//r.StaticFS("/assets", http.Dir("assets"))
	// ③
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	r.Run(":8080")
}

// gin 注册路由中间件
func costTime() gin.HandlerFunc {
	return func(c *gin.Context) {

		now := time.Now()
		c.Next()
		fmt.Printf(" reuqestUrl is %+v costTime is %+v\n", c.Request.URL.String(), time.Since(now))
	}
}
func main6() {

	r := gin.Default()
	r.Use(costTime())
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "sunweiming",
		})
	})
	r.Run("127.0.0.1:8080")
}

// gin 中间件执行原理分析
type Handler func(gc *GContext)
type HanclerChain []Handler

type GContext struct {
	Handlers HanclerChain
	index    int
}

func (g *GContext) Next() {
	g.index++
	for g.index < len(g.Handlers) {
		g.Handlers[g.index](g)
		g.index++
	}
}

func (g *GContext) Use(handlers ...Handler) {
	g.Handlers = append(g.Handlers, handlers...)
}

func (g *GContext) Handle(hs ...Handler) {
	g.Handlers = append(g.Handlers, hs...)
}
func (g *GContext) Start() {
	g.Next()
}
func one(gc *GContext) {
	fmt.Println("one")
	gc.Next()
	fmt.Println("one-afer")
}
func two(gc *GContext) {
	fmt.Println("two")
	gc.Next()
	fmt.Println("two-afer")
}
func three(gc *GContext) {
	fmt.Println("three")
	gc.Next()
	fmt.Println("three-afer")
}
func four(gc *GContext) {
	fmt.Println("four")
	gc.Next()
	fmt.Println("four-afer")
}
func main7() {
	gc := &GContext{
		index:    -1,
		Handlers: HanclerChain{},
	}
	gc.Use(one, two)
	gc.Handle(
		func(gc *GContext) {
			fmt.Println("main process")
		},
	)
	//gc.Use(three, four)
	gc.Start()
}

func main() {

	r := gin.Default()
	r.GET("/index/:id", func(c *gin.Context) {
		//id := c.Query("id")
		//id, ok := c.GetQuery("id")
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
			//"ok": ok,
		})
	})
	r.Run(":8080")
}
