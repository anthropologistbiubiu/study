package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

// 在这里定义一个全局的token bucket 限流算法 实现对服务的限流保护

func main() {
	// 创建一个令牌桶，每秒产生5个令牌
	limiter := rate.NewLimiter(5, 1)
	// 模拟一些请求
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %d processed\n", i+1)
		} else {
			fmt.Printf("Request %d dropped\n", i+1)
		}
		time.Sleep(200 * time.Millisecond) // 模拟请求间隔
	}
}
