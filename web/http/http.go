package main

import (
	"log"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: ":9090",
		//Handler: http.HandleFunc("http", httpRequest),
	}
	http.HandleFunc("/fetch", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL)
	w.Write([]byte("hello world"))
}
