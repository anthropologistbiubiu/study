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
func main() {
	c := make(chan int)
	go counter(c)
	for i := 0; i < 11; i++ {
		p := <-c
		fmt.Println(p)
		primes := make(chan int)
		go filter(c, primes, p)
		c = primes
	}
}
