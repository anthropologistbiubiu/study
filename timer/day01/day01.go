package main

import (
	"fmt"
	"time"
)

func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(3 * time.Second)
	select {
	case <-conn:
		timer.Stop()
		return true
	case <-timer.C: //超时
		fmt.Println("WaitChannel timeout")
		return false
	}
}

func main() {

	conn := make(chan string, 1)
	//conn <- "sunweiming"
	flag := WaitChannel(conn)
	fmt.Println("flag", flag)

}
