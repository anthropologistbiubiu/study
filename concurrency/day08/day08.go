package main

import "sync"

// 今天写一个两个协程交替打印的程序

var bridge = make(chan int, 1)

func doWork1() {
	defer func() {
		wg.Done()
	}()
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			println(i)
			bridge <- i
		}
	}
}

func doWork2() {
	defer wg.Done()
	for i, j := range bridge {
		println(i, j)
	}
}

var wg = sync.WaitGroup{}

func main() {

	wg.Add(2)
	go doWork1()
	go doWork2()
	wg.Wait()
}
