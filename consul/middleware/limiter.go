package middleware

import (
	utils "consul/utils/limiter"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"time"
)

// 直接在这里封装一个限流中间件 完成对服务端的限流请求的任务，后期可以优化为对限流外的请求进行阻塞等待
func LimiterInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 地址
	})
	// 限流参数
	maxTokens := 10 // 最大令牌数
	refillRate := 2 // 令牌生成速率（每秒）
	refillInterval := time.Second / time.Duration(refillRate)
	limiter := utils.NewLimiterBucket("bucket_limiter", "limiter_timestamp",
		int64(maxTokens), refillInterval, client)
	if !limiter.Allow() {
		fmt.Println("limiter success!")
		return nil, nil
	}
	response, err := handler(ctx, req)
	if err != nil {
		return nil, nil
	}
	return response, nil
}
