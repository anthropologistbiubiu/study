package main

import (
	"fmt"
	"sync"
	"time"
)

/*
可以先实现伪代码
要求每个一分钟间隔内的最大请求为10
固定窗口实现
*/

type Limiter struct {
	maxRequest int
	count      int
	lastTime   int64
	duration   time.Duration
	lock       sync.Mutex
}

func Init(maxRequest int, duration time.Duration) *Limiter {
	return &Limiter{
		maxRequest: maxRequest,
		lastTime:   time.Now().Unix(),
		duration:   duration,
		lock:       sync.Mutex{},
	}
}
func (l *Limiter) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.count++
	if time.Now().Unix()-l.lastTime > int64(l.duration) {
		l.lastTime = time.Now().Unix()
		l.count = 1
		return true
	}
	if l.count <= l.maxRequest {
		return true
	}
	return false
}

func main() {
	// 模拟并发请求 实际上还是串行执行
	limiter := Init(3, 5)
	i := 0
	for range time.Tick(time.Second) {
		i++
		if limiter.Allow() {
			fmt.Printf("reuest %d is  allowed\n", i)
		} else {
			fmt.Printf("reuest %d is not allowed\n", i)
		}
	}
}
