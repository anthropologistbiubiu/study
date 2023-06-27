package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func server() {
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(conn)
		go func(net.Conn) {
			fmt.Println(conn)
			for {
				input := bufio.NewReader(conn)
				recv := make([]byte, 0, 1024)
				n, err := input.Read(recv)
				fmt.Println("server n ", n, string(recv))
				time.Sleep(5 * time.Second)
				if err != nil {
					fmt.Println(err)
				}
				if n > 0 {
					fmt.Printf("server recive from %+v content:%+v\n", conn.RemoteAddr(), string(recv[:n]))
					send := "ok"
					conn.Write([]byte(send))
				}
			}
		}(conn)
	}
}

func main() {
	server()
}
