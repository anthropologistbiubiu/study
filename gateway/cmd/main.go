package main

import (
	"context"
	"fmt"
	"gateway/protos"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime" // 注意v2版本
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strings"
)

type server struct {
	protos.UnimplementedOrderServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GetOrderInfo(ctx context.Context, in *protos.GetOrderReq) (*protos.GetOrderRsp, error) {
	out := &protos.GetOrderRsp{
		OrderId:   in.OrderId,
		OrderInfo: "hello order" + in.Name,
	}
	return out, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册Greeter service到server
	srv := NewServer()
	protos.RegisterOrderServer(s, srv)
	// 8080端口启动gRPC Server
	fmt.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// 创建一个连接到我们刚刚启动的 gRPC 服务器的客户端连接
	// gRPC-Gateway 就是通过它来代理请求（将HTTP请求转为RPC请求）
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		//grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		fmt.Println("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
	// 注册OrderReq
	err = protos.RegisterOrderHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	fmt.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	fmt.Println(gwServer.ListenAndServe())
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

// 在这里启动一个服务，通过http 的rest请求来完整对grpc 服务的调用
// 1.如果gw.pb.go文件可以完成，先通过这种方式来完成
// 2.再通过http.pb.go文件或者两个文件来共同完成http->grpc 的服务。
// 3.其次通过这个服务过渡到 kratos 服务，来完成启动es 服务和数据库配置，并通过 orm 来完成数据库的操作过程。
