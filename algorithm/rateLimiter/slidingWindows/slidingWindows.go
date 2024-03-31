package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

var (
	rdb = &redis.Client{}
	ctx = context.Background()
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // 设置密码，若没有密码则留空
		DB:       0,                // 使用默认数据库
	})
	// 检查连接是否成功
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()
	gin.ForceConsoleColor()
	r.Use(SlidingWindowsMiddleWare())
	r.POST("/rate", func(c *gin.Context) {
		c.String(http.StatusOK, "golang ~\n")
	})
	r.Run(":8080")
}

func SlidingWindowsMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		url := c.Request.URL.String()
		if !IsActionAllow(url, 1, 1) {
			fmt.Println("每秒该接口最大请求数为1,请检查请求频率")
			c.Abort()
		}
		c.Next()
	}
}

func IsActionAllow(actionKey string, period, maxCount int64) bool {

	key := "limiter:" + actionKey
	fmt.Println("key", key)
	//现在毫秒时间戳
	nowTime := time.Now().UnixNano() / 1e6
	//设置一个值，score和value都为毫秒时间戳
	mem := redis.Z{
		Score:  float64(nowTime),
		Member: nowTime,
	}
	rdb.ZAdd(context.Background(), key, mem)
	//移除时间窗口之前的值，剩下的就是当前的时间窗口
	Min := "-inf"
	Max := fmt.Sprintf("%d", nowTime-period*1000)
	result, err := rdb.ZRemRangeByScore(ctx, key, Min, Max).Result()
	fmt.Println("result", result, err)
	//获取时间窗口内的数量
	count, err := rdb.ZCard(ctx, key).Result()
	fmt.Println("count", count, err)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return count <= maxCount
}
