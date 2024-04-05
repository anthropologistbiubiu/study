package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	pctx := context.Background()
	cctx, cancelfunc := context.WithCancel(pctx)
	sctx, _ := context.WithTimeout(cctx, time.Duration(2*time.Second))
	go Contexttest(sctx)
	time.Sleep(time.Second * 1)
	cancelfunc()
	//cancel()
	time.Sleep(5 * time.Second)
}

func Contexttest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("nihao")
			return
		default:
			fmt.Println("wo bu hao")
			time.Sleep(1 * time.Second)
		}
	}
}
