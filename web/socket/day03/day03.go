package main

import (
	"bufio"
	"fmt"
	"net"
)

func server() {
	listener, err := net.Listen("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		for {
			reader := bufio.NewReader(conn)
			//recv := make([]byte, 0, 128)
			recv := make([]byte, 128)
			n, err := reader.Read(recv)
			if err != nil {
				fmt.Println(err)
			}
			if n > 0 {
				fmt.Printf("server recive from %+v content:%+v\n", conn.RemoteAddr(), string(recv[:n]))
				send := "ok"
				conn.Write([]byte(send))
			}
		}
	}
}

func main() {
	server()
}
