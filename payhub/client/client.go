package main

import (
	"context"
	"errors"
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
	"sync"
	"time"
)

type customBreaker struct {
	breaker      kratosCircuitbreaker.CircuitBreaker
	name         string
	resetTimeout time.Duration
	openTime     time.Time
	isOpen       bool
	mu           sync.Mutex
}

func (cb *customBreaker) Allow() error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.isOpen {
		if time.Since(cb.openTime) >= cb.resetTimeout {
			// 冷却时间已过，尝试关闭熔断器
			cb.isOpen = false
			log.Printf("Circuit breaker '%s' cooling period over, attempting to close", cb.name)
		} else {
			// 仍在冷却时间内，拒绝请求
			log.Printf("Circuit breaker '%s' is OPEN,is in coding time: Request blocked", cb.name)
			return kratosCircuitbreaker.ErrNotAllowed
		}
	}

	err := cb.breaker.Allow()
	if err != nil {
		// 熔断器触发，记录打开时间
		cb.isOpen = true
		cb.openTime = time.Now()
		log.Printf("Circuit breaker '%s' transitioned to OPEN state", cb.name)
		return kratosCircuitbreaker.ErrNotAllowed
	}
	// 请求被允许
	log.Printf("Circuit breaker '%s' transitioned to Closed state", cb.name)
	return nil
}

func (cb *customBreaker) MarkSuccess() {
	cb.breaker.MarkSuccess()
}

func (cb *customBreaker) MarkFailed() {
	cb.breaker.MarkFailed()
}

type loggingBreaker struct {
	kratosCircuitbreaker.CircuitBreaker
	name         string
	resetTimeout time.Duration
	openTime     time.Time
}

// 重写与自定义熔断器

func (lb *loggingBreaker) Allow() error {
	err := lb.CircuitBreaker.Allow()
	if err != nil {
		log.Printf("Circuit breaker '%s' is open, request blocked\n", lb.name)
	} else {
		log.Printf("Circuit breaker '%s' is closed, request allowed\n", lb.name)
	}
	return fmt.Errorf("request is not allowed:%v", err)
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
			sre.WithBucket(10),
			sre.WithRequest(10),
			sre.WithSuccess(0.9),
		)
		return &customBreaker{
			breaker:      breaker,
			name:         "my-custom-breaker",
			resetTimeout: time.Second * 300, // 冷却时间
			isOpen:       false,
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
	for i := 0; i < 50; i++ {
		client := v1.NewPaymentSerivceHTTPClient(hConn)
		_, err := client.CreatePaymentOrder(context.Background(), &v1.PaymentCreateRequest{
			Merchantid: "merchant123",
			Amount:     "456",
		})
		if err != nil {
			if errors.Is(err, kratosCircuitbreaker.ErrNotAllowed) {
				log.Printf("Request %d failed: circuit breaker is open\n", i)
			} else {
				log.Printf("Request %d failed;err:%v", i, err)
			}
		} else {
			log.Printf("Request %d succeeded\n", i)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
