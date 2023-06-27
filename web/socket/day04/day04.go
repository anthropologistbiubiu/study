package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	for {
		input := bufio.NewReader(os.Stdin)
		msg, err := input.ReadString('\n')
		msg = strings.TrimSpace(msg)
		fmt.Printf("msg :%+v\n", []byte(msg))
		if err != nil {
			fmt.Println(err)
		}
		n, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
		//var recv = []byte{}
		var recv = [128]byte{}
		n, err = conn.Read(recv[:])
		if err != nil {
			fmt.Println(err)
		}
		if n > 0 {
			fmt.Printf("client recive from remote server:%+v ,content is %+v\n", conn.RemoteAddr(), string(recv[:]))
		}
	}
}

func main() {
	client()
}
