package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"log"
	"strings"
	"time"
)

func init() {
	resolver.Register(&EtcdResolverBuilder{})
}

type EtcdResolverBuilder struct{}

func (*EtcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &EtcdResolver{
		target: target,
		cc:     cc,
	}
	r.start()
	return r, nil
}

func (*EtcdResolverBuilder) Scheme() string {
	return "etcd"
}

type EtcdResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
}

func (r *EtcdResolver) start() {
	// 连接etcd服务器
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd服务地址
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("failed to connect to etcd: %v", err)
	}
	defer cli.Close()

	// 监听服务地址变化
	rch := cli.Watch(context.Background(), r.target.Endpoint())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			// 解析服务地址并更新到gRPC客户端
			r.updateAddresses(string(ev.Kv.Value))
		}
	}
}

func (r *EtcdResolver) updateAddresses(addresses string) {
	var newAddresses []resolver.Address
	// 解析以逗号分隔的服务地址列表
	for _, addr := range strings.Split(addresses, ",") {
		newAddresses = append(newAddresses, resolver.Address{Addr: addr})
	}
	r.cc.UpdateState(resolver.State{Addresses: newAddresses})
}

func (*EtcdResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (*EtcdResolver) Close()                                  {}

// 向 etcd 注册服务地址
func RegisterServiceWithEtcd(serviceName, address string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd服务地址
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("failed to connect to etcd: %v", err)
	}
	defer cli.Close()

	key := fmt.Sprintf("%s", serviceName)
	fmt.Println("KEY", key)
	_, err = cli.Put(context.Background(), key, address)
	if err != nil {
		log.Fatalf("failed to register service: %v", err)
	}
	fmt.Printf("Service registered: %s\n", address)
}

// 从 etcd 注销服务地址
func UnregisterServiceWithEtcd(serviceName, address string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd服务地址
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("failed to connect to etcd: %v", err)
	}
	defer cli.Close()

	key := fmt.Sprintf("/services/%s/%s", serviceName, address)
	_, err = cli.Delete(context.Background(), key)
	if err != nil {
		log.Fatalf("failed to unregister service: %v", err)
	}
	fmt.Printf("Service unregistered: %s\n", address)
}
