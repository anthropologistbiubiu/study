package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	"os"
	"os/signal"
	"sign/proto"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	etcdEndpoints = []string{"localhost:2379"}
	etcdKeyPrefix = "sign-service/"
)

func etcdregist() {
	serviceAddr := "localhost:50051" // 替换为实际的服务器地址
	if err := registerService(serviceAddr); err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}
}

type SignServer struct {
}

/*
func (s *sign) mustEmbedUnimplementedSignServiceRequestServer() {
}
*/
// 现在已经通过grpc 改造了签名服务器，现在剩下的就是改造每个业务层，实现这个业务的完整性，在业务层中添加 orm 的过程。
// 添加服务的注册与发现  添加负载均衡
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
	// 在服务器退出时，注销服务
	if err := deregisterService(serviceAddr); err != nil {
		log.Fatalf("Failed to deregister service: %v", err)
	}
}

func registerService(serviceAddr string) error {
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: etcdEndpoints,
	})
	if err != nil {
		return err
	}
	defer etcdCli.Close()

	key := etcdKeyPrefix + serviceAddr
	_, err = etcdCli.Put(context.Background(), key, serviceAddr)
	return err
}
func deregisterService(serviceAddr string) error {
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: etcdEndpoints,
	})
	if err != nil {
		return err
	}
	defer etcdCli.Close()

	key := etcdKeyPrefix + serviceAddr
	_, err = etcdCli.Delete(context.Background(), key)
	return err
}
func init() {
	// 注册etcd的服务发现解析器
	resolver.Register(&etcdResolverBuilder{})
}

type etcdResolverBuilder struct{}

func (*etcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &etcdResolver{
		target: target,
		cc:     cc,
	}
	go r.watch()
	return r, nil
}
func (*etcdResolverBuilder) Scheme() string {
	return "etcd"
}

type etcdResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
}

func (r *etcdResolver) ResolveNow(options resolver.ResolveNowOptions) {}

func (r *etcdResolver) Close() {}
func (r *etcdResolver) watch() {
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: etcdEndpoints,
	})
	if err != nil {
		log.Fatalf("Failed to create etcd client: %v", err)
	}
	defer etcdCli.Close()

	key := etcdKeyPrefix + r.target.Endpoint
	for {
		resp, err := etcdCli.Get(context.Background(), key)
		if err != nil {
			log.Printf("Failed to get etcd key %s: %v", key, err)
			time.Sleep(time.Second)
			continue
		}

		var addresses []resolver.Address
		for _, kv := range resp.Kvs {
			addresses = append(addresses, resolver.Address{Addr: strings.TrimPrefix(string(kv.Value), "http://")})
		}

		r.cc.NewAddress(addresses)
		// 等待下一次etcd变更通知
		r.watchKey(key, resp.Header.Revision+1)
	}
}

func (r *etcdResolver) watchKey(key string, revision int64) {
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: etcdEndpoints,
	})
	if err != nil {
		log.Fatalf("Failed to create etcd client: %v", err)
	}
	defer etcdCli.Close()

	rch := etcdCli.Watch(context.Background(), key, clientv3.WithRev(revision))
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				r.cc.NewAddress([]resolver.Address{{Addr: strings.TrimPrefix(string(ev.Kv.Value), "http://")}})
			case clientv3.EventTypeDelete:
				r.cc.NewAddress([]resolver.Address{})
			}
		}
	}
}
func (r *etcdResolver) watchKey(key string, revision int64) {
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: etcdEndpoints,
	})
	if err != nil {
		log.Fatalf("Failed to create etcd client: %v", err)
	}
	defer etcdCli.Close()

	rch := etcdCli.Watch(context.Background(), key, clientv3.WithRev(revision))
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				r.cc.NewAddress([]resolver.Address{{Addr: strings.TrimPrefix(string(ev.Kv.Value), "http://")}})
			case clientv3.EventTypeDelete:
				r.cc.NewAddress([]resolver.Address{})
			}
		}
	}
}
