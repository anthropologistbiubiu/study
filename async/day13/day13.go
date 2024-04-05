package main

import "time"

func main() {

	var chan1 = make(chan int, 10)

	for i := 1; i <= 10; i++ {
		chan1 <- i
	}

	go func() {
		for i := 1; i <= 11; i++ {
			j := <-chan1
			println(j)
		}
	}()

	/*
		go func() {
			for j := range chan1 {
				println(j)
			}
		}()
	*/
	// 哇神奇 ，在协程中处理管道 ，竟然没有报错
	k := <-chan1
	println("k", k)
	time.Sleep(5 * time.Second)
}
