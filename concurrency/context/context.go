package main

import (
	"context"
	"fmt"
	"sync"
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

// 通过管道实现并发控制
func Process(ch chan int) {
	//Do some work...
	time.Sleep(time.Second)
	ch <- 1 //管道中写入一个元素表示当前协程已结束
}
func concurrency_control_channel() {
	channels := make([]chan int, 10) //创建一个10个元素的切片，元素类型为channel
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int) //切片中放入一个channel
		go Process(channels[i])      //启动协程，传一个管道用于通信
	}
	for i, ch := range channels { //遍历切片，等待子协程结束
		<-ch
		fmt.Println("Routine ", i, " quit!")
	}
}

// 信号量控制
func context_signal_control() {
	var wg sync.WaitGroup
	wg.Add(2) //设置计数器，数值即为goroutine的个数
	go func() {
		//Do some work
		time.Sleep(1 * time.Second)

		fmt.Println("Goroutine 1 finished!")
		wg.Done() //goroutine执行结束后将计数器减1
	}()

	go func() {
		//Do some work
		time.Sleep(1 * time.Second)

		fmt.Println("Goroutine 2 finished!")
		wg.Done() //goroutine执行结束后将计数器减1
	}()

	wg.Wait() //主goroutine阻塞等待计数器变为0
	fmt.Printf("All Goroutine finished!")
}

// 取消控制
func HandelRequest(ctx context.Context) {
	go WriteRedis(ctx)
	go WriteDatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}
func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}
func WriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}
func context_control() {
	ctx, cancel := context.WithCancel(context.Background())
	go HandelRequest(ctx)
	time.Sleep(1 * time.Second)
	fmt.Println("It's time to stop all sub goroutines!")
	cancel()
	//Just for test whether sub goroutines exit or not
	time.Sleep(5 * time.Second)
}

//超时控制

func context_timeout_context() {
	HttpHandler()
}

func NewContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 4*time.Second)
}

func HttpHandler() {
	ctx, cancel := NewContextWithTimeout()
	defer cancel()
	deal(ctx)
}

func deal(ctx context.Context) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n", i)
		}
	}
}

func ValueHandelRequest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value("parameter"))
			time.Sleep(2 * time.Second)
		}
	}
}
func context_value_control() {
	ctx := context.WithValue(context.Background(), "parameter", "1")
	go ValueHandelRequest(ctx)
	time.Sleep(2 * time.Second)
}
func main() {
	t := time.Tick(time.Second)
	fmt.Println(t)
	fmt.Printf("%+v", t)
	for {
		select {
		case meg := <-t:
			fmt.Println(meg)
		default:
			fmt.Println("meg")

		}
	}
}
