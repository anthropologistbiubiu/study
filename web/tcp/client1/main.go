package main

import (
	"fmt"
	"net"
)

func sendData(conn net.Conn, data string) {
	fmt.Println("Sending data:", data)
	conn.Write([]byte(data))
	//time.Sleep(100 * time.Millisecond) // Introduce delay to simulate slow processing
}

func main() {
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	sendData(conn, "Packet 1 ")
	sendData(conn, "Packet 2 ")
	sendData(conn, "Packet 3 ")
}
