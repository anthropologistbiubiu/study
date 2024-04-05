package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//设置超时控制WithDeadline，超时时间2
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	case <-ctx.Done():
		//2到了到了，执行该代码
		fmt.Println(ctx.Err())
	}

}
