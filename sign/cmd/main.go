package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sign/proto"
	"time"
)

type SignServer struct {
}

/*
func (s *sign) mustEmbedUnimplementedSignServiceRequestServer() {

}
*/
func (s *SignServer) mustEmbedUnimplementedSignServiceRequestServer() {}

func (s *SignServer) GetSign(ctx context.Context, req *proto.SignRequest) (*proto.SignReponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	time.Sleep(5 * time.Second) //模拟数据库操作
	hash := sha256.New()
	hash.Write(data)
	hashValue := hash.Sum(nil)
	response := &proto.SignReponse{
		Sign: hex.EncodeToString(hashValue),
		Code: 200,
	}
	return response, nil
}
func main() {

	listener := ":8080"
	listen, err := net.Listen("tcp", listener)
	if err != nil {
		fmt.Println("", err)
	}
	server := grpc.NewServer()
	reflection.Register(server)
	proto.RegisterSignServiceRequestServer(server, &SignServer{})
	fmt.Println("start success!")
	if err = server.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
