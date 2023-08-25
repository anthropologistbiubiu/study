package main

import (
	"consul/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"math/rand"
	"strconv"
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

func main() {

	cc, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Printf("api.NewClient failed, err:%v\n", err)
		return
	}
	serviceName := "jobservice"

	// 创建一个新的Catalog服务查询实例
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
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
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
