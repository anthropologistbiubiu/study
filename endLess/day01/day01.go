package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gorilla/mux"
)

/*
首先是看懂 endless 不停机重启的过程；然后看明白 在不停机重启的过程中。
开启一个协程。这个协程中 注册一个信号量。监听信号量。然后传递信号量给主进程。查看主进程的执行。
二，如果在协程内不注册协程。完成程序的优雅退出。
*/
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("模拟数据库操作")
	panic("error")
	time.Sleep(10 * time.Second)
}

func main() {
	mux1 := mux.NewRouter()
	mux1.HandleFunc("/sleep", handler)
	srv := endless.NewServer("127.0.0.1:5003", mux1)
	sigHooks := map[os.Signal]func(){
		os.Interrupt:    func() { fmt.Println("test data1") },
		syscall.SIGTERM: func() { fmt.Println("test data2") },
	}
	for sig, hook := range sigHooks {
		if _, ok := srv.SignalHooks[endless.PRE_SIGNAL][sig]; ok {
			srv.SignalHooks[endless.PRE_SIGNAL][sig] = append(srv.SignalHooks[endless.PRE_SIGNAL][sig], hook)
		}
	}

	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println("process task!")
			case <-sigchan:
				fmt.Println("模拟关闭消费者 ")
				return
			}
		}
	}()
	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
	log.Println("All servers stopped. Exiting.")
}
