package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	r.POST("/hello", func(c *gin.Context) {
		fmt.Println("NNNNNNNNN run:8082")
		c.JSON(200, gin.H{
			"message": "sunweiming",
		})
	})
	r.Run(":8082")
}
