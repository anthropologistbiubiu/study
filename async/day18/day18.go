package main

func counter() chan int {
	out := make(chan int)
	go func() {
		for i := 2; ; i++ {
			out <- i
		}
	}()
	return out
}

func filter(prime int, in chan int) chan int {

	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := counter()

	for i := 1; i <= 10; i++ {
		prime := <-ch
		println(prime)
		primes := filter(prime, ch)
		ch = primes
	}
}
