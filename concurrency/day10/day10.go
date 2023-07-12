package main

import "time"

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

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			println("work2", i)
			<-switchChan
		} else {
			continue
		}
	}
}

func main() {

	go work1()
	go work2()
	time.Sleep(5 * time.Second)
}
