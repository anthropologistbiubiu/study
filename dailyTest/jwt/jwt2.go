// resource_service.go
package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var secretKey = []byte("your-secret-key") // 替换为实际的密钥

func main() {
	router := mux.NewRouter()

	// 受保护的路由
	router.HandleFunc("/protected", protectedEndpoint).Methods("GET")

	// 启动服务器
	http.Handle("/", router)
	fmt.Println("Resource service is running on port 8082...")
	http.ListenAndServe(":8082", nil)
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	// 从请求头中获取JWT令牌
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 解析JWT令牌
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 如果令牌有效，返回受保护资源
	w.Write([]byte("Hello, this is a protected resource!"))
}

func extractTokenFromHeader(header string) (string, error) {
	// 从 Authorization 头中提取令牌
	splitToken := strings.Split(header, "Bearer ")
	if len(splitToken) != 2 {
		return "", fmt.Errorf("Invalid token format")
	}
	return strings.TrimSpace(splitToken[1]), nil
}
