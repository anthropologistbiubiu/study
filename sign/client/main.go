package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"sign/proto"
)

func main() {
	serverAddr := "localhost:8080"
	// 连接gRPC服务器
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to connect: %v", err)
	}
	defer conn.Close()
	// 创建gRPC客户端
	client := proto.NewSignServiceRequestClient(conn)
	req := &proto.SignRequest{
		Name:   "sunweiming",
		Email:  "1319847967@qq.com",
		Phone:  "1319847957",
		Amount: 500,
	}
	response, err := client.GetSign(context.Background(), req)
	if err != nil {
		fmt.Println("GetSign Err", err)
	}
	fmt.Printf("%v\n", response)

}
