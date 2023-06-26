package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func client() {

	conn, err := net.Dial("tcp", ":9091")
	if err != nil {
		fmt.Println(err)
	}
	for {
		input := bufio.NewReader(os.Stdin)
		var buff = []byte{}
		_, err := input.Read(buff)
		if err != nil {
			fmt.Println(err)
		}
		conn.Write(buff)
		recv := make([]byte, 0, 1024)
		_, err = conn.Read(recv)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("recive from remote server:%+v ,content is %+v\n", conn.RemoteAddr(), string(recv))
	}

}

func main() {
	client()
}