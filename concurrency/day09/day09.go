package main

import (
	"fmt"
	"time"
)

func main() {

	var workChan = make(chan int)
	go func() {
		fmt.Println("gorutine start")
		workChan <- 1
		println("gorutine end")
	}()
	//<-workChan
	println("end")
	time.Sleep(2 * time.Second)
}
