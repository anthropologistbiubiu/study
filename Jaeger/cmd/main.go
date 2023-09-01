package main

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

func main() {
	// 初始化 Jaeger 导出器
	exporter, err := jaeger.New(jaeger.WithAgentEndpoint(jaeger.WithAgentHost("localhost"), jaeger.WithAgentPort(6831)))
	if err != nil {
		log.Fatalf("Failed to create Jaeger exporter: %v", err)
	}
	defer exporter.Shutdown(context.Background())

	// 创建全局的 Tracer
	provider := otel.GetTracerProvider()
	provider.RegisterResource(resource.NewWithAttributes(resource.Attributes{
		"service.name": "grpc-service",
	}))
	provider.RegisterSpanProcessor(exporter)

	// 创建 gRPC 服务器
	server := grpc.NewServer(
		grpc.UnaryInterceptor(otelUnaryServerInterceptor()),
	)

	// 注册 gRPC 服务
	// 在这里添加你的 gRPC 服务定义和实现

	// 启动 gRPC 服务器
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// otelUnaryServerInterceptor 返回 gRPC 的 UnaryServerInterceptor，用于链路追踪
func otelUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		tr := otel.Tracer("grpc-server")
		ctx, span := tr.Start(ctx, info.FullMethod)
		defer span.End()

		// 可以在 span 上添加属性、事件等信息

		resp, err := handler(ctx, req)

		if err != nil {
			st, _ := status.FromError(err)
			span.SetStatus(codes.Error, st.Message())
		}

		return resp, err
	}
}
