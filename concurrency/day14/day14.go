// 素数筛子
// 一般写法
package main

import (
	"fmt"
)

func generate(ch chan int) {
	for i := 2; i < 100; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
