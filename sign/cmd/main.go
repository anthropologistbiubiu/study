package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sign/proto"
)

type sign struct{}

func (s *sign) GetSign() {

}
func main() {

	listener := ":8080"
	listen, err := net.Listen("tcp", listener)
	if err != nil {
		fmt.Println("NNNNNNNNNNNNN", err)
	}
	server := grpc.NewServer()
	fmt.Println(server)
	proto.RegisterSignServiceRequestServer(server, &sign{})
	server.RegisterService(s, nil)
	if err = server.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
