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

// 这里给出解析 etcd 的过程
var (
	etcdEndpoints = []string{"localhost:2379"}
	etcdKeyPrefix = "sign-service/"
)

func init() {
	// 注册etcd的服务发现解析器
	resolver.Register(&EtcdResolverBuilder{})
}

type EtcdResolverBuilder struct{}

func (*EtcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &etcdResolver{
		target: target,
		cc:     cc,
	}
	go r.watch()
	return r, nil
}

func (*EtcdResolverBuilder) Scheme() string {
	return "etcd"
}

func RegisterService(serviceAddr string) error {
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: etcdEndpoints,
	})
	if err != nil {
		return err
	}
	defer etcdCli.Close()
	key := etcdKeyPrefix + serviceAddr
	fmt.Println("NNNNNNNNNNNNNNN", key)
	_, err = etcdCli.Put(context.Background(), key, serviceAddr)
	return err
}

func DeregisterService(serviceAddr string) error {
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

	key := etcdKeyPrefix + r.target.Endpoint()
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
