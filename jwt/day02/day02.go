package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var jwtKey = []byte("your-secret-key")

func main() {
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/profile", authMiddleware(handleProfile))

	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// 实际场景中会有用户验证过程，这里简化为一个示例用户名
	username := "user123"

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: token.Claims.(jwt.MapClaims)["exp"].(time.Time),
	})
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Login successful")
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractToken(r)
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Accessing profile...")
}

func extractToken(r *http.Request) string {
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		return ""
	}
	return tokenCookie.Value
}
