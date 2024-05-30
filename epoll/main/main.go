package main

import (
	"fmt"
	"net"
	"syscall"
)

func main() {
	// 创建监听器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()

	// 获取监听器的文件描述符
	lfd := listener.(*net.TCPListener).File().Fd()

	// 创建 epoll 实例
	epfd, err := syscall.EpollCreate1(0)
	if err != nil {
		fmt.Println("Error creating epoll instance:", err)
		return
	}
	defer syscall.Close(epfd)

	// 将监听器的文件描述符注册到 epoll 实例中
	err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, int(lfd), &syscall.EpollEvent{
		Events: syscall.EPOLLIN,
		Fd:     int32(lfd),
	})
	if err != nil {
		fmt.Println("Error adding listener fd to epoll:", err)
		return
	}

	// 事件循环
	events := make([]syscall.EpollEvent, 100)
	for {
		n, err := syscall.EpollWait(epfd, events, -1)
		if err != nil {
			fmt.Println("Error waiting for epoll events:", err)
			return
		}

		for i := 0; i < n; i++ {
			if int(events[i].Fd) == int(lfd) {
				// 处理新连接
				conn, err := listener.Accept()
				if err != nil {
					fmt.Println("Error accepting connection:", err)
					continue
				}

				// 获取新连接的文件描述符
				cfd := conn.(*net.TCPConn).File().Fd()
				// 将新连接的文件描述符注册到 epoll 实例中
				err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, int(cfd), &syscall.EpollEvent{
					Events: syscall.EPOLLIN | syscall.EPOLLET,
					Fd:     int32(cfd),
				})
				if err != nil {
					fmt.Println("Error adding connection fd to epoll:", err)
					conn.Close()
					continue
				}
			} else {
				// 处理已注册的文件描述符上的事件
				go handleConnection(int(events[i].Fd))
			}
		}
	}
}

func handleConnection(fd int) {
	defer syscall.Close(fd)

	buf := make([]byte, 1024)
	for {
		n, err := syscall.Read(fd, buf)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		if n == 0 {
			// Connection closed
			fmt.Println("Connection closed")
			return
		}

		// 简单地将接收到的数据回传给客户端
		_, err = syscall.Write(fd, buf[:n])
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}
