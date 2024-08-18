package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/filter"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hashicorp/consul/api"
	"io"
	url2 "net/url"
)

func main() {
	// 创建 Consul 客户端
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	// 使用 Consul 作为服务发现
	r := consul.New(consulClient)
	// 创建路由 Filter：筛选版本号为"2.0.0"的实例
	versionFilter := filter.Version("2.0.0")
	// 设置全局的 Selector，使用 wrr 算法
	selector.SetGlobalSelector(wrr.NewBuilder())

	// 创建 HTTP 客户端
	hConn, err := http.NewClient(
		context.Background(),
		http.WithEndpoint("discovery:///payhub"),
		http.WithDiscovery(r),
		http.WithNodeFilter(versionFilter),
	)
	if err != nil {
		panic(err)
	}
	// 使用客户端发送请求
	url := "/payment/create"
	req := &http.Request{
		Method: "POST",
		URL: &url2.URL{
			Path: url,
		},
	}
	response, err := hConn.Do(req)
	if err != nil {
		panic(err)
	}
	if resBytes, err := io.ReadAll(response.Body); err != nil {
		fmt.Println("///////////", err, resBytes)
	}
}
