package main

import "fmt"

func main() {
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
