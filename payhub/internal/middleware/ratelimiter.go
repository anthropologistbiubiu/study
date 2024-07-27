package middleware

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	limiter "github.com/juju/ratelimit"
)

func RateLimitMiddleware() middleware.Middleware {
	bucket := limiter.NewBucketWithRate(1, 10) // 每秒1个请求，最多积累5个
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if bucket.TakeAvailable(1) < 1 {
				return nil, fmt.Errorf("rate limit exceeded")
			}
			return handler(ctx, req)
		}
	}
}

const ApiKey = "sunweiming"

func ApiAuthMiddleWare() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				apiKey := tr.RequestHeader().Get("apiKey")
				if apiKey != ApiKey {
					return nil, fmt.Errorf("api key error")
				}
				fmt.Println("apiKey is correct")
			} else {
				return nil, fmt.Errorf("api key error")
			}
			return handler(ctx, req)
		}
	}
}

var jwtSecret = []byte("your-secret-key")

func JWTMiddleware() middleware.Middleware {
	return jwt.Server(func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	}, jwt.WithSigningMethod(jwt.SigningMethodHS256))
}
