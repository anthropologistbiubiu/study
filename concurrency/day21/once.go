package main

import (
	"log"
	"os"
	"sync"
	"time"
)

type AppConfig struct {
	version  string
	serverIp string
	port     int
}

func init() {

	_ = os.Setenv("version", "1.0")
}

var (
	//config = new(AppConfig)
	once = sync.Once{}
)

func ReadConfig() {
	once.Do(func() {
		conig := &AppConfig{
			version:  os.Getenv("version"),
			serverIp: "localhost",
			port:     8080,
		}
		log.Println(conig)
	})
}
func main() {
	for i := 0; i < 10; i++ {
		go ReadConfig()
	}
	time.Sleep(time.Second * 10)
}
