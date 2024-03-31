package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func RateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusForbidden, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}
func tokenBucket() {
	fillInterval := 10 * time.Microsecond        //添加令牌的时间间隔
	capacity := 100                              //令牌桶的容量
	tokenBucket := make(chan struct{}, capacity) //初始化一个令牌桶

	//每隔一秒钟往令牌桶里添加令牌，如果桶已经满了，则直接放弃
	fillToken := func() {
		ticker := time.NewTicker(fillInterval)
		for {
			select {
			case <-ticker.C:
				select {
				case tokenBucket <- struct{}{}:
				default:
				}

			}
		}
	}
	go fillToken()
	select {}
}

func main() {
	r := gin.Default()
	gin.ForceConsoleColor()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "golang ~")
	})
	r.Run(":8080")
}
