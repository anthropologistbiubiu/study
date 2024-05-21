package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	http.HandleFunc("/", serveStaticFiles)
	http.HandleFunc("/api/hello", helloHandler)

	// 设置服务器配置
	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Starting server on :8080")
	log.Fatal(server.ListenAndServe())
}

// serveStaticFiles 处理静态文件请求
func serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "static/index.html")
		return
	}

	// 获取文件路径
	path := filepath.Join("static", r.URL.Path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, path)
}

// helloHandler 处理API请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "Hello, World!"}`)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
