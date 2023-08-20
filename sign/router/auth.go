package router

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	port      = ":50051"
	secretKey = "my_secret_key"
)

// 可以自行先判断一次token 是否过期，如果过期就重新 请求，如果没有过期就重新请求
// 传递数据使用metadata 的方式去传递数据。

// 服务端使用拦截器对请求拦截，首先判断token 的有效性，如果有效，执行业务，如果无效

// 如果是有效期过期就 重新签发token 和过期时间

// 那么客户端怎么限定对哪些数据进行保护 限制为两台服务之间的通信过程。

func JwtAuthorizeInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	token := md.Get("authorization")
	if len(token) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is missing")
	}
	// 验证 JWT
	claims := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(token[0], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !tkn.Valid {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	// 进行鉴权逻辑，这里可以根据 claims 中的信息来进行授权判断

	// 也应该判断是否存在这个用户
	if claims["username"] != "admin" {
		return nil, status.Errorf(codes.PermissionDenied, "user does not have permission")
	}
	return handler(ctx, req)
}

func generatorJwtToken() {

}

func refreshJwtToken() {

}
