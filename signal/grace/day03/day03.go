package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

// 平滑关闭 serve
func main() {
	router := gin.Default()
	router.GET("/home", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	server.RegisterOnShutdown(func() {
		log.Println("start execute out shutown")
	})

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Println("Server closed under request")
			} else {
				log.Fatal("Server closed unexpect")
			}
		}
	}()

	<-quit
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Close:", err)
	}
	log.Println("Server exiting")
}
