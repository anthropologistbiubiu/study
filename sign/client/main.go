package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"sign/proto"
)

const (
	etcdEndpoints = "localhost:2379" // etcd 服务器的地址
)

func main() {
	// 指定服务名称
	/*
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   []string{etcdEndpoints},
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			log.Fatalf("Failed to create etcd client: %v", err)
		}
		defer cli.Close()
		resolver.Register(&etcd.EtcdResolverBuilder{Cli: cli})
		target := "etcd:///sign-service"
		conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithResolvers(&etcd.EtcdResolverBuilder{Cli: cli}), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
		if err != nil {
			fmt.Println("failed to connect: %v", err)
		}
		defer conn.Close()
	*/
	conn, err := grpc.Dial("localhost:55001", grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewSignServiceRequestClient(conn)
	req := &proto.SignRequest{
		Name:   "sunweiming",
		Email:  "1319847967@qq.com",
		Phone:  "1319847957",
		Amount: 500,
		Type:   "md5",
	}
	// 遍历服务列表并打印
	response, err := client.GetSign(context.Background(), req)
	if err != nil {
		fmt.Println("GetSign Err", err)
	}
	fmt.Printf("response %v\n", response)
}
