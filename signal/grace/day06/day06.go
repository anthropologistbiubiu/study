package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func WaitExit(c chan os.Signal, exit func()) {
	for i := range c {
		switch i {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("receive exit signal ", i.String(), ",exit...")
			exit()
			os.Exit(0)
		}
	}
}

// NewShutdownSignal new normal Signal channel
func NewShutdownSignal() chan os.Signal {
	c := make(chan os.Signal)
	// SIGHUP: terminal closed
	// SIGINT: Ctrl+C
	// SIGTERM: program exit
	// SIGQUIT: Ctrl+/
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	return c
}

// Recover the go routine
func Recover() {
	/*
		for _, cleanup := range cleanups {
			cleanup()
		}
	*/
	if err := recover(); err != nil {
		fmt.Println("recover error", err)
	}
}

// GoSafe instead go func()
func GoSafe(ctx context.Context, fn func(ctx context.Context)) {
	/*
		go func(ctx context.Context) {
			defer Recover()
		}(ctx)
	*/
	go func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		if fn != nil {
			fn(ctx)
		}
	}(ctx)
}

func main() {
	// a gin http server
	//gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.GET("/hello", func(context *gin.Context) {
		// 被 gin 所在 goroutine 捕获
		/*
			defer func() {
				err := recover()
				fmt.Println(err)
			}()
		*/
		defer Recover()
		panic("i am panic")
	})

	httpSrv := &http.Server{
		Addr:    "127.0.0.1:8060",
		Handler: g,
	}
	fmt.Println("server run on:", httpSrv.Addr)
	go httpSrv.ListenAndServe()

	// a custom dangerous go routine, 10s later app will crash!!!!
	GoSafe(context.Background(), func(ctx context.Context) {
		//time.Sleep(time.Second * 5)
		panic("dangerous")
	})

	// wait until exit
	signalChan := NewShutdownSignal()
	WaitExit(signalChan, func() {
		// your clean code
		if err := httpSrv.Shutdown(context.Background()); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("http server closed")
	})
}
