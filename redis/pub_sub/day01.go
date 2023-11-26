package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func write(names ...string) {

	for {
		time.Sleep(1 * time.Second)
		log.Println("write", time.Now().Second())
	}
}

func read() {
	for {
		time.Sleep(1 * time.Second)
		log.Println("read", time.Now().Second())
	}
}

func main() {

	sigchan := make(chan os.Signal, 1)
	go write()
	go read()
	signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGINT)
	<-sigchan
	time.Sleep(5 * time.Second)

}
