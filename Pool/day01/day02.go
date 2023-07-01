package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var array = make([]int, 0, 5)
	wg.Add(2)
	go func(group *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			array = append(array, i)
		}
	}(&wg)

	go func(group *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println(array[i])
		}
	}(&wg)
	wg.Wait()
}
