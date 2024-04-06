package main

import "fmt"

func main() {

	chans := make(chan int, 10)
	chans <- 1
	chans <- 2
	chans <- 3
	chans <- 4
	close(chans)
	for range chans {
	}
	for v := range chans {
		fmt.Println(v)
	}

}
