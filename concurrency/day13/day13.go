package main

func main() {

	var chan1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		chan1 <- i
	}
	/*
		for i := 1; i <= 10; i++ {
			j := <-chan1
			println(j)
		}
	*/

	for j := range chan1 {
		println(j)
	}
	k := <-chan1
	println("k", k)
}
