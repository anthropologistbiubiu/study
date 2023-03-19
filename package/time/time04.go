package main

import (
	"fmt"
	"log"
	"time"
)

func demo(t interface{}) {
	for {
		select {
		case <-t.(*time.Ticker).C: //断言
			println("ls timer")
		}

	}
}

// bad
func WorngTicker() {
	for {
		select {
		case <-time.Tick(1 * time.Second):
			log.Println("资源泄露")
		}
	}
}

func AfterDemo() {
	log.Println(time.Now)
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println(" waiting for a second!", time.Now())
		}
	}
}
func AfterFuncDemo() {
	log.Println("AfterFuncDemo start", time.Now())
	time.AfterFunc(1*time.Second, func() {
		log.Println("AfterFuncDemo end", time.Now())
	})

	time.Sleep(2 * time.Second) //等待协程退出
}

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

// good
func main() {
	conn := make(chan string)
	WaitChannel(conn)
}
