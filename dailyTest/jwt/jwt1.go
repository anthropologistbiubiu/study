// user_service.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var secretKey = []byte("your-secret-key") // 替换为实际的密钥

// User 结构体用于模拟用户信息
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	router := mux.NewRouter()

	// 登录接口
	router.HandleFunc("/login", handleLogin).Methods("POST")

	// 启动服务器
	http.Handle("/", router)
	fmt.Println("User service is running on port 8081...")
	http.ListenAndServe(":8081", nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// 在实际应用中，这里可以验证用户名和密码是否正确
	// 这里简单演示，假设用户名为 "user"，密码为 "password"
	if user.Username == "user" && user.Password == "password" {
		// 生成JWT令牌
		token, err := generateToken(user.Username)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		// 返回令牌给前端
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}

func generateToken(username string) (string, error) {
	// 设置过期时间为 1 小时
	expirationTime := time.Now().Add(1 * time.Hour)

	// 创建JWT令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
	})

	// 使用密钥进行签名
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
