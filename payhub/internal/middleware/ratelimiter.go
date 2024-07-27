package middleware

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	//"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	limiter "github.com/juju/ratelimit"
)

func RateLimitMiddleware() middleware.Middleware {
	bucket := limiter.NewBucketWithRate(1, 1) // 每秒1个请求，最多积累5个
	return func(handler middleware.Handler) middleware.Handler {
		fmt.Println("||||||||||")
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if bucket.TakeAvailable(1) < 1 {
				return nil, fmt.Errorf("rate limit exceeded")
			}
			return handler(ctx, req)
		}
	}
}
