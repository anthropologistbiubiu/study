package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

// Request structure
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
func PackRequest(req Request) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, req.Type); err != nil {
		return nil, err
	}
	dataLen := uint16(len(req.Data))
	if err := binary.Write(buf, binary.BigEndian, dataLen); err != nil {
		return nil, err
	}
	if _, err := buf.Write(req.Data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

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
	if err := binary.Write(buf, binary.BigEndian, resp.Code); err != nil {
		return nil, err
	}
	dataLen := uint16(len(resp.Data))
	if err := binary.Write(buf, binary.BigEndian, dataLen); err != nil {
		return nil, err
	}
	if _, err := buf.Write(resp.Data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnpackResponse unpacks a byte slice into a response
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

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	var dataBuffer bytes.Buffer
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		dataBuffer.Write(buffer[:n])

		for {
			if dataBuffer.Len() < 3 {
				// Not enough data to read type and length
				break
			}
			data := dataBuffer.Bytes()
			dataLen := binary.BigEndian.Uint16(data[1:3])
			packetLen := 3 + int(dataLen)
			if dataBuffer.Len() < packetLen {
				// Not enough data to read the whole packet
				break
			}
			req, err := UnpackRequest(data[:packetLen])
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
			dataBuffer.Next(packetLen)
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
		fmt.Println("Handling business type 3")
		return Response{Code: 1, Data: []byte("Unknown request type")}
	}
}

func handleBusinessType1(req Request) Response {
	fmt.Println("Handling business type 1")
	return Response{Code: 0, Data: []byte("Response from business type 1")}
}

func handleBusinessType2(req Request) Response {
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
