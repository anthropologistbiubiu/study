package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Println(err)
	}
	for {
		input := bufio.NewReader(os.Stdin)
		msg, err := input.ReadString('\n')
		fmt.Printf("msg :%+v", msg)
		if err != nil {
			fmt.Println(err)
		}
		conn.Write([]byte(msg))
		var recv = []byte{}
		n, err := conn.Read(recv)
		if err != nil {
			fmt.Println(err)
		}
		if n > 0 {
			fmt.Printf("client recive from remote server:%+v ,content is %+v\n", conn.RemoteAddr(), string(recv))
		}
	}
}

func main() {
	client()
}
