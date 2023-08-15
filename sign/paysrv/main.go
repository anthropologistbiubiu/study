package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.New()
	r.POST("/payment/create")
	r.POST("/payment/info")
	r.POST("/cashout/create")
	r.POST("/cashout/info")
	r.POST("/refound/create")
	r.POST("/refound/info")
	r.Run(":5555")

}
