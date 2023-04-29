package main

import (
	"context"
	"day03/pb"
	"flag"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpc 客户端
// 调用server端的 SayHello 方法

var name = flag.String("name", "sunweiming", "")

// valid 校验认证信息.
/*
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// 执行token认证的逻辑
	// 这里是为了演示方便简单判断token是否与"some-secret-token"相等
	return token == "some-secret-token"
}

// unaryInterceptor 客户端一元拦截器
func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var credsConfigured bool
	for _, o := range opts {
		_, ok := o.(grpc.PerRPCCredsCallOption)
		if ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	end := time.Now()
	fmt.Printf("RPC: %s, start time: %s, end time: %s, err: %v\n", method, start.Format("Basic"), end.Format(time.RFC3339), err)
	return err
}
*/
// unaryCallWithMetadata 普通RPC调用客户端metadata操作
func unaryCallWithMetadata(c pb.GreeterClient, name string) {
	// 创建metadata
	md := metadata.Pairs(
		"token", "app-test-sunweiming",
		"request_id", "1234567",
	)
	// 基于metadata创建context.
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// RPC调用
	var header, trailer metadata.MD
	r, err := c.UnarySayHello(
		ctx,
		&pb.HelloRequest{Name: name},
		grpc.Header(&header),   // 接收服务端发来的header
		grpc.Trailer(&trailer), // 接收服务端发来的trailer
	)
	if err != nil {
		log.Printf("failed to call SayHello: %v", err)
		return
	}
	// 从header中取location
	if t, ok := header["location"]; ok {
		fmt.Printf("location from header: %+v\n", t)
		/*
			for i, e := range t {
				fmt.Printf("%d %s\n", i, e)
			}
		*/
	} else {
		log.Printf("location expected but doesn't exist in header")
		return
	}
	if t, ok := header["name"]; ok {
		fmt.Printf("name from header: %+v\n", t)
		/*
			for i, e := range t {
				fmt.Printf("%d %s\n", i, e)
			}
		*/
	} else {
		log.Printf("name expected but doesn't exist in header")
		return
	}
	// 获取响应结果
	fmt.Printf("get response: %s\n", r.Reply)
	// 从trailer中取timestamp
	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer: tailer %+v\n", t)
		/*
			for i, e := range t {
				fmt.Printf("%d %s\n", i, e)
			}
		*/
	} else {
		log.Printf("timestamp expected but doesn't exist in trailer")
	}
	if t, ok := trailer["name"]; ok {
		fmt.Printf("name from trailer: tailer %+v\n", t)
		/*
			for i, e := range t {
				fmt.Printf("%d %s\n", i, e)
			}
		*/
	} else {
		log.Printf("name expected but doesn't exist in trailer")
	}
}

// bidirectionalWithMetadata 流式RPC调用客户端metadata操作
func bidirectionalWithMetadata(c pb.GreeterClient, name string) {
	// 创建metadata和context.
	md := metadata.Pairs("token", "app-test-sunweiming")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 使用带有metadata的context执行RPC调用.
	stream, err := c.BidirectionalStreamingSayHello(ctx)
	if err != nil {
		log.Fatalf("failed to call BidiHello: %v\n", err)
	}

	go func() {
		// 当header到达时读取header.
		header, err := stream.Header()

		if err != nil {
			log.Fatalf("failed to get header from stream: %v", err)
		}
		// 从返回响应的header中读取数据.
		if l, ok := header["name"]; ok {
			fmt.Printf("name from header: %+v \n", l)
			/*
				for i, e := range l {
					fmt.Printf("%d %s\n", i, e)
				}
			*/
		} else {
			log.Println("name expected but doesn't exist in header")
			return
		}

		// 发送所有的请求数据到server.
		for i := 0; i < 5; i++ {
			if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
				log.Fatalf("failed to send streaming: %v\n", err)
			}
		}
		stream.CloseSend()
	}()

	// 读取所有的响应.
	var rpcStatus error
	for {
		r, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		fmt.Printf("%s\n", r.Reply)
	}
	if rpcStatus != io.EOF {
		log.Printf("failed to finish server streaming: %v", rpcStatus)
		return
	}

	// 当RPC结束时读取trailer
	trailer := stream.Trailer()
	// 从返回响应的trailer中读取metadata.
	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer: %+v\n", t)
		/*
			for i, e := range t {
				fmt.Printf(" %d. %s\n", i, e)
			}
		*/
	} else {
		log.Printf("timestamp expected but doesn't exist in trailer")
	}
}
func main() {
	flag.Parse() // 解析命令行参数

	// 连接server
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}
	defer conn.Close()
	// 创建客户端
	c := pb.NewGreeterClient(conn) // 使用生成的Go代码
	// 调用RPC方法
	/*
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
			resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
			if err != nil {
				s := status.Convert(err)        // 将err转为status
				for _, d := range s.Details() { // 获取details
					switch info := d.(type) {
					case *errdetails.QuotaFailure:
						fmt.Printf("Quota failure: %+v\n", info)
					default:
						fmt.Printf("Unexpected type: %s\n", info)
					}
				}
				fmt.Printf("c.SayHello failed, err:%v\n", err)
				return
			}
			// 拿到了RPC响应
			log.Printf("resp:%v\n", resp.GetReply())
	*/
	bidirectionalWithMetadata(c, "sunweiming")
}
