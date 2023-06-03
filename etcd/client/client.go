package main

import (
	"context"
	"etcd/proto"
	"etcd/register"
	"flag"
	"google.golang.org/grpc"
	"log"
	"time"
)

//var serAddr = flag.String("addr", "localhost:8000", "the address to connect to")

const registerDialPrefix = "register://localhost:2379/"

func main() {
	// 解析命令行参数
	flag.Parse()

	service, err := register.NewLocalDefNamingService("my1")
	if err != nil {
		log.Println("Create naming service error: %v", err)
	}
	resolver, err := service.NewEtcdResolver()
	if err != nil {
		log.Fatalf("Create register resolver error: %v", err)
	}
	// 连接服务端
	conn, err := grpc.Dial(registerDialPrefix+service.GetPathServerName("s1"), grpc.WithInsecure(), grpc.WithResolvers(resolver),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		log.Fatalf("Conn server error: %v", err)
	}
	log.Printf("Conn success: %v", conn.GetState())
	// 执行完方法自动关闭资源
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Close conn error: %v", err)
			return
		}
		log.Println("Close conn success")
	}()

	// 创建客户端
	c := proto.NewGreeterClient(conn)

	log.Println("5秒中之后调用SayHello方法")
	time.Sleep(time.Second * 5)
	num := 10
	for i := 0; i < num; i++ {
		// 创建2秒超时ctx
		ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
		// 发起RPC请求
		log.Println("开始调用SayHello方法")
		res, err := c.SayHello(ctx, &proto.HelloRequest{Name: "一号"})
		if err != nil {
			log.Fatalf("请求失败: %v", err)
		}
		log.Printf("请求结果: %s", res.GetMessage())
	}

	// 睡眠一会再结束
	log.Println("3秒后结束，客户端自动断开连接")
	time.Sleep(time.Second * 3)

}
