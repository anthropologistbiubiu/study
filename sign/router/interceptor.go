package router

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"sign/utils/log"
	"time"
)

func AccessLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	msg := ""
	response, err := handler(ctx, req)
	cost := time.Since(start)
	p, _ := peer.FromContext(ctx)
	ip := p.Addr.String()
	request, err := json.Marshal(req)
	if err != nil {
		log.Info(msg, zap.String("error", fmt.Sprintf("%s", err)))
	}
	resdata, err := json.Marshal(response)
	if err != nil {
		log.Info(msg, zap.String("string", fmt.Sprintf("%s", err)))
	}
	log.Info(msg, zap.Any("method", info.FullMethod),
		zap.Any("request", string(request)),
		zap.Any("response", string(resdata)),
		zap.String("ip", ip),
		zap.Duration("cost", cost),
	)
	return response, nil
}
