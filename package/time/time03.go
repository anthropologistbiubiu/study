package main

import (
	"log"
	"time"
)

func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		log.Println("ticker...")
	}
}

func main() {
	TickerDemo()
}
