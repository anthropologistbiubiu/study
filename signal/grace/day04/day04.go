package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	time.Sleep(5 * time.Second)
	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}
