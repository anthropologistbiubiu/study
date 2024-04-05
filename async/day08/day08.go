package main

import (
	"fmt"
	"sync"
)

// 今天写一个两个协程交替打印的程序

var bridge = make(chan int, 100)

func doWork1() {
	defer func() {
		wg.Done()
	}()
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			println("gorutine 1", i)
			bridge <- i
		}
	}
}

func doWork2() {
	defer wg.Done()
	for i := range bridge {
		println("gorutine 2", i+1)
		if i == 9 {
			return
		}
	}
}

// 我在这里要搞定channel 的用法

var wg = sync.WaitGroup{}

func main() {

	//wg.Add(2)
	//go doWork1()

	//go doWork2()

	//wg.Wait()
	for i := 0; i < 10; i++ {
		bridge <- i
	}
	for data := range bridge {
		fmt.Println(data)
	}
}
