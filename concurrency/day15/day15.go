package main

func filter(prime int, recv chan int, send chan int) {
	for i := range recv {
		if i%prime != 0 {
			//println("send <- ", i)
			send <- i
		}
	}
}

func counter(c chan int) {
	i := 2
	for {
		c <- i
		//println("counter c <-", i)
		i++
	}
}

func main() {
	c := make(chan int)
	go counter(c)
	primes1 := make(chan int)
	go filter(2, c, primes1)
	i := <-primes1
	println("i", i)

	primes2 := make(chan int)
	go filter(i, primes1, primes2)

	for i := 0; i < 10; i++ {
		p := <-primes2
		println(p)
	}
}
