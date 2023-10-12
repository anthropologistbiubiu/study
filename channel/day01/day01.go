package main

import (
	"fmt"
	"time"
)

// 熟悉一下channel 的属性
func main() {

	channel := make(chan int, 10)
	for i := 0; i < 10; i++ {
		channel <- i
	}
	//close(channel)
	go func() {
		for v := range channel {
			fmt.Println(v)
		}
	}()
	time.Sleep(1 * time.Second)
}
