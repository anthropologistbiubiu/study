package main

import (
	"fmt"
	"time"
)

type LeakyBucket struct {
	capacity   int // 漏桶容量
	rate       int // 漏桶速率，每秒处理请求数
	waterLevel int // 漏桶中当前的水量
}

func NewLeakyBucket(capacity, rate int) *LeakyBucket {
	return &LeakyBucket{
		capacity:   capacity,
		rate:       rate,
		waterLevel: 0,
	}
}

func (lb *LeakyBucket) AddRequest() bool {
	if lb.waterLevel < lb.capacity {
		lb.waterLevel++
		return true
	}
	return false
}

func (lb *LeakyBucket) Drain() {
	for range time.Tick(time.Second) {
		if lb.waterLevel > 0 {
			lb.waterLevel--
		}
	}
}

// 漏桶如何实现1分钟内处理最多600个请求的需求

// 漏桶如何实现1分钟内处理最多1个请求的需求
func main() {
	// 创建一个漏桶，容量为10，速率为1（每秒处理1个请求）
	leakyBucket := NewLeakyBucket(3, 1)
	// 启动漏桶的漏水任务
	go leakyBucket.Drain()
	// 模拟请求到达
	for i := 0; i < 20; i++ {
		// 当漏桶满了时，拒绝新的请求
		if !leakyBucket.AddRequest() {
			fmt.Println("请求被拒绝：漏桶已满")
		} else {
			fmt.Println("请求被接受：漏桶中的水量为", leakyBucket.waterLevel)
		}
		time.Sleep(time.Millisecond * 500) // 模拟请求间隔
	}
}
