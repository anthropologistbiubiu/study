package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 优雅退出（退出信号）
func waitElegantExit(c chan os.Signal) {
	for i := range c {
		switch i {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			// 这里做一些清理操作或者输出相关说明，比如 断开数据库连接
			fmt.Println("receive exit signal ", i.String(), ",exit...")
			os.Exit(0)
		}
	}
}

func main() {
	//
	// 你的业务逻辑
	//
	fmt.Println("server run on: 127.0.0.1:8000")

	c := make(chan os.Signal)
	// SIGHUP: terminal closed
	// SIGINT: Ctrl+C
	// SIGTERM: program exit
	// SIGQUIT: Ctrl+/
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 阻塞，直到接受到退出信号，才停止进程
	waitElegantExit(c)
}
