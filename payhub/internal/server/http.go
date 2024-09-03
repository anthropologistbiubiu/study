package server

import (
	_ "context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	prom "go.opentelemetry.io/otel/exporters/prometheus"
	_ "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"payhub/api/v1"
	"payhub/internal/conf"
	"payhub/internal/middleware"
	"payhub/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer1(c *conf.Server, pay *service.PaymentOrderService, logger log.Logger) *http.Server {
	promExporter, err := prom.New()
	if err != nil {
		log.Fatalf("failed to initialize Prometheus exporter: %v", err)
	}
	meterProvider := metric.NewMeterProvider(metric.WithReader(promExporter))
	otel.SetMeterProvider(meterProvider)
	meter := otel.Meter("my-payhub1-meter")
	if err != nil {
		log.Fatalf("failed to create Int64Counter: %v", err)
	}
	requestsCounter, err := meter.Int64Counter(
		"payhub1_http_requests_total",
	)
	if err != nil {
		log.Fatalf("failed to create Int64Counter: %v", err)
	}
	var opts = []http.ServerOption{
		http.Middleware(
			/*
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte("testKey"), nil
				}),
			*/
			//ratelimit.Server(ratelimit.WithLimiter(mylimiter)),
			//middleware.IpWhiteMiddleware(middleware.WhiteList),
			middleware.RateLimitMiddleware1(),
			middleware.AccessLogMiddleware(logger),
			metrics.Server(
				metrics.WithRequests(requestsCounter)),
		),
	}
	if c.Http1.Network != "" {
		opts = append(opts, http.Network(c.Http1.Network))
	}
	if c.Http1.Addr != "" {
		opts = append(opts, http.Address(c.Http1.Addr))
	}
	if c.Http1.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http1.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.Handle("/metrics", promhttp.Handler())
	v1.RegisterPaymentSerivceHTTPServer(srv, pay)
	return srv
}

func NewHTTPServer2(c *conf.Server, pay *service.PaymentOrderService, logger log.Logger) *http.Server {
	/// type Handler func(ctx context.Context, req interface{}) (interface{}, error)

	//mylimiter := bbr.NewLimiter(bbr.WithWindow(1*time.Hour), bbr.WithBucket(1))

	var opts = []http.ServerOption{
		http.Middleware(
			/*
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte("testKey"), nil
				}),
			*/
			//ratelimit.Server(ratelimit.WithLimiter(mylimiter)),
			//middleware.IpWhiteMiddleware(middleware.WhiteList),
			middleware.RateLimitMiddleware2(),
			middleware.AccessLogMiddleware(logger),
		),
	}
	if c.Http2.Network != "" {
		opts = append(opts, http.Network(c.Http2.Network))
	}
	if c.Http2.Addr != "" {
		opts = append(opts, http.Address(c.Http2.Addr))
	}
	if c.Http2.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http2.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	//srv.Handle("/metrics", promhttp.Handler())
	v1.RegisterPaymentSerivceHTTPServer(srv, pay)
	return srv
}
