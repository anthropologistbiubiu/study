package main

import (
	"context"
	"fmt"
	kratosCircuitbreaker "github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-kratos/aegis/circuitbreaker/sre"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/filter"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	kratos_http "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hashicorp/consul/api"
	"log"
	"payhub/api/v1"
	"time"
)

type loggingBreaker struct {
	kratosCircuitbreaker.CircuitBreaker
	name string
}

func (lb *loggingBreaker) Allow() error {
	err := lb.CircuitBreaker.Allow()
	if err != nil {
		log.Printf("Circuit breaker '%s' is open, request blocked\n", lb.name)
	} else {
		log.Printf("Circuit breaker '%s' is closed, request allowed\n", lb.name)
	}
	return err
}

func (lb *loggingBreaker) MarkSuccess() {
	lb.CircuitBreaker.MarkSuccess()
	log.Printf("Circuit breaker '%s': request succeeded\n", lb.name)
}

func (lb *loggingBreaker) MarkFailed() {
	lb.CircuitBreaker.MarkFailed()
	log.Printf("Circuit breaker '%s': request failed\n", lb.name)
}
func main() {
	// 创建 Consul 客户端
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	// 使用 Consul 作为服务发现
	r := consul.New(consulClient)
	// 创建路由 Filter：筛选版本号为"2.0.0"的实例
	versionFilter := filter.Version("1.0.0")
	// 设置全局的 Selector，使用 wrr 算法
	selector.SetGlobalSelector(wrr.NewBuilder())
	// 配置熔断器的参数
	// 定义生成熔断器的函数
	genBreakerFunc := func() kratosCircuitbreaker.CircuitBreaker {
		// 创建一个 SRE 熔断器，可以自定义参数
		breaker := sre.NewBreaker(
			sre.WithWindow(time.Second*10),
			sre.WithBucket(1),
			sre.WithSuccess(2),
		)
		return &loggingBreaker{
			CircuitBreaker: breaker,
			name:           "my-circuit-breaker",
		}
	}
	// 创建 HTTP 客户端
	hConn, err := kratos_http.NewClient(
		context.Background(),
		//kratos_http.WithEndpoint("discovery:///payhub"),
		kratos_http.WithDiscovery(r),
		kratos_http.WithNodeFilter(versionFilter),
		kratos_http.WithBlock(),
		kratos_http.WithMiddleware(
			circuitbreaker.Client(
				circuitbreaker.WithCircuitBreaker(genBreakerFunc),
			),
		),
		kratos_http.WithEndpoint("127.0.0.1:8005"),
	)
	if err != nil {
		panic(err)
	}
	// 使用客户端发送请求
	client := v1.NewPaymentSerivceHTTPClient(hConn)
	rsp, err := client.CreatePaymentOrder(context.Background(), &v1.PaymentCreateRequest{
		Merchantid: "merchant123",
		Amount:     "456",
	})
	fmt.Printf("rsp:%+v,err:%v \n", rsp, err)
}
