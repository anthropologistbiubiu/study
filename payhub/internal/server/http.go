package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "payhub/api/helloworld/v1"
	"payhub/internal/conf"
	"payhub/internal/middleware"
	"payhub/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, pay *service.PaymentOrderService, logger log.Logger) *http.Server {
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
			middleware.RateLimitMiddleware(),
			middleware.ApiAuthMiddleWare(),
		),
		/*
			http.Middleware(
			),
		*/
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterPaymentSerivceHTTPServer(srv, pay)
	return srv
}
