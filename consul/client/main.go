package main

import (
	"consul/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"
	"strconv"
	"sync"
	"time"
)

func http_main() {
	cc, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Printf("api.NewClient failed, err:%v\n", err)
		return
	}
	// 返回的是一个 map[string]*api.AgentService
	// 其中key是服务ID，值是注册的服务信息
	services, _, err := cc.Catalog().Services(nil)
	if err != nil {
		fmt.Println(err)
	}
	for serviceName := range services {
		fmt.Println("Service:", serviceName)
	}
	serviceMap, err := cc.Agent().ServicesWithFilter("Service==`hello`")
	if err != nil {
		fmt.Printf("query service from consul failed, err:%v\n", err)
		return
	}
	// 选一个服务机（这里选最后一个）
	var addr string
	for k, v := range serviceMap {
		fmt.Printf("%s:%#v\n", k, v)
		addr = v.Address + ":" + strconv.Itoa(v.Port)
	}
	fmt.Println("result", addr)
}

const consulScheme = "consul"

type consulResolverBuilder struct{}

func (b *consulResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &consulResolver{
		target: target,
		cc:     cc,
	}
	r.start()
	return r, nil
}

func (*consulResolverBuilder) Scheme() string {
	return "consul"
}

type consulResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
	mu     sync.Mutex
}

func (r *consulResolver) start() {
	// 解析服务名
	serviceName := r.target.Endpoint()
	fmt.Printf("开始解析服务 ：%+v\n", serviceName)
	// 初始化 Consul 客户端
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Printf("创建 Consul 客户端失败：%v\n", err)
		return
	}
	// 查询服务实例
	instances, _, err := consulClient.Catalog().Service(serviceName, "", nil)
	if err != nil {
		fmt.Printf("查询服务实例失败：%v\n", err)
		return
	}
	// 构建解析器状态
	var addresses []resolver.Address
	for _, instance := range instances {
		addresses = append(addresses, resolver.Address{
			Addr: fmt.Sprintf("%s:%d", instance.ServiceAddress, instance.ServicePort),
		})
	}
	// 更新解析器状态
	fmt.Println("wwwww 查看解析出来的地址", addresses)
	r.cc.UpdateState(resolver.State{
		Addresses: addresses,
	})

}
func (r *consulResolver) ResolveNow(o resolver.ResolveNowOptions) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// 在此处可以实现主动刷新服务实例列表
	r.cc.UpdateState(resolver.State{Addresses: []resolver.Address{{Addr: "localhost:50051"}}})
}

func (r *consulResolver) Close() {
	// 关闭资源
}

func (r *consulResolver) Scheme() string {
	return consulScheme
}

func (r *consulResolver) watch() {
	// 在此处可以监听 Consul 服务变化，并更新服务实例列表
}
func grpc_main() {

	/*
		cc, err := api.NewClient(api.DefaultConfig())
		if err != nil {
			fmt.Printf("api.NewClient failed, err:%v\n", err)
			return
		}
	*/
	//serviceName := "jobservice"
	// consul服务
	//address := "consul://127.0.0.1:8500/jobservice"

	// 创建一个新的Catalog服务查询实例
	/*
		catalog := cc.Catalog()

		// 查询指定服务的实例
		serviceEntries, _, err := catalog.Service(serviceName, "", nil)
		if err != nil {
			fmt.Println("UUU", err)
		}
		// 打印服务实例的信息
		for _, entry := range serviceEntries {
			fmt.Printf("Service: %s, Address: %s, Port: %d\n", entry.ServiceName, entry.Address, entry.ServicePort)
		}
		serviceMap, err := cc.Agent().ServicesWithFilter("Service==`jobservice`")
		fmt.Println("WWWWWWWWWWC", serviceMap)
		if err != nil {
			fmt.Printf("query service from consul failed, err:%v\n", err)
			return
		}
		// 选一个服务机（这里选最后一个）
		var addr string
		adds := []string{}
		for k, v := range serviceMap {
			fmt.Printf("%s:%#v\n", k, v)
			addr = v.Address + ":" + strconv.Itoa(v.Port)
			adds = append(adds, addr)
		}

		rand.Seed(time.Now().UnixMilli())
		index := rand.Intn(100) % len(adds)
		addr = adds[index]
	*/
	resolver.Register(&consulResolverBuilder{})

	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s:///%s", consulScheme, "jobservice"),                     // gRPC 服务名
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), // 使用轮询负载均衡
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    30 * time.Second,
			Timeout: 10 * time.Second,
		}),
	)
	//conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		fmt.Println("failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewJobServicevRequestClient(conn)
	req := &pb.Request{
		Name: "sunweiming",
		Job:  "doctor",
	}
	response, err := client.GetJobService(context.Background(), req)
	if err != nil {
		fmt.Println("GetJobService Err", err)
	}
	fmt.Printf("response %v\n", response)
}

func main() {
	for i := 0; i < 20; i++ {
		grpc_main()
	}
}
