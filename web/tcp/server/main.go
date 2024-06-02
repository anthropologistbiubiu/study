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

// PackRequest packs a request into a byte slice

// UnpackRequest unpacks a byte slice into a request
func UnpackRequest(data []byte) (Request, error) {
	buf := bytes.NewReader(data)
	var reqType byte
	var dataLen uint16

	if err := binary.Read(buf, binary.BigEndian, &reqType); err != nil {
		return Request{}, err
	}

	if err := binary.Read(buf, binary.BigEndian, &dataLen); err != nil {
		return Request{}, err
	}

	reqData := make([]byte, dataLen)
	if _, err := buf.Read(reqData); err != nil {
		return Request{}, err
	}

	return Request{Type: reqType, Data: reqData}, nil
}

// PackResponse packs a response into a byte slice
func PackResponse(resp Response) ([]byte, error) {
	buf := new(bytes.Buffer)
	// Write response code
	if err := binary.Write(buf, binary.BigEndian, resp.Code); err != nil {
		return nil, err
	}
	// Write data length
	dataLen := uint16(len(resp.Data))
	if err := binary.Write(buf, binary.BigEndian, dataLen); err != nil {
		return nil, err
	}
	// Write data
	if _, err := buf.Write(resp.Data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnpackResponse unpacks a byte slice into a response

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		req, err := UnpackRequest(buffer[:n])
		if err != nil {
			fmt.Println("Error unpacking request:", err)
			return
		}

		resp := routeRequest(req)
		respData, err := PackResponse(resp)
		if err != nil {
			fmt.Println("Error packing response:", err)
			return
		}

		_, err = conn.Write(respData)
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}

func routeRequest(req Request) Response {
	switch req.Type {
	case 1:
		return handleBusinessType1(req)
	case 2:
		return handleBusinessType2(req)
	default:
		return Response{Code: 1, Data: []byte("Unknown request type")}
	}
}

func handleBusinessType1(req Request) Response {
	// Example business logic for type 1
	fmt.Println("Handling business type 1")
	return Response{Code: 0, Data: []byte("Response from business type 1")}
}

func handleBusinessType2(req Request) Response {
	// Example business logic for type 2
	fmt.Println("Handling business type 2")
	return Response{Code: 0, Data: []byte("Response from business type 2")}
}

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on :12345")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
