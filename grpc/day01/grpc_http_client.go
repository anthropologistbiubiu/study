package main

import (
	"fmt"
	"log"
	"net/rpc"
)

/*
type Args struct {
	X, Y int
}
*/

func main() {
	// 建立HTTP连接
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9093")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 同步调用
	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}
	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.X, args.Y, reply)

	// 异步调用
	var reply2 int
	divCall := client.Go("ServiceA.Add reply2", args, &reply2, nil)
	replyCall := <-divCall.Done // 接收调用结果
	fmt.Println("replayCall.Error", replyCall.Error)
	fmt.Println("replay2", reply2)
}
