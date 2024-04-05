package main

import (
	"fmt"
	"time"
)

func main() {
	// 1 在这里需要你写算法
	// 2 要求每秒钟调用一次proc函数
	// 3 要求程序不能退出
	t := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-t.C:
			go func() {
				defer func() {
					fmt.Println(recover())
				}()
				proc()
			}()
		}
	}

}

func proc() {
	panic("ok")
}
