package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var jwtKey = []byte("my_secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {

	http.HandleFunc("/sign", signFunc)
	http.HandleFunc("/sign", welcomeFunc)
	http.ListenAndServe(":8080", nil)
}

func signFunc(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	// 获取JSON正文并解码为凭据
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// 如果主体结构错误，则返回HTTP错误
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 从我们的map中获取用户的密码
	expectedPassword, ok := users[creds.Username]

	// 如果设置的用户密码与我们收到的密码相同，那么我们可以继续。
	// 如果不是，则返回“未经授权”状态。
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// 在这里声明令牌的到期时间，我们将其保留为5分钟
	expirationTime := time.Now().Add(5 * time.Minute)
	// 创建JWT声明，其中包括用户名和有效时间
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 使用用于签名的算法和令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 创建JWT字符串
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// 如果创建JWT时出错，则返回内部服务器错误
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 最后，我们将客户端cookie token设置为刚刚生成的JWT
	// 我们还设置了与令牌本身相同的cookie到期时间
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
func welcomeFunc(w http.ResponseWriter, r *http.Request) {
	// 我们可以从每个请求的Cookie中获取会话令牌
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// 如果未设置cookie，则返回未授权状态
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// 对于其他类型的错误，返回错误的请求状态。
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 从Cookie获取JWT字符串
	tknStr := c.Value

	// 初始化`Claims`实例
	claims := &Claims{}

	// 解析JWT字符串并将结果存储在`claims`中。
	// 请注意，我们也在此方法中传递了密钥。
	// 如果令牌无效（如果令牌已根据我们设置的登录到期时间过期）或者签名不匹配,此方法会返回错误.
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// 最后，将欢迎消息以及令牌中的用户名返回给用户
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}