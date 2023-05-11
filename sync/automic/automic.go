package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	atomicAction()
}

// atomic原子操作
func atomicAction() {
	var b int64
	var a atomic.Value
	atomic.AddInt64(&b, 1)
	a.Store(b)
	fmt.Println(b)
	fmt.Println(a.Load())
}
