package main

import "time"

func main() {

	var chan1 = make(chan int, 1)
	go func() {
		println("gorutine start")
		i := <-chan1
		println("gorutine end", i)
	}()
	time.Sleep(5 * time.Second)
	chan1 <- 5
	time.Sleep(5 * time.Second)
}
