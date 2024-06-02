package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type Request struct {
	Type byte
	Data []byte
}

// Response structure
type Response struct {
	Code byte
	Data []byte
}

func UnpackResponse(data []byte) (Response, error) {
	buf := bytes.NewReader(data)
	var respCode byte
	var dataLen uint16

	if err := binary.Read(buf, binary.BigEndian, &respCode); err != nil {
		return Response{}, err
	}

	if err := binary.Read(buf, binary.BigEndian, &dataLen); err != nil {
		return Response{}, err
	}

	respData := make([]byte, dataLen)
	if _, err := buf.Read(respData); err != nil {
		return Response{}, err
	}

	return Response{Code: respCode, Data: respData}, nil
}
func PackRequest(req Request) ([]byte, error) {
	buf := new(bytes.Buffer)
	// Write request type
	if err := binary.Write(buf, binary.BigEndian, req.Type); err != nil {
		return nil, err
	}
	// Write data length
	dataLen := uint16(len(req.Data))
	if err := binary.Write(buf, binary.BigEndian, dataLen); err != nil {
		return nil, err
	}
	// Write data
	if _, err := buf.Write(req.Data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func sendRequest(conn net.Conn, reqType byte, data []byte) {
	req := Request{Type: reqType, Data: data}
	reqData, err := PackRequest(req)
	if err != nil {
		fmt.Println("Error packing request:", err)
		return
	}

	_, err = conn.Write(reqData)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	resp, err := UnpackResponse(buffer[:n])
	if err != nil {
		fmt.Println("Error unpacking response:", err)
		return
	}

	fmt.Printf("Received response: Code=%d, Data=%s\n", resp.Code, resp.Data)
}

func main() {
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	sendRequest(conn, 1, []byte("Request for business type 1"))
	sendRequest(conn, 2, []byte("Request for business type 2"))
	sendRequest(conn, 3, []byte("Request for unknown business type"))
}
