package main

import (
	"context"
	"fmt"
	//"log"
	"github.com/Jeffail/tunny"
	"golang.org/x/sync/semaphore"
	"sync"
	"time"
)

const (
	Limit  = 3
	Weight = 1
)

var s = semaphore.NewWeighted(10)
var w sync.WaitGroup

func SignalConcurrency() {

	for i := 0; i < 1000; i++ {
		err := s.Acquire(context.Background(), 1)
		if err != nil {
			fmt.Println(err)
		}
		go doSomething(i, s)
	}
}

func doSomething(u int, s *semaphore.Weighted) {
	defer func() {
		s.Release(1)
	}()
}

var workChan = make(chan struct{}, 10)

func ChannelConcurrency() {
	for i := 0; i < 1000; i++ {
		select {
		case workChan <- struct{}{}:
			go doWork(i)
		}
	}
}

var Chan = make(chan struct{}, 10)

func ChannelConcurrencyPro() {
	for i := 0; i < 1000; i++ {
		Chan <- struct{}{}
		go do(i)
	}
}
func do(i int) {

	defer func() {
		<-Chan
	}()
}

func doWork(i int) {
	defer func() {
		<-workChan
	}()
}

func poolConcurrency() {
	pool := tunny.NewFunc(10, func(i interface{}) interface{} {
		return nil
	})
	defer pool.Close()

	for i := 0; i < 1000; i++ {
		go pool.Process(i)
	}
	time.Sleep(time.Second * 1)
}
