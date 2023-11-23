package main

import (
	"log"
	"os"
	"sync"
)

type AppConfig struct {
	version  string
	serverIp string
	port     int
}

var (
	config = new(AppConfig)
	once   = sync.Once{}
)

func ReadConfig() {
	once.Do(func() {
		conig = &AppConfig{
			version:  os.Getenv("version"),
			serverIp: "localhost",
			port:     8080,
		}
	})
}
func main() {
	for i := 0; i < 10; i++ {
		go ReadConfig()
	}
	log.Println(config)
}
