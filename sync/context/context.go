package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	contextTest()
}

func contextTest() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ctx context.Context) {
		defer func() {
			wg.Done()
		}()
		for {
			select {
			case <-ctx.Done():
				time.Sleep(1 * time.Second)
				fmt.Println("停止了")
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("还在跑")
			}
		}
	}(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}
