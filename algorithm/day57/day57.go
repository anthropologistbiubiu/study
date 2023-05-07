package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

const (
	Limit  = 3 // 同時并行运行的goroutine上限
	Weight = 1 // 每个goroutine获取信号量资源的权重
)

/*
func main() {
	urls := []string{
		"http://www.example.com",
		"http://www.example.net",
		"http://www.example.net/foo",
		"http://www.example.net/bar",
		"http://www.example.net/baz",
	}
	s := semaphore.NewWeighted(Limit)
	var w sync.WaitGroup
	for _, u := range urls {
		w.Add(1)
		go func(u string) {
			s.Acquire(context.Background(), Weight)
			doSomething(u)
			s.Release(Weight)
			w.Done()
		}(u)
	}
	w.Wait()

	fmt.Println("All Done")
}
*/

var s = semaphore.NewWeighted(1000)
var w sync.WaitGroup

// 首先使用信号量机制带来的问题是什么
func main() {

	for i := 0; i < 10000; i++ {
		err := s.Acquire(context.Background(), 1)
		if err != nil {
			fmt.Println(err)
		}
		w.Add(1)
		go doSomething(i, s)
	}
	w.Wait()
}

func doSomething(u int, s *semaphore.Weighted) { // 模拟抓取任务的执行
	fmt.Println(u)
	time.Sleep(2 * time.Second)
	defer func() {
		s.Release(1)
		w.Done()
	}()
}
