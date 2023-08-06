package etcd

import (
	"context"
	"errors"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"log"
	"strings"
	"time"
)

var _ resolver.Builder = &EtcdResolverBuilder{}
var _ resolver.Resolver = &EtcdResolver{}

func (r *EtcdResolver) ResolveTarget(target resolver.Target) (*addressIterator, error) {
	// 解析服务名称
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

type EtcdResolverBuilder struct {
	Cli *clientv3.Client
}

func (e *EtcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &EtcdResolver{
		cli:    e.Cli,
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
	cli    *clientv3.Client
	target resolver.Target
	cc     resolver.ClientConn
}

func (r *EtcdResolver) start() {
	// 连接etcd服务器
	defer r.cli.Close()
	// 监听服务地址变化
	fmt.Println("r.target.endpoint", r.target.Endpoint())
	rch := r.cli.Watch(context.Background(), r.target.Endpoint())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			// 解析服务地址并更新到gRPC客户端
			fmt.Println("wresp", string(ev.Kv.Value), string(ev.Kv.Key))
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
	fmt.Println("update service", newAddresses)
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

	key := fmt.Sprintf("/services/%s/%s", serviceName, address)
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
