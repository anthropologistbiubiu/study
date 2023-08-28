package main

import (
	"consul/middleware"
	"consul/pb"
	"consul/service"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func grpc_main() {

	// 写一个grpc 服务注册到
	consul, err := middleware.NewConsul("127.0.0.1:8500")
	if err != nil {
		fmt.Println(err)
	}
	consul.RegisterService("jobservice", "127.0.0.1", 8082)
	address := ":8082"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Print("wwww", err)
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(middleware.LimiterInterceptor))
	pb.RegisterJobServicevRequestServer(server, &service.JobServiceServer{})
	fmt.Println("server2 start")
	if err := server.Serve(listener); err != nil {
		fmt.Println("NNN", err)
	}
}

func main() {
	grpc_main()
}
