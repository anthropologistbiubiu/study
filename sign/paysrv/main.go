package main

import "github.com/gin-gonic/gin"

// 这里主要是暴露一个服务
func main() {
	r := gin.New()
	r.POST("/payment/create")
	r.POST("/payment/info")
	r.POST("/cashout/create")
	r.POST("/cashout/info")
	r.POST("/refound/create")
	r.POST("/refound/info")
	r.POST("/payment/notify")
	r.POST("/cashout/notify")
	r.Run(":5555")
}
