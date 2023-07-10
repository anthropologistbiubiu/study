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

func DelayFunction() {
	timer := time.NewTimer(5 * time.Second)
	select {
	case <-timer.C:
		fmt.Println("Delayed 5s，...")
	}
}

func AfterDemo() {
	fmt.Println(time.Now().Second())
	<-(time.After(1 * time.Second))
	fmt.Println(time.Now().Second())
}

func AfterFuncDemo() {
	fmt.Println("AfterFuncDemo start", time.Now().Format("20060102"))
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("AfterFuncDemo end", time.Now().Format("20060102 13:14:15"))
	})
	fmt.Println("wait")
	time.Sleep(4 * time.Second) //等待协程退出
}

func WorngTicker() {
	for {
		select {
		case <-time.Tick(1 * time.Second):
			fmt.Println("资源泄露")
		}
	}
}
func demo(t interface{}) {
	for {
		select {
		case <-t.(*time.Ticker).C:
			println("1s timer")
		}
	}
}

func main1() {
	t := time.NewTicker(time.Second * 1)
	go demo(t)
	select {}
}

func main() {

	//conn := make(chan string, 1)
	//conn <- "sunweiming"
	//flag := WaitChannel(conn)
	//fmt.Println("flag", flag)
	//DelayFunction()
	//AfterDemo()
	AfterFuncDemo() //time.AfterFunc()是异步执行的，
	// 所以需要在函数最后sleep等待指定的协程退出，否则可能函数结束时协程还未执行。
	main1()
}
