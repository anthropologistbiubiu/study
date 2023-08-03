package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"sign/proto"
	"sync"
	"syscall"
	"time"
)

type SignServer struct {
}

/*
func (s *sign) mustEmbedUnimplementedSignServiceRequestServer() {
}
*/
// 现在已经通过grpc 改造了签名服务器，现在剩下的就是改造每个业务层，实现这个业务的完整性，在业务层中添加 orm 的过程。
// 添加服务的注册与发现
// 添加 log 层的日志归档和记录
// kafa 实现请求的限流的熔断

func (s *SignServer) mustEmbedUnimplementedSignServiceRequestServer() {}

func (s *SignServer) GetSign(ctx context.Context, req *proto.SignRequest) (*proto.SignReponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	var flag bool
	timer := time.After(10 * time.Second)
	for !flag {
		select {
		case <-timer:
			fmt.Println("10s is coming")
			flag = true
		default:
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Second())
		}
	}
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
	var wg sync.WaitGroup
	wg.Add(1)
	// 处理优雅退出信号
	go func() {
		defer wg.Done()
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		fmt.Println("\nReceived shutdown signal. Gracefully shutting down...")
		// 关闭gRPC服务器
		server.GracefulStop()
	}()

	fmt.Println("start success!")
	// 启动gRPC服务器
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// 等待所有活动的请求处理完成
	wg.Wait()
}
