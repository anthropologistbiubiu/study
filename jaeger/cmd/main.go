package main

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func main() {
	// 创建 Jaeger 导出器
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		panic(fmt.Sprintf("failed to initialize exporter: %v", err))
	}

	// 创建 TracerProvider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("test-service"),
		)),
	)
	defer func() { _ = tp.Shutdown(context.Background()) }()

	// 设置全局 TracerProvider
	otel.SetTracerProvider(tp)

	// 获取 Tracer
	tracer := otel.Tracer("test-tracer")

	// 创建 Span
	_, span := tracer.Start(context.Background(), "main")
	defer span.End()

	// 模拟工作
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Trace has been sent to Jaeger")
}
