package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sign/proto"
)

type SignServer struct {
}

/*
func (s *sign) mustEmbedUnimplementedSignServiceRequestServer() {

}
*/
func (s *SignServer) mustEmbedUnimplementedSignServiceRequestServer() {}

func (s *SignServer) GetSign(ctx context.Context, req *proto.SignRequest) (*proto.SignReponse, error) {
	fmt.Println(ctx)
	data, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("WWWWWWWWWWWWWWWWWWW", string(data))
	response := &proto.SignReponse{
		Sign: "sign123456",
		Code: 200,
	}
	return response, nil
}
func main() {

	listener := ":8080"
	listen, err := net.Listen("tcp", listener)
	if err != nil {
		fmt.Println("NNNNNNNNNNNNN", err)
	}
	server := grpc.NewServer()
	fmt.Println(server)
	proto.RegisterSignServiceRequestServer(server, &SignServer{})
	if err = server.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
