package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("my-secret-key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/protected", protectedHandler)
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("password", password)

	// 假设在真实场景中你会校验用户名和密码

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Welcome, %s!", claims.Username)))
}
