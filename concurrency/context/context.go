package main

import (
	"fmt"
	"time"
)

func Goroutine_simple_many_who() {
	str := []string{"1111", "2222", "3333", "4444"}
	for _, pr := range str {
		//fmt.Println(pr)
		go func() {
			fmt.Println(pr)
		}()
	}
}

func Goroutine_simple_many_order() {
	str := []string{"1111", "2222", "3333", "4444"}
	for i, pr := range str {
		go func(i int, name string) {
			fmt.Println(i, pr)
		}(i, pr)
	}
}
func main() {
	//Goroutine_simple_many_order()
	Goroutine_simple_many_who()
	time.Sleep(10 * time.Second)
}
