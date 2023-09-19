package main

/*
import (
	"context"
	"fmt"
	"gateway/protos"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime" // 注意v2版本
	"google.golang.org/grpc"
	"log"
	"net"
)

type orderServer struct {
	//protos.UnimplementedOrderServer
}

func NewOrderServer() *orderServer {
	return &orderServer{}
}

func (s *orderServer) GetOrderInfo(ctx context.Context, in *protos.GetOrderReq) (*protos.GetOrderRsp, error) {
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
		fmt.Println("Failed to listen:", err)
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
	// 注册OrderReq
	// 注册OrderReq
	mux := runtime.NewServeMux()
	protos.RegisterOrderHandlerServer(context.Background(), mux)
	// 8090端口提供gRPC-Gateway服务
	fmt.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	fmt.Println(gwServer.ListenAndServe())
}

// 在这里启动一个服务，通过http 的rest请求来完整对grpc 服务的调用
// 1.如果gw.pb.go文件可以完成，先通过这种方式来完成
// 2.再通过http.pb.go文件或者两个文件来共同完成http->grpc 的服务。
// 3.其次通过这个服务过渡到 kratos 服务，来完成启动es 服务和数据库配置，并通过 orm 来完成数据库的操作过程。

*/
