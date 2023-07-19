// 素数筛子
// 一般写法
package main

import "fmt"

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
func counter(c chan int) {
	i := 2
	for {
		c <- i
		i++
	}
}

/*
	func main() {
		c := make(chan int)
		go counter(c)
		primes := make(chan int)
		go filter(c, primes, 2)
		for i := 0; i < 10; i++ {
			p := <-primes
			fmt.Println(p)
		}
	}
*/
func main() {
	c := make(chan int)
	go counter(c)
	primes := make(chan int)
	fmt.Println(2)
	go filter(c, primes, 2)
	i := <-primes
	fmt.Println(i)
	c = primes
	primes = make(chan int)
	go filter(c, primes, i)
	for i := 0; i < 10; i++ {
		p := <-primes
		fmt.Println(p)
	}
}
