package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
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

func main() {
	r := gin.Default()
	gin.ForceConsoleColor()
	r.Use(RateLimitMiddleware(time.Second, 100, 100))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "golang ~")
	})
	r.Run(":8080")
}

// zset 滑动窗口实现的限流函数   传入的参数应该是paychan value成员就不用传 用时间戳来就行  这个函数可以做到的就是对某个用户来限制请求

// 限制对某个接口个请求

func IsActionAllow(userId, actionKey string, period, maxCount int) bool {
	//用userId和actionKey拼接成zset的key
	key := "limiter:" + userId + ":" + actionKey
	//现在毫秒时间戳
	nowTime := time.Now().UnixNano() / 1e6
	//设置一个值，score和value都为毫秒时间戳
	redis.zadd(key, nowTime, nowTime)
	//移除时间窗口之前的值，剩下的就是当前的时间窗口
	redis.zremrangeByScore(key, 0, nowTime-period*1000)
	//获取时间窗口内的数量
	count := redis.zcard(key)
	return count <= maxCount
}

// 这段程序应该走那段路径才是对的

func main1() {
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
