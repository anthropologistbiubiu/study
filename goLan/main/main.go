package main

import (
	"fmt"
	"sync"
)

func concurrency_test1() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go func() {
		close(chan2)
		fmt.Println(<-chan2)
	}()
	go func() {
		close(chan1)
		fmt.Println(<-chan1)
	}()
	select {
	case <-chan1:
		fmt.Println("chan1 ready.", <-chan1)
	case <-chan2:
		fmt.Println("chan2 ready.")
	}
	fmt.Println("main exit.")
}

// 两个协程交替打印出100
func concurrency_test2() {
	ch1 := make(chan struct{}, 1)
	var wait sync.WaitGroup
	wait.Add(2)
	ch1 <- struct{}{}
	var i int = 1
	go func() {
		defer wait.Done()
		for i <= 9 {
			<-ch1
			fmt.Println("gorutine1", i)
			i++
			ch1 <- struct{}{}
		}
	}()
	go func() {
		defer wait.Done()
		for i <= 9 {
			<-ch1
			fmt.Println("gorutine1", i)
			i++
			ch1 <- struct{}{}
		}
	}()
	wait.Wait()
}

func main() {
	ch := make(chan struct{}, 1)
	var wait sync.WaitGroup
	wait.Add(2)
	ch <- struct{}{}
	go func() {

	}()
	go func() {

	}()
	select {}
}
