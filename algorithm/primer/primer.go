package main

import "fmt"

func gen(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func sieve(input <-chan int, output chan<- int, prime int) {
	for num := range input {
		if num%prime != 0 {
			output <- num
		}
	}
}

func main() {
	ch := make(chan int)
	go gen(ch)
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)
		newch := make(chan int)
		go sieve(ch, newch, prime)
		ch = newch
	}
}
