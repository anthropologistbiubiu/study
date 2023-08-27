package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

var Limiter *rate.Limiter

func init() {
	Limiter = rate.NewLimiter(1000, 1500)
}

func GetToken() {

}

func Allow() {

}

func BlockWait() {

}

// 是需要写一个限流中间件的

// 在限流中间件中使用redis 存储当前限流器状态

// 并且使用ridis 分布式锁去完成对并发的控制

func test_tokenBucket() {
	// 创建一个令牌桶，每秒产生5个令牌
	limiter := rate.NewLimiter(5, 3)
	// 模拟一些请求
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %d processed\n", i+1)
		} else {
			fmt.Printf("Request %d dropped\n", i+1)
		}
		time.Sleep(500 * time.Millisecond) // 模拟请求间隔
	}
	for i := 0; i < 4; i++ {
		if limiter.Allow() {
			fmt.Printf("abnormal Request %d processed\n", i+1)
		} else {
			if err := limiter.Wait(context.Background()); err == nil {
				fmt.Printf("abnormal block Request %d processed\n", i+1)
			} else {
				fmt.Printf("abnormal Request %d dropped err:%+v \n", i+1, err)
			}
		}
	}
	for i := 0; i < 5; i++ {
		if limiter.Allow() {
			fmt.Printf("normal Request %d processed\n", i+1)
		} else if err := limiter.Wait(context.Background()); err == nil {
			fmt.Printf("normal Request %d processed\n", i+1)
		} else {
			fmt.Printf("normal Request %d dropped\n", i+1)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	test_tokenBucket()
}
