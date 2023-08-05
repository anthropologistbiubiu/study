package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"sign/etcd"
	"sign/proto"
)

// 客户端给出 负载均衡的调用过程
func main() {
	resolver.Register(&etcd.EtcdResolverBuilder{})
	// 连接gRPC服务器
	conn, err := grpc.Dial("etcd:///sign-service", grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		fmt.Println("Failed to connect: %v", err)
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
