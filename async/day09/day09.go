package main

import (
	"time"
)

func main() {

	var workChan = make(chan int)
	go func() {
		println("gorutine start")
		workChan <- 1
		println("gorutine end")
	}()
	//<-workChan
	println("end")
	time.Sleep(2 * time.Second)
}

var switchChan = make(chan struct{})
var dataChan = make(chan int, 1)

func work1() {

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			println("work1", i)
			switchChan <- struct{}{}
		} else {
			continue
		}
		//dataChan<-1
	}
}

func work2() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			println("work2", i)
			switchChan <- struct{}{}
		} else {
			continue
		}
	}
}
