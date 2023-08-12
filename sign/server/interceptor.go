package main

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"sign/utils/log"
	"time"
)

// 首先是这里是etcd 的服务发现 ，使用的是分布式签名系统

// 这里主要写一个拦截器文件来实现拦截器对grpc服务access日志的封装

// 1.怎么写拦截器
// 只需要实现某个接口就好
func accessLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	//start := time.Now()
	// 执行RPC调用
	path := "sunweiming"
	//cost := ""
	log.Info(path,
		zap.Duration("cost", time.Duration(5)),
		zap.Any("request", req),
		zap.Any("reponse", info))
	// 记录访问日志
	// 访问日志的拦截器需要自定义   // 可以参考gin框架
	// log.Info(zap.String("requst", req.(string)))
	return nil, nil
}

// 2.access日志的封装过程

// 3.使用grpc 元数据来实现封装 jwt 的用户认证过程

// 4.signServer 配置redis 分布式缓存
