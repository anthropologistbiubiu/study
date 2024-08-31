package server

import (
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"payhub/api/v1"

	//prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus"

	"payhub/internal/conf"
	"payhub/internal/middleware"
	"payhub/internal/service"
)

var (
	// Name is the name of the compiled software.
	Name = "metrics"
	// Version is the version of the compiled software.
	// Version = "v1.0.0"

	_metricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "duration_sec",
		Help:      "server requests duration(sec).",
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"kind", "operation"})

	_metricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer1(c *conf.Server, pay *service.PaymentOrderService, logger log.Logger) *http.Server {
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
			middleware.RateLimitMiddleware1(),
			//middleware.ApiAuthMiddleWare(),
		),
		/*
			http.Middleware(
			),
		*/
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
			metrics.Server(
				metrics.WithRequests(middleware.RequestCounter), // 中间件采集业务请求数
			),
			//middleware.ApiAuthMiddleWare(),
		),
		/*
			http.Middleware(
			),
		*/
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
	v1.RegisterPaymentSerivceHTTPServer(srv, pay)
	return srv
}
