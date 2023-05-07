package main

// 用一个栈空间和一个读写锁实现并发控制

import (
	"fmt"

	"sync"

	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i = i + 1 {

		wg.Add(1)

		go func(n int) {

			// defer wg.Done()

			defer wg.Add(-1)

			EchoNumber(n)

		}(i)

	}

	wg.Wait()

}

func EchoNumber(i int) {

	time.Sleep(3e9)

	fmt.Println(i)

}
