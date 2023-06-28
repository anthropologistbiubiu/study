package main

import (
	"fmt"
	"time"
)

func demo(count int) {
	for i := 1; i < 10; i++ {
		fmt.Println(count, ":", i)
	}
}

func main() {
	for i := 1; i < 10; i++ {
		go demo(i)
	}
	//添加休眠时间等待goroutine执行结束
	time.Sleep(3e9)
}
