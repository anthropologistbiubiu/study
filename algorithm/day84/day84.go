package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("主协程结束")
				return
			default:
				go func() {
					//fmt.Println("child start")
					time.Sleep(time.Second * 4)
					fmt.Println("child end")
				}()
			}
		}
	}()
	time.Sleep(3 * time.Second)
}
