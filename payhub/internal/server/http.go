package server

import (
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	jwt2 "github.com/golang-jwt/jwt/v5"
	v1 "payhub/api/helloworld/v1"
	"payhub/internal/conf"
	"payhub/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type MyLimiter struct {
}

func (l MyLimiter) Allow() (ratelimit.DoneFunc, error) {
	return nil, nil
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, pay *service.PaymentOrderService, logger log.Logger) *http.Server {
	/// type Handler func(ctx context.Context, req interface{}) (interface{}, error)
	var opts = []http.ServerOption{
		http.Middleware(
			jwt.Server(func(token *jwt2.Token) (interface{}, error) {
				return []byte("testKey"), nil
			}),
		),
		http.Middleware(
			ratelimit.Server((ratelimit.WithLimiter(bbr.NewLimiter()))),
		),
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
