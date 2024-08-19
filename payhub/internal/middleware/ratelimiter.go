package middleware

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	limiter "github.com/juju/ratelimit"
	"strings"
)

func AccessLogMiddleware(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if ok {
				logger.Log(log.LevelInfo, "method", tr.Operation(), "path", tr.Endpoint())
			}
			return handler(ctx, req)
		}
	}
}

var WhiteList = []string{
	"192.168.1.1",
	"10.0.0.1",
	"127.0.0.1",
}

func IpWhiteMiddleware(whiteList []string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				clientIP := getClientIP(tr)
				if !isIPWhitelisted(clientIP, whiteList) {
					return nil, fmt.Errorf("IP %s is not allowed", clientIP)
				}
				fmt.Printf("clientIP:%s is allowed\n", clientIP)
			} else {
				return nil, fmt.Errorf("could not extract transport information from context")
			}
			return handler(ctx, req)
		}
	}
}
func getClientIP(tr transport.Transporter) string {
	if ht, ok := tr.(*http.Transport); ok {
		req := ht.Request()

		// 尝试从 X-Forwarded-For 头中获取
		xff := req.Header.Get("X-Forwarded-For")
		if xff != "" {
			ips := strings.Split(xff, ",")
			return strings.TrimSpace(ips[0])
		}

		// 尝试从 X-Real-Ip 头中获取
		xri := req.Header.Get("X-Real-Ip")
		if xri != "" {
			return xri
		}

		// 最后从请求的远程地址中获取
		ip := req.RemoteAddr
		if colon := strings.LastIndex(ip, ":"); colon != -1 {
			ip = ip[:colon]
		}
		return ip
	}
	return ""
}

func isIPWhitelisted(ip string, whiteList []string) bool {
	for _, whitelistedIP := range whiteList {
		if ip == whitelistedIP {
			return true
		}
	}
	return false
}
func IpWhiteMiddleware1() middleware.Middleware {

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				xff := tr.RequestHeader().Get("X-Forwarded-For")
				fmt.Println("|||||||||||", xff)
				if len(xff) > 0 {
					ips := strings.Split(xff, ",")
					if len(ips) > 0 {
						fmt.Println(ips[0])
					}
				} else {
					return nil, fmt.Errorf("parse ip error")
				}
			} else {
				return nil, fmt.Errorf("bad request")
			}
			return handler(ctx, req)
		}
	}
}

func RateLimitMiddleware1() middleware.Middleware {
	bucket := limiter.NewBucketWithRate(10, 10) // 每秒1个请求，最多积累5个
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if bucket.TakeAvailable(1) < 1 {
				return nil, fmt.Errorf("rate limit exceeded")
			}
			fmt.Printf("http server1 req rate is normal,%f \n", bucket.Rate())
			return handler(ctx, req)
		}
	}
}

func RateLimitMiddleware2() middleware.Middleware {
	bucket := limiter.NewBucketWithRate(1, 10) // 每秒1个请求，最多积累5个
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if bucket.TakeAvailable(1) < 1 {
				return nil, fmt.Errorf("rate limit exceeded")
			}
			fmt.Printf("http server2 req rate is normal,%f \n", bucket.Rate())
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

/*
var jwtSecret = []byte("your-secret-key")

func JWTMiddleware() middleware.Middleware {
	return jwt.Server(func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	}, jwt.WithSigningMethod(jwt.SigningMethodHS256))
}

*/
