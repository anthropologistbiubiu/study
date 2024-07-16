package server

import (
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	v1 "payhub/api/helloworld/v1"
	"payhub/internal/conf"
	"payhub/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, pay *service.PaymentOrderService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
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
	http.Middleware(
		jwt.Server(func(ctx, token string) (interface{}, error) {
			return []byte("testkey"), nil
		}))
	v1.RegisterPaymentSerivceHTTPServer(srv, pay)
	return srv
}
