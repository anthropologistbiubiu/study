package main

import (
	"runtime"
	"time"
)

// 用两个无缓冲管道  实现两个协程交替打印
var chan1 = make(chan struct{}, 1)
var chan2 = make(chan struct{}, 1)

func work1() {

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			<-chan2
			println("work1", i)
			chan1 <- struct{}{}
		} else {
			continue
		}
		//dataChan<-1
	}
}

func work2() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			<-chan1
			println("work2", i)
			chan2 <- struct{}{}
		} else {
			continue
		}
	}
}

// 验证一下无缓冲管道和有缓冲管道 buff=1 在并发控制中的不同
func main() {

	runtime.GOMAXPROCS(2)
	chan2 <- struct{}{}
	go work1()
	go work2()

	time.Sleep(5 * time.Second)
}
