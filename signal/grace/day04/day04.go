package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		for {
			fmt.Println("test")
		}
	}()

	// Block until a signal is received.
	<-c
	fmt.Println("Got signal end")
}
