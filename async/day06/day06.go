package main

import (
	"fmt"
	"sync"
)

// 本文主要记录了并发例题
/*
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var x, y int
	go func() {
		x = 1
		y = 1
		wg.Done()
	}()
	go func() {
		r1 := y
		r2 := x

		fmt.Println(r1, r2) // ❶ r1 = 1, r2 = 0 可能吗?

		wg.Done()
	}()
	wg.Wait()
}


*/
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var x, y int
	go func() {
		x = 1
		r1 := y
		fmt.Println(r1) // ❶
		wg.Done()
	}()
	go func() {
		y = 1
		r2 := x
		fmt.Println(r2) // ❷
		wg.Done()
	}()
	wg.Wait()
}
