package main

import (
	"context"
	"fmt"
	"gateway/protos"
	http "github.com/go-kratos/kratos/v2/transport/http"
	"log"
	//"net/http"
)

func main() {
	httpClient, err := http.NewClient(context.Background())
	// 创建 OrderHTTPClient 实例
	orderClient := protos.NewOrderHTTPClient(httpClient)
	// 准备请求参数
	req := &protos.GetOrderReq{
		OrderId: "123456",
	}
	// 发送 HTTP 请求
	resp, err := orderClient.GetOrderInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("HTTP request failed: %v", err)
	}
	// 处理响应结果
	fmt.Printf("Order info: %+v\n", resp)
}
