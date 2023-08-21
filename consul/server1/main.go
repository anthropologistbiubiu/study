package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Print("hellowrld1")
	r := gin.New()

	r.POST("/hello", func(c *gin.Context) {
		fmt.Println("NNNNNNNNNNNN run8081")
		c.String(200, "success!")
	})

	r.Run(":8081")

	// 注册http 服务
	// gin框架

}
