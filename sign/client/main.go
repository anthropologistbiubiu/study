package main

import (
	"context"
	"errors"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"sign/proto"
	"strings"
	"time"
)

const (
	etcdEndpoints = "localhost:2379" // etcd 服务器的地址
)

// etcdResolver 实现了 gRPC 的 Resolver 接口
type etcdResolver struct {
	cli *clientv3.Client
}

func (r *etcdResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	return &etcdResolver{
		cli: r.cli,
	}, nil
}

// 客户端给出 负载均衡的调用过程
func (r *etcdResolver) Scheme() string {
	return "etcd"
}

// ResolveNow 当前不会主动触发名字解析
func (r *etcdResolver) ResolveNow(resolver.ResolveNowOptions) {}

// Close 关闭解析器
func (r *etcdResolver) Close() {}

// ResolveTarget 实现了名字解析的逻辑
func (r *etcdResolver) ResolveTarget(target resolver.Target) (*addressIterator, error) {
	// 解析服务名称
	fmt.Println("(((((((((((")
	resp, err := r.cli.Get(context.Background(), target.Endpoint())
	if err != nil {
		return nil, fmt.Errorf("failed to get service addresses from etcd: %v", err)
	}

	var addresses []resolver.Address
	for _, kv := range resp.Kvs {
		// 解析服务地址列表
		addrs := strings.Split(string(kv.Value), ",")
		for _, addr := range addrs {
			addresses = append(addresses, resolver.Address{Addr: addr})
		}
	}

	// 将解析后的地址列表更新到 gRPC 客户端连接中
	// 返回一个实现了 resolver.AddressIterator 接口的对象
	return &addressIterator{addresses: addresses}, nil
}

type addressIterator struct {
	addresses []resolver.Address
	next      int
}

func (it *addressIterator) Next() (resolver.Address, error) {
	if it.next >= len(it.addresses) {
		return resolver.Address{}, errors.New("err")
	}
	addr := it.addresses[it.next]
	it.next++
	return addr, nil
}
func main() {
	// 指定服务名称
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdEndpoints},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to create etcd client: %v", err)
	}
	defer cli.Close()
	// 注册 etcd 解析器
	r := &etcdResolver{
		cli: cli,
	}
	resolver.Register(r)
	// 创建 gRPC 连接
	target := "etcd:///sign-service"
	//resolver.Register(&etcd.EtcdResolverBuilder{})
	conn, err := grpc.Dial(target, grpc.WithResolvers(r),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	fmt.Printf("conn %+v \n", conn.Target())
	if err != nil {
		fmt.Println("failed to connect: %v", err)
	}
	defer conn.Close()
	fmt.Println("NNNNNNNNNNNNNNNNNNNN")
	client := proto.NewSignServiceRequestClient(conn)
	req := &proto.SignRequest{
		Name:   "sunweiming",
		Email:  "1319847967@qq.com",
		Phone:  "1319847957",
		Amount: 500,
	}
	fmt.Println("Client >>>>>>>", client)
	for i := 0; i < 10; i++ {
		fmt.Println("CCCCCCCCCCCC")
		response, err := client.GetSign(context.Background(), req)
		if err != nil {
			fmt.Println("GetSign Err", err)
		}
		fmt.Printf("%v\n", response)
	}
}
